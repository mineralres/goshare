package tdxclient

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/service/dcenter"
)

// PoolOptions PoolOptions
type PoolOptions struct {
	ServerList       []string `json:"serverList"`
	ExternServerList []string `json:"externServerList"`
}

// Pool Pool
// 暂时都采用同步连接
type Pool struct {
	clients        []*SyncClient
	clientsIndex   int
	clientsLock    sync.RWMutex
	exclients      []*SyncExternClient
	exclientsIndex int
	exclientsLock  sync.RWMutex
	requestID      int64
}

// MakePool MakePool
func MakePool(op *PoolOptions) *Pool {
	p := &Pool{}
	for i := range op.ServerList {
		cl := makeSyncClient(op.ServerList[i])
		cl.init()
		p.clients = append(p.clients, cl)
	}
	for i := range op.ExternServerList {
		cl := makeSyncExternClient(op.ExternServerList[i])
		cl.init()
		p.exclients = append(p.exclients, cl)
	}
	go p.checker()
	return p
}

func (p *Pool) checker() {
	timer := time.NewTicker(time.Second * 60)
	for {
		var clientCount, exClientCount, failedClientCount, failedExClientCount int
		select {
		case <-timer.C:
			p.clientsLock.Lock()
			clientCount = len(p.clients)
			for i := range p.clients {
				if !p.clients[i].ready {
					failedClientCount++
				}
			}
			p.clientsLock.Unlock()
			p.exclientsLock.Lock()
			exClientCount = len(p.exclients)
			for i := range p.exclients {
				p.exclients[i].GetInstrumentCount()
				if !p.exclients[i].ready {
					failedExClientCount++
				}
			}
			p.exclientsLock.Unlock()
		}
		log.Printf("检查连接结果, 普通连接[%d]个,失败[%d]个, 扩展连接[%d]个, 失败[%d]个", clientCount, failedClientCount, exClientCount, failedExClientCount)
	}
}

func (p *Pool) getExternClient() *SyncExternClient {
	p.exclientsLock.RLock()
	defer p.exclientsLock.RUnlock()
	if len(p.exclients) == 0 {
		panic("no extern server provided")
	}
	var ret *SyncExternClient
	j := p.exclientsIndex + 1
	for {
		if j == p.exclientsIndex-1 {
			break
		}
		if j > len(p.exclients)-1 {
			j = 0
		}
		ret = p.exclients[j]
		if ret.ready {
			break
		}
		j++
	}
	p.exclientsIndex = j
	return ret
}

// GetTick GetTick
func (p *Pool) GetTick(ctx *dcenter.DataSourceContext, s *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	var ret pb.MarketDataSnapshot
	if (s.Exchange >= 0 && s.Exchange <= pb.ExchangeType_CFFEX) || s.Exchange <= pb.ExchangeType_INE {
		// 期货行情还是使用自己接收的行情字段比较完整
		return &ret, errors.New("tdxunsupported")
		// 扩展行情
		cl := p.getExternClient()
		if cl == nil {
			panic("cl==nil")
		}
		req := &ReqGetInstrumentQuote{Code: s.Code}
		req.Code = strings.ToUpper(req.Code)
		switch s.Exchange {
		case pb.ExchangeType_CZCE:
			req.Market = 28
			req.Code = req.Code[:2] + "1" + req.Code[2:]
		case pb.ExchangeType_DCE:
			req.Market = 29
		case pb.ExchangeType_SHFE, pb.ExchangeType_INE:
			req.Market = 30
		case pb.ExchangeType_CFFEX:
			req.Market = 47
		}
		md, err := cl.GetInstrumentQuote(req)
		if err != nil {
			log.Println("err", err, s)
			return &ret, err
		}
		*md.Symbol = *s
		return md, err
	} else if s.Exchange >= pb.ExchangeType_SSE || s.Exchange >= pb.ExchangeType_SZE {
		// 股票行情，普通行情接口
	}
	return &ret, errors.New("tdxunsupported")
}

// GetLastTickSerires GetLastTickSerires
func (p *Pool) GetLastTickSerires(*dcenter.DataSourceContext, *pb.Symbol) (*pb.TickSeries, error) {
	// log.Println("TDXPOOL GetLastTickSerires")
	return nil, errors.New("tdxunsupported")
}

// GetTradingInstrument GetTradingInstrument
func (p *Pool) GetTradingInstrument(*dcenter.DataSourceContext, *pb.Symbol) (*pb.TradingInstrument, error) {
	// log.Println("TDXPOOL GetTradingInstrument")
	return nil, errors.New("tdxunsupported")
}

// TradingInstrumentList TradingInstrumentList
func (p *Pool) TradingInstrumentList(*dcenter.DataSourceContext, *pb.ReqGetTradingInstrumentList) ([]*pb.TradingInstrument, error) {
	log.Println("TDXPOOL TradingInstrumentList")
	return nil, errors.New("tdxunsupported")
}

// RGetKlineSeries RGetKlineSeries
func (p *Pool) RGetKlineSeries(ctx *dcenter.DataSourceContext, s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error) {
	// log.Println("TDXPOOL RGetKlineSeries")
	return nil, errors.New("tdxunsupported")
}

// GetKlineSeries GetKlineSeries
func (p *Pool) GetKlineSeries(ctx *dcenter.DataSourceContext, s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error) {
	// log.Println("TDXPOOL GetKlineSeries")
	return nil, errors.New("tdxunsupported")
}
