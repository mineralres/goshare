package httpapi

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mineralres/goshare/pkg/base"
	"github.com/mineralres/goshare/pkg/pb"
)

type handlerx struct {
	path    string
	handler func(*gin.Context, *pb.UserSession) (interface{}, error)
}

// HTTPHandler HTTPHandler
type HTTPHandler struct {
	handlerList1 []handlerx
}

// Run Run works
func (h *HTTPHandler) Run(port string) {
	h.registerHandler()
	r := gin.New()
	r.Use(h.httpHook)
	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	s.SetKeepAlivesEnabled(false)
	log.Printf("HTTP serve on %s ", port)
	s.ListenAndServe()
}

func (h *HTTPHandler) registerHandler() {
	h.handlerList1 = []handlerx{
		handlerx{"klineSeries", h.klineSeries},
		handlerx{"sseOptionTQuote", h.sseOptionTQuote},
		handlerx{"lastTick", h.lastTick},
		handlerx{"cnStockIndexSummary", h.cnStockIndexSummary},
		handlerx{"apiTest", h.apiTest},
		handlerx{"sseStockOptionList", h.sseStockOptionList},
	}
}

func (h *HTTPHandler) httpHook(context *gin.Context) {
	pathItems := strings.Split(context.Request.RequestURI, "/")
	if len(pathItems) < 3 {
		res := &base.HTTPResponse{Success: false}
		context.JSON(404, res)
		return
	}

	tag := pathItems[1]
	path := pathItems[2]
	indexx := strings.Index(path, "?")
	if indexx > 0 {
		path = path[0:indexx]
	}

	res := &base.HTTPResponse{}
	log.Println("访问链接", tag, path)
	var hl []handlerx
	if tag == "v1" {
		hl = h.handlerList1
	}
	err := base.Err404
	var rd interface{}
	for i := range hl {
		h := &hl[i]
		if h.path == path {
			rd, err = h.handler(context, nil)
		}
	}
	if err == nil {
		res.Success = true
	} else {
		res.Success = false
		res.Msg = err.Error()
	}
	res.Data = rd
	if err == base.Err404 {
		context.JSON(404, res)
		log.Println("404 Not found ", context.Request.RequestURI, tag, path)
	} else if err == base.ErrAbort {

	} else {
		context.JSON(200, res)
	}

}
