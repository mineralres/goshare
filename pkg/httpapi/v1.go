package httpapi

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mineralres/goshare/pkg/base"
	"github.com/mineralres/goshare/pkg/goshare"
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
	log.Println(tag, path)
	var hl []handlerx
	if tag == "gosharev1" {
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

func validKline(k *pb.Kline) bool {
	pmax := 99999999.99
	if k.Time == 0 {
		return false
	}
	if k.Open > pmax || k.Open < 0 {
		return false
	}
	if k.High > pmax || k.High < 0 {
		return false
	}
	if k.Low > pmax || k.Low < 0 {
		return false
	}
	if k.Close > pmax || k.Close < 0 {
		return false
	}
	return true
}

func (h *HTTPHandler) klineSeries(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var req struct {
		Exchange  int    `json:"exchange"`
		Code      string `json:"code"`
		Period    int    `json:"period"`
		StartTime int64  `json:"startTime"`
		EndTime   int64  `json:"endTime"`
	}
	var err error
	err = c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	var svc goshare.Service
	var filter []pb.Kline
	ret, err := svc.GetKData(&pb.Symbol{Exchange: pb.ExchangeType(req.Exchange), Code: req.Code}, pb.PeriodType(req.Period), req.StartTime, req.EndTime, 1)
	for i := range ret.List {
		k := &ret.List[i]
		if validKline(k) {
			filter = append(filter, *k)
		}
	}
	ret.List = filter
	return ret, err
}

func (h *HTTPHandler) sseOptionTQuote(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var req struct {
		Month string `json:"month"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	var svc goshare.Service
	return svc.GetOptionSinaTick(req.Month)
}
