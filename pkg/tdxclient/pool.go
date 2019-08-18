package tdxclient

import (
	"errors"
	"log"
	"sort"
	"sync"
	"time"
)

// PoolOptions PoolOptions
type PoolOptions struct {
	ServerList       []string `json:"serverList"`
	ExternServerList []string `json:"externServerList"`
}

// Pool Pool
type Pool struct {
	clients   []*SyncQuoteClient
	exclients []*SyncExternClient
	mu        sync.RWMutex
}

// NewPool NewPool
func NewPool(op *PoolOptions) *Pool {
	p := &Pool{}
	for i := range op.ServerList {
		cl, err := NewSyncQuoteClient(op.ServerList[i], time.Second*3)
		if err != nil {
			log.Println(err)
		}
		p.clients = append(p.clients, cl)
	}
	for i := range op.ExternServerList {
		cl, err := NewSyncExternClient(op.ExternServerList[i], time.Second*3)
		if err != nil {
			log.Println(err)
		}
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
			p.mu.Lock()
			clientCount = len(p.clients)
			for _, c := range p.clients {
				// 相当于心跳包检查
				_, err := c.ReqQryStockCount()
				if err != nil {
					c.ready = false
				}
				if !c.ready {
					failedClientCount++
				}
			}
			exClientCount = len(p.exclients)
			for _, c := range p.exclients {
				// 相当于心跳包
				_, err := c.GetInstrumentCount()
				if err != nil {
					c.ready = false
				}
				if !c.ready {
					failedExClientCount++
				}
			}
			p.mu.Unlock()
		}
		log.Printf("检查连接结果, 普通连接[%d]个,失败[%d]个, 扩展连接[%d]个, 失败[%d]个", clientCount, failedClientCount, exClientCount, failedExClientCount)
	}
}

// GetExternClient get extern client
func (p *Pool) GetExternClient() (*SyncExternClient, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	l := p.exclients
	for _, c := range l {
		if c.ready {
			c.referenceCount++
			sort.Slice(l, func(i, j int) bool {
				return l[i].referenceCount < l[j].referenceCount
			})
			return c, nil
		}
	}
	return nil, errors.New("no ex client valid")
}

// GetQuoteClient get quote client
func (p *Pool) GetQuoteClient() (*SyncQuoteClient, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	l := p.clients
	for _, c := range l {
		if c.ready {
			c.referenceCount++
			sort.Slice(l, func(i, j int) bool {
				return l[i].referenceCount < l[j].referenceCount
			})
			return c, nil
		}
	}
	return nil, errors.New("no ex client valid")
}
