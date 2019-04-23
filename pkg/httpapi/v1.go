package httpapi

import (
	"github.com/gin-gonic/gin"
	"github.com/mineralres/goshare/pkg/goshare"
	"github.com/mineralres/goshare/pkg/pb"
)

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
	var filter []pb.Kline
	return filter, err
}

func (h *HTTPHandler) sseOptionTQuote(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var req struct {
		Month string `json:"month"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	var svc goshare.DataSource
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

func (h *HTTPHandler) lastTick(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var req pb.Symbol
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}
	var svc goshare.DataSource
	return svc.GetLastTick(&req)
}

func (h *HTTPHandler) apiTest(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var svc goshare.DataSource
	return svc.GetLastTick(&pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "601398"})
}

func (h *HTTPHandler) cnStockIndexSummary(c *gin.Context, s *pb.UserSession) (interface{}, error) {
	var ret struct {
		// 上证综指
		SSE000001 pb.MarketDataSnapshot
		// 深圳综指
		SZE399001 pb.MarketDataSnapshot
		// 创业板指
		SZE399006 pb.MarketDataSnapshot
		// 上证50指数
		SSE000016 pb.MarketDataSnapshot
	}
	var svc goshare.DataSource
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
	ret.SSE000016.Symbol = pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "000016"}
	mds, err = svc.GetLastTick(&ret.SSE000016.Symbol)
	if err == nil {
		ret.SSE000016 = *mds
		ret.SSE000016.Name = "上证50"
	}
	return &ret, nil
}
