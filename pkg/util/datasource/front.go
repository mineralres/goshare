package datasource

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
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
	g.POST("/updateTick", f.updateTick)
	g.GET("/ws/uploadTick", f.uploadTick)

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

func (f *Front) subscribe(c *gin.Context) {
	log.Println("subscribe")
	defer log.Println("subscribe exited")
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	// 先接收订阅
	// conn.SetReadDeadline(time.Now().Add(time.Second * 15))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var connLock sync.RWMutex
	for {
		// 读取
		t, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(t, p, err)
			return
		}
		if t != websocket.TextMessage {
			log.Println("t != websocket.TextMessage")
			return
		}
		var l pb.SymbolList
		err = jsonpb.Unmarshal(bytes.NewReader(p), &l)
		if err != nil {
			log.Println(string(p), err)
			return
		}
		go func() {
			var req pb.ReqSubscribe
			req.List = l.List
			ch := make(chan *pb.MarketDataSnapshot, 100)
			f.cache.subscribe(&req, ch)
			for {
				select {
				case md := <-ch:
					str, err := (&jsonpb.Marshaler{EmitDefaults: true}).MarshalToString(md)
					if err != nil {
						log.Println("err", err)
						return
					}
					connLock.Lock()
					err = conn.WriteMessage(websocket.TextMessage, []byte(str))
					connLock.Unlock()
					if err != nil {
						log.Println("err", err)
						return
					}
				case <-ctx.Done():
					f.cache.unsubscribe(&pb.ReqUnSubscribe{}, ch)
					break
				}
			}
		}()
	}
}

func (f *Front) setTradingInstrument(c *gin.Context) {
	if !f.checkToken(c) {
		return
	}
	var req pb.ReqSetTradingInstrument
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
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

func (f *Front) updateTick(c *gin.Context) {
	if !f.checkToken(c) {
		return
	}
}

func (f *Front) uploadTick(c *gin.Context) {
	if !f.checkToken(c) {
		return
	}
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()
	// 先接收订阅
	// conn.SetReadDeadline(time.Now().Add(time.Second * 15))
	var n int
	for {
		// 读取
		t, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(t, p, err)
			return
		}
		if t != websocket.BinaryMessage {
			log.Println("t != websocket.TextMessage")
			return
		}
		var md pb.MarketDataSnapshot
		err = proto.Unmarshal(p, &md)
		if err != nil {
			continue
		}
		f.cache.Update(&md)
		n++
		log.Printf("received[%d]", n)
	}
}
