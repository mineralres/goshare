package httpapi

import (
	"log"
	"net/http"
	"strings"
	"time"

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
		handlerx{"klineSeriesTest", h.klineSeriesTest},
		handlerx{"lastTick", h.lastTick},
		handlerx{"cnStockIndexSummary", h.cnStockIndexSummary},
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
	// log.Println(tag, path)
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
	if k.Open == 0 || k.High == 0 || k.Low == 0 || k.Close == 0 {
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
	var svc goshare.SinaSource
	var filter []pb.Kline
	if req.Period == int(pb.PeriodType_M1) && strings.Index(req.Code, "1000") == 0 {
		req.EndTime = req.StartTime + 1
	}
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
	var svc goshare.SinaSource
	return svc.GetOptionTQuote(req.Month)
}

func transformPeriodType(p string) pb.PeriodType {
	switch p {
	case "604800000":
		return pb.PeriodType_W1
	case "86400000":
		return pb.PeriodType_D1
	case "3600000":
		return pb.PeriodType_H1
	case "1800000":
		return pb.PeriodType_M30
	case "900000":
		return pb.PeriodType_M15
	case "300000":
		return pb.PeriodType_M5
	case "60000":
		return pb.PeriodType_M1
	}
	return pb.PeriodType_D1
}

func (h *HTTPHandler) klineSeriesTest(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	symbol := c.Query("symbol")
	rangex := c.Query("range")
	since := c.Query("since")
	prevTradeTime := c.Query("prevTradeTime")
	// log.Println(symbol, rangex, since, prevTradeTime)

	var ret struct {
		Depths struct {
			Asks [][2]float64 `json:"asks"`
			Bids [][2]float64 `json:"bids"`
		} `json:"depths"`
		Lines  [][6]float64 `json:"lines"`
		Trades []struct {
			Amount float64 `json:"amount"`
			Price  float64 `json:"price"`
			Tid    int64   `json:"tid"`
			Time   int64   `json:"time"`
			Type   string  `json:"type"`
		} `json:"trades"`
	}

	timex := (time.Now().Unix())
	if since != "" {
		timex = int64(base.ParseInt(since))
	}

	var svc goshare.SinaSource
	sx := base.MakeSymbol(symbol)
	l, err := svc.GetKData(&sx, transformPeriodType(rangex), 0, 0, 1)
	if err == nil {
		for i := range l.List {
			k := &l.List[i]
			if k.Time >= timex {
				ret.Lines = append(ret.Lines, [6]float64{float64(k.Time), k.Open, k.High, k.Low, k.Close, k.Volume})
			}
		}
	}

	price := 100.00
	if len(ret.Lines) > 0 {
		price = ret.Lines[len(ret.Lines)-1][4]
	}
	vol := 20.00
	for i := 0; i < 5; i++ {
		ret.Depths.Asks = append(ret.Depths.Asks, [2]float64{price + float64(30-i)*5, vol})
		ret.Depths.Bids = append(ret.Depths.Bids, [2]float64{price - float64(i)*5, vol})
	}

	timex = time.Now().Unix()
	if prevTradeTime != "" {
		timex = int64(base.ParseInt(prevTradeTime)) / 1000
	}

	for i := timex + 1; i <= time.Now().Unix(); i++ {
		ret.Trades = append(ret.Trades, struct {
			Amount float64 `json:"amount"`
			Price  float64 `json:"price"`
			Tid    int64   `json:"tid"`
			Time   int64   `json:"time"`
			Type   string  `json:"type"`
		}{19.79, price, 1585662041877811201, i * 1000, "sell"})
	}

	return &ret, nil
}

func (h *HTTPHandler) lastTick(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var req pb.Symbol
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	var svc goshare.SinaSource
	return svc.GetLastTick(&req)
}

func (h *HTTPHandler) cnStockIndexSummary(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var ret struct {
		// 上证综指
		SSE000001 pb.MarketDataSnapshot
		// 深圳综指
		SZE399001 pb.MarketDataSnapshot
		// 创业板指
		SZE399006 pb.MarketDataSnapshot
	}
	var svc goshare.SinaSource
	ret.SSE000001.Symbol = pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "000001"}
	mds, err := svc.GetLastTick(&ret.SSE000001.Symbol)
	if err == nil {
		ret.SSE000001 = *mds
		ret.SSE000001.Name = "上证综指"
	}
	ret.SZE399001.Symbol = pb.Symbol{Exchange: pb.ExchangeType_SZE, Code: "399001"}
	mds, err = svc.GetLastTick(&ret.SZE399001.Symbol)
	if err == nil {
		ret.SZE399001 = *mds
		ret.SZE399001.Name = "深圳综指"
	}
	ret.SZE399006.Symbol = pb.Symbol{Exchange: pb.ExchangeType_SZE, Code: "399006"}
	mds, err = svc.GetLastTick(&ret.SZE399006.Symbol)
	if err == nil {
		ret.SZE399006 = *mds
		ret.SZE399006.Name = "创业指数"
	}
	return &ret, nil
}
