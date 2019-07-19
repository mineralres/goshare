package spider

import (
	pb "github.com/mineralres/goshare/pkg/pb/goshare"
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
func (s *Spider) GetTick(ex, symbol string) (*pb.MarketDataSnapshot, error) {
	return s.GetLastTick(ex, symbol)
}

// GetLastTickSerires 取最新一天的tick序列
func (s *Spider) GetLastTickSerires(ex, symbol string) (*pb.TickSeries, error) {
	return nil, nil
}

// GetTradingInstrument 读取合约信息
func (s *Spider) GetTradingInstrument(ex, symbol string) (*pb.Instrument, error) {
	return nil, nil
}

// TradingInstrumentList 全部合约信息
func (s *Spider) TradingInstrumentList(*pb.ReqGetTradingInstrumentList) ([]*pb.Instrument, error) {
	return nil, nil
}

// RGetKlineSeries 反向取K线
func (s *Spider) RGetKlineSeries(ex, symbol string, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error) {
	return nil, nil
}

// GetKlineSeries K线
func (s *Spider) GetKlineSeries(ex, symbol string, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error) {
	return nil, nil
}
