package dcenter

import pb "github.com/mineralres/goshare/pkg/pb/goshare"

// CTPRealtime 用于从CTP订阅实时行情
type CTPRealtime struct {
}

// MakeCTPRealtime MakeCTPRealtime
func MakeCTPRealtime() *CTPRealtime {
	ret := &CTPRealtime{}
	return ret
}

// Subscribe 订阅
func (ctp *CTPRealtime) Subscribe(*pb.ReqSubscribe, chan *pb.MarketDataSnapshot) (*pb.RspSubscribe, error) {
	return nil, nil
}

// UnSubscribe 取消订阅
func (ctp *CTPRealtime) UnSubscribe(*pb.ReqUnSubscribe, chan *pb.MarketDataSnapshot) (*pb.RspUnSubscribe, error) {
	return nil, nil
}
