package ctp

import (
	"fmt"
	"sync"
	"time"
)

// SyncPool  as pool
type SyncPool struct {
	options *SyncPoolOptions
	aMap    map[string]*SyncAdapter
	mu      sync.Mutex
}

// SyncPoolOptions  options
type SyncPoolOptions struct {
	IdleTimeout time.Duration
	GatewayHost string
	NewSyncAdapter  func(brokerID, account string) *SyncAdapter
}

// NewSyncPool create new pool
func NewSyncPool(options *SyncPoolOptions) *SyncPool {
	if options.NewSyncAdapter == nil {
		panic("NewSyncAdapter should not be nil")
	}
	ret := &SyncPool{options: options}
	return ret
}

// GetAdapter get adapter from pool
func (p *SyncPool) GetAdapter(brokerID, account string) *SyncAdapter {
	p.mu.Lock()
	defer p.mu.Unlock()
	key := fmt.Sprintf("%s-%s", brokerID, account)
	a, ok := p.aMap[key]
	if ok {
		return a
	}
	a = p.options.NewSyncAdapter(brokerID, account)
	p.aMap[key] = a
	return a
}
