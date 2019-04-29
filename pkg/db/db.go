package db

import "github.com/mineralres/goshare/pkg/pb"

// XDataBase 数据存储引擎
type XDataBase interface {
	// 保存单根K线
	Save(*pb.Symbol, pb.PeriodType, *pb.Kline) error
	// 保存日tick序列
	SaveDayTickSeries(ts *pb.TickSeries) error
	// 反向取K线
	RGetKlineSeries(s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries
	// 取最新一天的tick序列
	GetLastTickSerires(s *pb.Symbol) *pb.TickSeries
	// 取独立ID
	GetUniqueID() (int64, error)
	// 保存user session
	SetUserSession(string, *pb.UserSession) error
	// 读取user session
	GetUserSession(string) (*pb.UserSession, error)
}
