package tdxclient

import (
	"log"
	"sync"
	"time"
)

// PoolOptions PoolOptions
type PoolOptions struct {
	ServerList       []string `json:"serverList"`
	ExternServerList []string `json:"externServerList"`
}

// Pool Pool
// 暂时都采用同步连接
type Pool struct {
	clients        []*SyncQuoteClient
	clientsIndex   int
	clientsLock    sync.RWMutex
	exclients      []*SyncExternClient
	exclientsIndex int
	exclientsLock  sync.RWMutex
	requestID      int64
}

// NewPool NewPool
func NewPool(op *PoolOptions) *Pool {
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

// GetExternClient get extern client
func (p *Pool) GetExternClient() *SyncExternClient {
	p.exclientsLock.RLock()
	defer p.exclientsLock.RUnlock()
	if len(p.exclients) == 0 {
		return nil
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
