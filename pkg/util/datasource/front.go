package datasource

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/pb"
)

var (
	wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Front 前置机
type Front struct {
	cache   *XCache
	options Options
}

// Options 选项
type Options struct {
	Token string `json:"token"`
	Port  int    `json:"port"`
}

// MakeFront MakeFront
func MakeFront(op *Options) *Front {
	ret := &Front{options: *op}
	ldb := MakeXLevelDB()
	ret.cache = makeXCache(ldb, ldb)
	return ret
}

// Run run
func (f *Front) Run() {
	r := gin.New()
	g := r.Group("/api/v1")
	g.POST("/setTradingInstrument", f.setTradingInstrument)
	g.POST("/getTradingInstrument", f.getTradingInstrument)
	g.GET("/ws/stream", f.handleStream)

	r.Run(fmt.Sprintf(":%d", f.options.Port))
}

type response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

func (f *Front) checkToken(c *gin.Context) bool {
	token := c.GetHeader("token")
	if token == f.options.Token {
		return true
	}
	log.Printf("check token failed %s, should be %s", token, f.options.Token)
	c.JSON(200, &response{Success: false, Code: -1, Msg: "invalide token"})
	return false
}

func (f *Front) setTradingInstrument(c *gin.Context) {
	if !f.checkToken(c) {
		return
	}
	var req pb.ReqSetTradingInstrument
	err := c.BindJSON(&req)
	if err != nil {
		log.Println("setTradingInstrument", err)
		return
	}
	log.Printf("setTradingInstrument [%d] 个 ", len(req.List))
	f.cache.setTradingInstrument(&req)
}

func (f *Front) getTradingInstrument(c *gin.Context) {
	var req pb.ReqGetTradingInstrument
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	f.cache.getTradingInstrument(&req)
}

func (f *Front) handleStream(c *gin.Context) {
	if !f.checkToken(c) {
		return
	}
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()
	chOut := make(chan *pb.Message, 100)
	chSub := make(chan *pb.MarketDataSnapshot, 9999)
	go func() {
		// 写函数
		for {
			select {
			case msg := <-chOut:
				d, err := proto.Marshal(msg)
				if err != nil {
					continue
				}
				err = conn.WriteMessage(websocket.BinaryMessage, d)
				if err != nil {
					return
				}
			case md := <-chSub:
				msg := new(pb.Message)
				msg.Type = pb.MessageType_RTN_MARKET_DATA_UPDATE
				d, err := proto.Marshal(md)
				if err != nil {
					continue
				}
				msg.Data = d
				chOut <- msg
			}
		}
	}()
	// 先接收订阅
	for {
		// 读取
		t, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if t != websocket.BinaryMessage {
			continue
		}
		msg := new(pb.Message)
		err = proto.Unmarshal(p, msg)
		if err != nil {
			continue
		}
		switch msg.Type {
		case pb.MessageType_UPLOAD_TICK:
			var tick pb.MarketDataSnapshot
			if err = proto.Unmarshal(msg.Data, &tick); err != nil {
				log.Println("req", err)
				continue
			}
			f.cache.Update(&tick)
		case pb.MessageType_REQ_UNSUBSCRIBE_MARKET_DATA:
			var req pb.ReqUnSubscribe
			if err = proto.Unmarshal(msg.Data, &req); err != nil {
				log.Println("req", req, err)
				continue
			}
			f.cache.unsubscribe(&req, chSub)
		case pb.MessageType_REQ_SUBSCRIBE_MARKET_DATA:
			// 订阅行情
			log.Printf("front received msg[%v] len[%d]", msg.Type, len(msg.Data))
			var req pb.ReqSubscribe
			if err = proto.Unmarshal(msg.Data, &req); err != nil {
				log.Println("req", req, err)
				continue
			}
			f.cache.subscribe(&req, chSub)
			defer f.cache.unsubscribe(&pb.ReqUnSubscribe{List: req.List}, chSub)
		case pb.MessageType_HEATBEAT:
			// 心跳
			log.Printf("front received msg[%v]", msg.Type)
			chOut <- msg
		default:
			log.Printf("front received msg[%v]", msg.Type)
		}
	}
}
