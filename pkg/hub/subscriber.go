package hub

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/mineralres/goshare/pkg/pb/ctp"
)

// Subscriber  subscribe tick
// 订阅，退订，和断线重连(重订)
type Subscriber struct {
	adapter       *Adapter
	mapInstrument map[string]string
	lastTickTime  int64
	mu            sync.Mutex
}

// NewSubscriber create subscriber
func NewSubscriber(host, account, password string, fronts []string, timeout time.Duration, onTick func(*ctp.CThostFtdcDepthMarketDataField)) (*Subscriber, error) {
	subscriber := &Subscriber{}
	subscriber.mapInstrument = make(map[string]string)
	xontick := func(rtn *ctp.CThostFtdcDepthMarketDataField) {
		subscriber.mu.Lock()
		ex, ok := subscriber.mapInstrument[rtn.InstrumentID]
		if ok {
			rtn.ExchangeID = ex
		}
		subscriber.lastTickTime = time.Now().Unix()
		subscriber.mu.Unlock()
		onTick(rtn)
	}
	adapter, err := createAdapter(host, account, password, fronts, timeout, xontick)
	subscriber.lastTickTime = time.Now().Unix()
	go func() {
		for {
			if !subscriber.isOK() {
				adapter, err = createAdapter(host, account, password, fronts, timeout, xontick)
				if err == nil {
					subscriber.mu.Lock()
					if subscriber.adapter != nil {
						subscriber.adapter.Close()
					}
					subscriber.adapter = adapter
					subscriber.lastTickTime = time.Now().Unix()
					m2 := subscriber.mapInstrument
					subscriber.mu.Unlock()
					// 重新订阅
					for symbol, exchange := range m2 {
						subscriber.Subscribe(exchange, symbol)
					}
				}
			}
			time.Sleep(time.Second * 5)
		}
	}()
	subscriber.adapter = adapter
	return subscriber, err
}

func createAdapter(host, account, password string, fronts []string, timeout time.Duration, onTick func(*ctp.CThostFtdcDepthMarketDataField)) (*Adapter, error) {
	var adapter *Adapter
	var err error
	sig := make(chan bool)
	adapter, err = NewAdapter(host, timeout, func(pkt *Packet) {
		switch ctp.CtpMessageType(pkt.MsgType) {
		case ctp.CtpMessageType_MD_OnRspSubMarketData:
		case ctp.CtpMessageType_MD_OnRtnDepthMarketData:
			var rtn ctp.CThostFtdcDepthMarketDataField
			if err := pkt.Get1(&rtn); err == nil {
				onTick(&rtn)
			}
		case ctp.CtpMessageType_MD_OnFrontConnected:
			var req ctp.CThostFtdcReqUserLoginField
			req.UserID = account
			req.Password = password
			adapter.Post(int32(ctp.CtpMessageType_MD_ReqUserLogin), &req, 1)
		case ctp.CtpMessageType_MD_OnRspUserLogin:
			var rsp ctp.CThostFtdcRspUserLoginField
			var rspInfo ctp.CThostFtdcRspInfoField
			if err := pkt.Get2(&rsp, &rspInfo); err == nil {
			}
			sig <- true
		default:
			log.Println(ctp.CtpMessageType(pkt.MsgType), len(pkt.BodyList))
		}
	})
	var req ctp.CThostFtdcReqRegisterFrontField
	req.Fronts = fronts
	adapter.Post(int32(ctp.CtpMessageType_MD_RegisterFront), &req, 1)
	adapter.Post(int32(ctp.CtpMessageType_MD_Init), nil, 2)
	select {
	case <-time.After(timeout):
		adapter.Close()
		return nil, errors.New("CreateSubscriberTimeout")
	case <-sig:
		break
	}
	return adapter, err
}

func (s *Subscriber) isOK() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return !(s.adapter == nil || s.adapter.closed || time.Now().Unix()-s.lastTickTime > 60)
}

// Subscribe quote
func (s *Subscriber) Subscribe(exchange, symbol string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mapInstrument[symbol] = exchange
	var req ctp.CThostFtdcReqSubscribeMarketData
	req.Instruments = append(req.Instruments, symbol)
	s.adapter.Post(int32(ctp.CtpMessageType_MD_SubscribeMarketData), &req, 2)
	return nil
}

// UnSubscribe un sub
func (s *Subscriber) UnSubscribe(exchange, symbol string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.mapInstrument, symbol)
	return nil
}
