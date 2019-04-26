package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

type wsFront struct {
	grpcConn     *grpc.ClientConn
	grpcEndPoint string
	connLock     sync.RWMutex
	port         int
}

func makeWsFront(grpcEndPoint string, port int) *wsFront {
	if port == 0 {
		panic("invlid ws port")
	}
	var front wsFront
	front.grpcEndPoint = grpcEndPoint
	front.port = port
	return &front
}

func (front *wsFront) makeGClient() pb.GoShareClient {
	front.connLock.Lock()
	defer front.connLock.Unlock()
	if front.grpcConn != nil {
		return pb.NewGoShareClient(front.grpcConn)
	}
	conn, err := grpc.Dial(front.grpcEndPoint, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	if conn == nil {
		panic("")
	}
	front.grpcConn = conn
	return pb.NewGoShareClient(front.grpcConn)
}

// 因为grpc-gateway对stream 的实现不是特别成熟,所以此处先用websocket完成stream方式的推送
func (front *wsFront) run() {
	r := gin.New()
	r.GET("/ws/uploadTick", front.uploadTick)
	s := &http.Server{
		Addr:    ":" + strconv.Itoa(front.port),
		Handler: r,
	}
	s.SetKeepAlivesEnabled(false)
	log.Printf("websocket listen on [%d] ", front.port)
	s.ListenAndServe()
}

// 推送tick到自建数据库
func (front *wsFront) uploadTick(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := front.makeGClient()
	if client == nil {
		return
	}
	for {
		conn.SetReadDeadline(time.Now().Add(15 * time.Second))
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("wsUploadTick exit", err)
			break
		} else {
			if messageType == websocket.BinaryMessage {
				var mds pb.MarketDataSnapshot
				if err := proto.Unmarshal(p, &mds); err == nil {
					if mds.Symbol.Code == "" {
						continue
					}
					resp, err := client.PushTick(context.Background(), &mds)
					if err != nil || resp == nil {
						return
					}
				}
			}
		}
	}
}
