package datasource

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/api"
	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	"github.com/mineralres/goshare/pkg/util"
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
	wsIndex int64
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

func (f *Front) checkToken(c *gin.Context) bool {
	token := c.GetHeader("token")
	if token == f.options.Token {
		return true
	}
	log.Printf("check token failed %s, should be %s", token, f.options.Token)
	type response struct {
		Success bool   `json:"success"`
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
	}
	c.JSON(200, &response{Success: false, Code: -1, Msg: "invalide token"})
	return false
}

func (f *Front) setTradingInstrument(r *http.Request) (interface{}, error) {
	var req pb.ReqSetTradingInstrument
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	log.Printf("setTradingInstrument [%d] 个 ", len(req.List))
	f.cache.setTradingInstrument(&req)
	return nil, nil
}

func (f *Front) getTradingInstrument(r *http.Request) (interface{}, error) {
	var req pb.ReqGetTradingInstrument
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	f.cache.getTradingInstrument(&req)
	return nil, nil
}

func (f *Front) handleStream(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()
	chOut := make(chan *pb.Message, 100)
	chSub := make(chan *pb.MarketDataSnapshot, 9999)
	index := atomic.AddInt64(&f.wsIndex, 1)
	ticker := time.NewTicker(time.Second * 10)
	var nUpload int64
	go func() {
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
		// 写函数
		for {
			select {
			case <-ticker.C:
				err := fsend(pb.MessageType_HEATBEAT, nil)
				if err != nil {
					return
				}
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
		conn.SetReadDeadline(time.Now().Add(15 * time.Second))
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
			nUpload++
			f.cache.Update(&tick)
			if nUpload%100 == 0 {
				log.Printf("front[%d] received msg[%v] len[%d]", index, msg.Type, len(msg.Data))
			}
		case pb.MessageType_REQ_UNSUBSCRIBE_MARKET_DATA:
			var req pb.ReqUnSubscribe
			if err = proto.Unmarshal(msg.Data, &req); err != nil {
				log.Println("req", req, err)
				continue
			}
			f.cache.unsubscribe(&req, chSub)
		case pb.MessageType_REQ_SUBSCRIBE_MARKET_DATA:
			// 订阅行情
			log.Printf("front[%d] received msg[%v] len[%d]", index, msg.Type, len(msg.Data))
			var req pb.ReqSubscribe
			if err = proto.Unmarshal(msg.Data, &req); err != nil {
				log.Println("req", req, err)
				continue
			}
			f.cache.subscribe(&req, chSub)
			defer f.cache.unsubscribe(&pb.ReqUnSubscribe{List: req.List}, chSub)
		case pb.MessageType_HEATBEAT:
			// 心跳
			log.Printf("front[%d] received msg[%v]", index, msg.Type)
		default:
			log.Printf("front[%d] received msg[%v]", index, msg.Type)
		}
	}
}

func bindJSON(r *http.Request, obj interface{}) error {
	return json.NewDecoder(r.Body).Decode(obj)
}

func (f *Front) instrumentList(r *http.Request) (interface{}, error) {
	var req pb.ReqGetTradingInstrumentList
	var err error
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return f.cache.ds.TradingInstrumentList(&api.Context{}, &req)
}

func (f *Front) lastTick(r *http.Request) (interface{}, error) {
	var req pb.Symbol
	var err error
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return f.cache.getTick(&req)
}

// Run run
func (f *Front) Run() {
	log.Printf("Datasource running on :%d", f.options.Port)
	util.RunTinyGateway("", f.options.Port, func(path string, w http.ResponseWriter, r *http.Request) (interface{}, error) {
		switch path {
		case "/api/v1/lastTick":
			return f.lastTick(r)
		case "/api/v1/setTradingInstrument":
			return f.setTradingInstrument(r)
		case "/api/v1/getTradingInstrument":
			return f.getTradingInstrument(r)
		case "/api/v1/instrumentList":
			return f.instrumentList(r)
		case "/api/v1/ws/stream":
			f.handleStream(w, r)
			return nil, errors.New("abort")
		}
		return nil, nil
	})
}
