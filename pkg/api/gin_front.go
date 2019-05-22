package api

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/service"
)

func pbJSON(c *gin.Context, o proto.Message) error {
	header := c.Writer.Header()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
	m := jsonpb.Marshaler{EmitDefaults: true}
	err := m.Marshal(c.Writer, o)
	if err != nil {
		return err
	}
	c.Writer.WriteHeader(200)
	return nil
}

var (
	wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// XRouter XRouter
type XRouter struct {
	Router *gin.Engine
	ucc    pb.UCenterClient
	dcc    pb.DCenterClient
}

// Authentication 鉴权
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// MakeAPIRouter API
func MakeAPIRouter() *XRouter {
	var xr XRouter
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	xr.Router = r
	conn, err := service.GetClientConn("srv.ucenter")
	if err != nil {
		panic(err)
	}
	xr.ucc = pb.NewUCenterClient(conn)
	conn, err = service.GetClientConn("srv.dcenter")
	if err != nil {
		panic(err)
	}
	xr.dcc = pb.NewDCenterClient(conn)
	r.Use(gin.Recovery())
	r.Use(Authentication())
	// UI静态文件
	r.Use(static.Serve("/", static.LocalFile("./ui-release", true)))

	// 注册API
	g := r.Group("/api/v1")
	g.GET("/user", xr.user)
	g.GET("/user/routes", xr.routes)
	g.GET("/users", xr.userList)
	g.GET("/strategies", xr.strategies)
	g.GET("/dcenter/info", xr.dcenterInfo)
	g.POST("/dcenter/getTradingInstrument", xr.getTradingInstrument)

	g = r.Group("/ws/v1/")
	g.GET("/dcenter/subscribe", xr.subscribe)
	return &xr
}
