package tradingaccount

import (
	"os"
	"sync"
	"unsafe"

	"github.com/golang/protobuf/proto"
	"github.com/mineralres/goshare/pkg/pb"
)

// Trader api
type Trader struct {
	inBuffer  []byte
	outBuffer []byte
	session
	sync.RWMutex
}

func (trader *Trader) callAPI(callType pb.BindingMessageType, d proto.Message) {
	// 因为会修改inBuffer，所以需要锁
	trader.Lock()
	defer trader.Unlock()
	trader.inType = uint64(callType)
	if d != nil {
		out, _ := proto.Marshal(d)
		copy(trader.inBuffer[:], out[:])
		trader.inDataLen = uint64(len(out))
	}
	// log.Println("ctp call ", callType)
	ctpTradeCall.Call(uintptr(unsafe.Pointer(&trader.session)))
}

func onTraderPopupMessage(s *session) {
	if s.gospi < 0 || s.gospi > 990 {
		panic("onTraderPopupMessage s.gospi < 0 or > 990 ")
	}
	data := (*[bufferSize]byte)(s.outData)[:s.outDataLen]
	// s.gospi 是index
	spi := HandlerList[s.gospi]
	// log.Println("pb.BindingMessageType(s.outType)", pb.BindingMessageType(s.outType))
	switch pb.BindingMessageType(s.outType) {
	case pb.BindingMessageType_CTP_ON_FRONT_CONNECTED:
		spi.OnFrontConnected()
	case pb.BindingMessageType_CTP_RSP_USER_LOGIN:
		var rsp pb.CTPRspInfo
		if err := proto.Unmarshal(data, &rsp); err == nil {
			spi.OnRspUserLogin(&rsp)
		}
	case pb.BindingMessageType_CTP_ON_RTN_ORDER:
		var order pb.Order
		if err := proto.Unmarshal(data, &order); err == nil {
			spi.OnRtnOrder(&order)
		}
	case pb.BindingMessageType_CTP_ON_RTN_TRADE:
		var trade pb.TradeReport
		if err := proto.Unmarshal(data, &trade); err == nil {
			spi.OnRtnTrade(&trade)
		}
	case pb.BindingMessageType_CTP_ON_RSP_ORDER_INSERT:
		var rsp pb.CTPOnRspOrderInsert
		if err := proto.Unmarshal(data, &rsp); err == nil {
			spi.OnRspOrderInsert(&rsp)
		}
	}
}

// MakeTrader 创建接口, 返回接口ID
func MakeTrader(r *pb.TradingRoute, spi Handler) *Trader {
	if spi == nil {
		panic("")
	}
	trader := &Trader{}
	trader.goapi = uintptr(unsafe.Pointer(trader))
	trader.gospi = insertHandler(spi)
	// log.Println("trader.gospi", trader.gospi, unsafe.Sizeof(trader.session))

	trader.inBuffer = make([]byte, bufferSize)
	trader.outBuffer = make([]byte, bufferSize)
	trader.inData = unsafe.Pointer(&trader.inBuffer[0])
	trader.outData = unsafe.Pointer(&trader.outBuffer[0])

	trader.callAPI(pb.BindingMessageType_CTP_CREATE_API, r)
	return trader
}

// Destory 销毁接口
func (trader *Trader) Destory() {
	trader.callAPI(pb.BindingMessageType_CTP_DELETE_API, nil)
}

// Login 登陆
func (trader *Trader) Login(ea *pb.TradingAccount) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_USER_LOGIN, ea)
	return nil
}

// InsertOrder 发单
func (trader *Trader) InsertOrder(order *pb.Order) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_INSERT_ORDER, order)
	return nil
}

// CancelOrder 撤单
func (trader *Trader) CancelOrder(req *pb.CancelOrderRequest) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_CANCEL_ORDER, req)
	return nil
}

// QuerySummary 查询资金
func (trader *Trader) QuerySummary() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QUERY_TRADING_ACCOUNT, nil)
	return nil
}

// QueryPositionDetail 查询持仓明细
func (trader *Trader) QueryPositionDetail() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QUERY_POSITION_DETAIL, nil)
	return nil
}

// QueryCommissionRate 查询手续费率
func (trader *Trader) QueryCommissionRate(s *pb.Symbol) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QUERY_COMMISSION_RATE, s)
	return nil
}

// QueryMarginRate 查询保证金率
func (trader *Trader) QueryMarginRate(s *pb.Symbol) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QUERY_MARGIN_RATE, s)
	return nil
}

// Connect 连接
func (trader *Trader) Connect(req *pb.CtpReqConnect) error {
	var err error
	req.Directory, err = os.Getwd()
	trader.callAPI(pb.BindingMessageType_CTP_REQ_CONNECT, req)
	return err
}

// Authencate 认证
func (trader *Trader) Authencate(req *pb.CtpReqAuthencate) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_AUTHENTICATE, req)
	return nil
}

// LoginOnly 登陆
func (trader *Trader) LoginOnly(req *pb.CtpReqUserLogin) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_USER_LOGIN, req)
	return nil
}

// ConfirmSettlementInfo 确认结算单
func (trader *Trader) ConfirmSettlementInfo(req *pb.CtpReqSettlementInfoConfirm) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_SETTLEMENT_INFO_CONFIRM, req)
	return nil
}

// QueryInvestor 查询投资者
func (trader *Trader) QueryInvestor(req *pb.CtpReqQryInvestor) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QUERY_INVESTOR, req)
	return nil
}

// QueryTransferBank 查询银行
func (trader *Trader) QueryTransferBank(req *pb.CtpReqQryTransferBank) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QUERY_TRANSFER_BANK, req)
	return nil
}

// Transfer 银行转期货
func (trader *Trader) Transfer(req *pb.CtpReqTransfer) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_TRANSFER, req)
	return nil
}

// QueryAccountRegister 查询注册资金账号
func (trader *Trader) QueryAccountRegister(req *pb.CtpReqQryAccountRegister) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QUERY_ACCOUNT_REGISTER, req)
	return nil
}

// Init 初始化
func (trader *Trader) Init() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_CALL_INIT, nil)
	return nil
}

// Disconnect 断开连接
func (trader *Trader) Disconnect() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_DISCONNECT, nil)
	return nil
}
