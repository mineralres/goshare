package spider

import (
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/service/dcenter"
)

// Spider Spider
type Spider struct {
}

// MakeSpider MakeSpider
func MakeSpider() *Spider {
	ret := &Spider{}
	return ret
}

// GetTick 读取tick
func (s *Spider) GetTick(ctx *dcenter.DataSourceContext, req *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	return s.GetLastTick(req)
}

// GetLastTickSerires 取最新一天的tick序列
func (s *Spider) GetLastTickSerires(*dcenter.DataSourceContext, *pb.Symbol) (*pb.TickSeries, error) {
	return nil, nil
}

// GetTradingInstrument 读取合约信息
func (s *Spider) GetTradingInstrument(*dcenter.DataSourceContext, *pb.Symbol) (*pb.TradingInstrument, error) {
	return nil, nil
}

// TradingInstrumentList 全部合约信息
func (s *Spider) TradingInstrumentList(*dcenter.DataSourceContext, *pb.ReqGetTradingInstrumentList) ([]*pb.TradingInstrument, error) {
	return nil, nil
}

// RGetKlineSeries 反向取K线
func (s *Spider) RGetKlineSeries(ctx *dcenter.DataSourceContext, symbol *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error) {
	return nil, nil
}

// GetKlineSeries K线
func (s *Spider) GetKlineSeries(ctx *dcenter.DataSourceContext, symbol *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error) {
	return nil, nil
}
