package datasource

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/util"
)

// Client Client
type Client struct {
	options      ClientOptions
	chUploadTick chan *pb.MarketDataSnapshot
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
	if options.WithUploader {
		go ret.uploader()
	}
	return ret
}

// Subscribe Subscribe
func (c *Client) Subscribe(req *pb.ReqSubscribe, ch chan *pb.MarketDataSnapshot) (*pb.RspSubscribe, error) {
	return nil, nil
}

// UnSubscribe UnSubscribe
func (c *Client) UnSubscribe(req *pb.ReqUnSubscribe, ch chan *pb.MarketDataSnapshot) (*pb.RspUnSubscribe, error) {
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

	for md := range c.chUploadTick {
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
