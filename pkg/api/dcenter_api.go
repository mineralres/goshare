package api

import (
	"bytes"
	"context"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/pb"
)

func (r *XRouter) dcenterInfo(c *gin.Context) {
	resp, err := r.dcc.GetDCenterInfo(context.Background(), &pb.ReqGetDCenterInfo{})
	if err != nil {
		log.Println(resp, err)
		c.JSON(500, "")
	}
	pbJSON(c, resp)
}

func (r *XRouter) subscribe(c *gin.Context) {
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
			stream, err := r.dcc.Subscribe(ctx, &req)
			if err != nil {
				log.Println("r.dcc.Subscribe", err)
				return
			}
			for {
				md, err := stream.Recv()
				if err != nil {
					log.Println("stream.Recv ", err)
					return
				}
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
			}
		}()
	}
}

func (r *XRouter) getTradingInstrument(c *gin.Context) {
	var req pb.ReqGetTradingInstrument
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	resp, err := r.dcc.GetTradingInstrument(context.Background(), &req)
	if err != nil {
		return
	}
	log.Println(resp, err)
}
