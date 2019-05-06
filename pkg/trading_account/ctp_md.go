package tradingaccount

import (
	"sync"
	"unsafe"

	"github.com/golang/protobuf/proto"
	"github.com/mineralres/goshare/pkg/pb"
)

// MarketDataSpi CtpMDSpi
type MarketDataSpi interface {
	OnFrontConnected()
	OnRspUserLogin(*pb.RspTradingAccountLogin)
	OnRtnDepthMarketData(*pb.MarketDataSnapshot)
}

// MarketDataAPI CtpMdApi
type MarketDataAPI struct {
	inBuffer  []byte
	outBuffer []byte
	session
	sync.RWMutex
}

func (api *MarketDataAPI) callAPI(callType pb.BindingMessageType, d proto.Message) {
	api.Lock()
	defer api.Unlock()
	api.inType = uint64(callType)
	out, _ := proto.Marshal(d)
	copy(api.inBuffer[:], out[:])
	api.inDataLen = uint64(len(out))
	ctpMDCall.Call(uintptr(unsafe.Pointer(&api.session)))
}

func onMDPopupMessage(s *session) {

	if s.gospi < 0 || s.gospi > 990 {
		panic("onTraderPopupMessage s.gospi < 0 || s.gospi > 990 ")
	}
	data := (*[bufferSize]byte)(s.outData)[:s.outDataLen]
	spi := mdSpiList[s.gospi]
	switch pb.BindingMessageType(s.outType) {
	case pb.BindingMessageType_CTP_ON_FRONT_CONNECTED:
		spi.OnFrontConnected()
	case pb.BindingMessageType_CTP_RSP_USER_LOGIN:
		var rsp pb.RspTradingAccountLogin
		if err := proto.Unmarshal(data, &rsp); err == nil {
			spi.OnRspUserLogin(&rsp)
		}
	case pb.BindingMessageType_CTP_ON_RTN_DEPTH_MARKET_DATA:
		var m pb.MarketDataSnapshot
		if err := proto.Unmarshal(data, &m); err == nil {
			spi.OnRtnDepthMarketData(&m)
		}
	}
}

// MakeMarketDataAPI NewMarketDataAPI
func MakeMarketDataAPI(r *pb.TradingRoute, spi MarketDataSpi) *MarketDataAPI {
	if spi == nil {
		panic("")
	}
	api := &MarketDataAPI{}
	api.goapi = uintptr(unsafe.Pointer(api))
	api.gospi = insertMarketDataSpi(spi)

	api.inBuffer = make([]byte, bufferSize)
	api.outBuffer = make([]byte, bufferSize)
	api.inData = unsafe.Pointer(&api.inBuffer[0])
	api.outData = unsafe.Pointer(&api.outBuffer[0])

	api.callAPI(pb.BindingMessageType_CTP_CREATE_API, r)
	return api
}

// Login 登陆
func (api *MarketDataAPI) Login(ea *pb.TradingAccount) error {
	api.callAPI(pb.BindingMessageType_CTP_REQ_USER_LOGIN, ea)
	return nil
}

// Subscribe 订阅行情
func (api *MarketDataAPI) Subscribe(symbolList *pb.SymbolList) error {
	api.callAPI(pb.BindingMessageType_CTP_REQ_SUBSCRIBE_MARKET_DATA, symbolList)
	return nil
}

// Init 初始化
func (api *MarketDataAPI) Init() error {
	api.callAPI(pb.BindingMessageType_CTP_REQ_CALL_INIT, nil)
	return nil
}
