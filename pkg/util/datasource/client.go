package datasource

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/api"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/util"
)

type symbolSubscriber struct {
	symbol pb.Symbol
	tick   *pb.MarketDataSnapshot
	chList []chan *pb.MarketDataSnapshot
}

// Client Client
type Client struct {
	options          ClientOptions
	chUploadTick     chan *pb.MarketDataSnapshot
	mapSubscriber    sync.Map
	chReqSubscribe   chan pb.ReqSubscribe
	chReqUnSubscribe chan pb.ReqUnSubscribe
}

// ClientOptions Options
type ClientOptions struct {
	URL   url.URL `json:"url"`
	Token string  `json:"token"`
}

// MakeClient MakeClient
func MakeClient(options *ClientOptions) *Client {
	ret := &Client{options: *options}
	ret.chUploadTick = make(chan *pb.MarketDataSnapshot, 10000)
	ret.chReqSubscribe = make(chan pb.ReqSubscribe, 100)
	ret.chReqUnSubscribe = make(chan pb.ReqUnSubscribe, 100)
	go ret.clientConn()
	return ret
}

// Subscribe Subscribe
func (c *Client) Subscribe(ctx *api.Context, req *pb.ReqSubscribe, ch chan *pb.MarketDataSnapshot) (*pb.RspSubscribe, error) {
	var ret pb.RspSubscribe
	for i := range req.List {
		symbol := req.List[i]
		v, ok := c.mapSubscriber.Load(makeKey(symbol))
		if ok {
			p := v.(*symbolSubscriber)
			p.chList = append(p.chList, ch)
			ch <- p.tick
		} else {
			p := &symbolSubscriber{symbol: *symbol}
			p.chList = append(p.chList, ch)
			c.mapSubscriber.Store(makeKey(symbol), p)
		}
	}
	c.chReqSubscribe <- *req
	return &ret, nil
}

// UnSubscribe UnSubscribe
func (c *Client) UnSubscribe(ctx *api.Context, req *pb.ReqUnSubscribe, ch chan *pb.MarketDataSnapshot) (*pb.RspUnSubscribe, error) {
	for i := range req.List {
		symbol := req.List[i]
		v, ok := c.mapSubscriber.Load(makeKey(symbol))
		if ok {
			sub := v.(*symbolSubscriber)
			var left []chan *pb.MarketDataSnapshot
			for i := range sub.chList {
				if sub.chList[i] != ch {
					left = append(left, sub.chList[i])
				}
			}
			sub.chList = left
			if len(sub.chList) == 0 {
				c.chReqUnSubscribe <- *req
			}
		}
	}
	return nil, nil
}

// SetTradingInstrument SetTradingInstrument
func (c *Client) SetTradingInstrument(req *pb.ReqSetTradingInstrument) error {
	url := c.options.URL
	var resp pb.RspSetTradingInstrument
	return util.PostSome(fmt.Sprintf("%s://%s/api/v1/setTradingInstrument", url.Scheme, url.Host), c.options.Token, req, resp)
}

// UpdateTick UpdateTick
func (c *Client) UpdateTick(req *pb.MarketDataSnapshot) error {
	c.chUploadTick <- req
	return nil
}

func (c *Client) clientConn() {
	defer func() {
		time.Sleep(time.Second)
		go c.clientConn()
	}()
	url := c.options.URL
	if url.Scheme == "https" {
		url.Scheme = "wss"
	} else {
		url.Scheme = "ws"
	}
	url.Path = "/api/v1/ws/stream"
	header := make(http.Header)
	header.Add("token", c.options.Token)
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), header)
	if err != nil {
		log.Printf("goshare服务器[%s] 连接失败", url.String())
		return
	}
	log.Printf("成功连接推送行情 %s ", url.String())
	defer conn.Close()

	fsend := func(t pb.MessageType, d proto.Message) error {
		var msg pb.Message
		msg.Type = t
		if d != nil {
			out, err := proto.Marshal(d)
			if err != nil {
				return err
			}
			msg.Data = out
		}
		out, err := proto.Marshal(&msg)
		if err != nil {
			return err
		}
		return conn.WriteMessage(websocket.BinaryMessage, out)
	}
	// write
	go func() {
		var req pb.ReqSubscribe
		c.mapSubscriber.Range(func(k, v interface{}) bool {
			s := new(pb.Symbol)
			*s = v.(*symbolSubscriber).symbol
			req.List = append(req.List, s)
			return true
		})
		if len(req.List) > 0 {
			c.chReqSubscribe <- req
			log.Println("重连之后重新发送订阅请求", len(req.List))
		}
		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case md := <-c.chUploadTick:
				err := fsend(pb.MessageType_UPLOAD_TICK, md)
				if err != nil {
					return
				}
			case <-ticker.C:
				err := fsend(pb.MessageType_HEATBEAT, nil)
				if err != nil {
					return
				}
			case req := <-c.chReqUnSubscribe:
				err := fsend(pb.MessageType_REQ_UNSUBSCRIBE_MARKET_DATA, &req)
				if err != nil {
					return
				}
			case req := <-c.chReqSubscribe:
				log.Println("发送订阅请求", len(req.List))
				err := fsend(pb.MessageType_REQ_SUBSCRIBE_MARKET_DATA, &req)
				if err != nil {
					return
				}
			}
		}
	}()
	// read
	for {
		conn.SetReadDeadline(time.Now().Add(15 * time.Second))
		t, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(t, p, err)
			return
		}
		if t != websocket.BinaryMessage {
			log.Println("t != websocket.TextMessage")
			return
		}
		msg := new(pb.Message)
		if err = proto.Unmarshal(p, msg); err != nil {
			continue
		}
		// log.Printf("client received msg [%v] len[%d]", msg.Type, len(msg.Data))
		switch msg.Type {
		case pb.MessageType_HEATBEAT:
			log.Printf("client received msg [%v] len[%d]", msg.Type, len(msg.Data))
		case pb.MessageType_RTN_MARKET_DATA_UPDATE:
			md := new(pb.MarketDataSnapshot)
			if err = proto.Unmarshal(msg.Data, md); err != nil {
				continue
			}
			v, ok := c.mapSubscriber.Load(makeKey(md.Symbol))
			if ok {
				sub := v.(*symbolSubscriber)
				sub.tick = md
				for i := range sub.chList {
					sub.chList[i] <- md
				}
			}
		}
	}

}

// GetKlineSeries GetKlineSeries
func (c *Client) GetKlineSeries(ctx *api.Context, req *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error) {
	return nil, errors.New("unsported")
}

// RGetKlineSeries RGetKlineSeries
func (c *Client) RGetKlineSeries(ctx *api.Context, req *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error) {
	return nil, errors.New("unsported")
}

// GetLastTick GetLastTick
func (c *Client) GetLastTick(ctx *api.Context, req *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	url := c.options.URL
	var resp pb.MarketDataSnapshot
	return &resp, util.PostSome(fmt.Sprintf("%s://%s/api/v1/lastTick", url.Scheme, url.Host), c.options.Token, req, &resp)
}

// GetTickSerires GetTickSerires
func (c *Client) GetTickSerires(ctx *api.Context, req *pb.ReqGetTickSeries) (*pb.RspGetTickSeries, error) {
	return nil, errors.New("unsported")
}

// GetTradingInstrument GetTradingInstrument
func (c *Client) GetTradingInstrument(ctx *api.Context, s *pb.Symbol) (*pb.TradingInstrument, error) {
	return nil, errors.New("unsported")
}

// TradingInstrumentList TradingInstrumentList
func (c *Client) TradingInstrumentList(ctx *api.Context, req *pb.ReqGetTradingInstrumentList) ([]*pb.TradingInstrument, error) {
	url := c.options.URL
	var resp []*pb.TradingInstrument
	return resp, util.PostSome(fmt.Sprintf("%s://%s/api/v1/instrumentList", url.Scheme, url.Host), c.options.Token, req, &resp)
}
