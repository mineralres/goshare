package gateway

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/util"
	// pb "github.com/mineralres/goshare/pkg/pb/goshare"
	"github.com/mineralres/goshare/pkg/api"
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

// Gateway Gateway
type Gateway struct {
	dsList       []api.DataSource          // 常规数据
	wsIndex int64
}

// NewGateway NewGateway
func NewGateway() *Gateway {
	ret := new(Gateway)
	return ret
}

// Run Run
func (g *Gateway) Run(staticDir string, port int) {
	log.Printf("RunTinyGateway on %d", port)
	util.RunTinyGateway(staticDir, port, func(path string, w http.ResponseWriter, r *http.Request) (interface{}, error) {
		switch path {
		case "/api/v1/ws/stream":
			g.handleStream(w, r)
			return nil, errors.New("abort")
		case "/api/v1/md/instrumentList":
			return g.instrumentList(r)
		case "/api/v1/md/mainContract":
			return g.mainContract(r)
		}
		return nil, nil
	})
}
