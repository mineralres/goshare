package datasource

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/util"
)

type symbolSubscriber struct {
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
	URL          url.URL `json:"url"`
	Token        string  `json:"token"`
	WithUploader bool    `json:"withUploader"`
}

// MakeClient MakeClient
func MakeClient(options *ClientOptions) *Client {
	ret := &Client{options: *options}
	ret.chUploadTick = make(chan *pb.MarketDataSnapshot, 10000)
	ret.chReqSubscribe = make(chan pb.ReqSubscribe, 100)
	ret.chReqUnSubscribe = make(chan pb.ReqUnSubscribe, 100)
	if options.WithUploader {
		go ret.uploader()
	}
	go ret.subscribeConn()
	return ret
}

// Subscribe Subscribe
func (c *Client) Subscribe(req *pb.ReqSubscribe, ch chan *pb.MarketDataSnapshot) (*pb.RspSubscribe, error) {
	var ret pb.RspSubscribe
	for i := range req.List {
		symbol := *req.List[i]
		v, ok := c.mapSubscriber.Load(symbol)
		if ok {
			p := v.(*symbolSubscriber)
			p.chList = append(p.chList, ch)
			ch <- p.tick
		} else {
			p := &symbolSubscriber{}
			p.chList = append(p.chList, ch)
			c.mapSubscriber.Store(symbol, p)
		}
	}
	c.chReqSubscribe <- *req
	return &ret, nil
}

func (c *Client) subscribeConn() {
	defer func() {
		time.Sleep(time.Second)
		go c.subscribeConn()
	}()
	url := c.options.URL
	if url.Scheme == "https" {
		url.Scheme = "wss"
	} else {
		url.Scheme = "ws"
	}
	url.Path = "/api/v1/ws/subscribe"
	header := make(http.Header)
	header.Add("token", c.options.Token)
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), header)
	if err != nil {
		log.Println("dial:", err, url.String())
		return
	}
	log.Printf("成功连接推送行情 %s ", url.String())
	defer conn.Close()

	// write
	go func() {
		for req := range c.chReqSubscribe {
			var l pb.SymbolList
			l.List = req.List
			str, err := (&jsonpb.Marshaler{}).MarshalToString(&l)
			err = conn.WriteMessage(websocket.TextMessage, []byte(str))
			if err != nil {
				c.chReqSubscribe <- req
				log.Println(err)
				// 连接断开重试
				return
			}
		}
	}()
	// read
	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(t, p, err)
			return
		}
		if t != websocket.TextMessage {
			log.Println("t != websocket.TextMessage")
			return
		}
		md := new(pb.MarketDataSnapshot)
		err = jsonpb.Unmarshal(bytes.NewReader(p), md)
		if err != nil {
			continue
		}
		v, ok := c.mapSubscriber.Load(*md.Symbol)
		if ok {
			sub := v.(*symbolSubscriber)
			sub.tick = md
			for i := range sub.chList {
				sub.chList[i] <- md
			}
		}
	}

}

// UnSubscribe UnSubscribe
func (c *Client) UnSubscribe(req *pb.ReqUnSubscribe, ch chan *pb.MarketDataSnapshot) (*pb.RspUnSubscribe, error) {
	for i := range req.List {
		symbol := *req.List[i]
		v, ok := c.mapSubscriber.Load(symbol)
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

// 负责上传数据
func (c *Client) uploader() {
	defer func() {
		time.Sleep(time.Second)
		go c.uploader()
	}()
	url := c.options.URL
	if url.Scheme == "https" {
		url.Scheme = "wss"
	} else {
		url.Scheme = "ws"
	}
	url.Path = "/api/v1/ws/uploadTick"
	header := make(http.Header)
	header.Add("token", c.options.Token)
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), header)
	if err != nil {
		log.Println("dial:", err, url.String())
		return
	}
	log.Printf("成功连接 %s ", url.String())
	defer conn.Close()

	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			conn.WriteMessage(websocket.TextMessage, []byte("heartbeat"))
		case md := <-c.chUploadTick:
			d, err := proto.Marshal(md)
			if err != nil {
				log.Println(err)
				continue
			}
			err = conn.WriteMessage(websocket.BinaryMessage, d)
			if err != nil {
				log.Println(err)
				// 连接断开重试
				break
			}
		}
	}
}
