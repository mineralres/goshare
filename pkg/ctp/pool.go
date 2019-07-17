package ctp

import (
	"errors"
	"fmt"
	"sync"
)

// SyncPool  as pool
type SyncPool struct {
	options *SyncPoolOptions
	aMap    map[string]*SyncAdapter
	mu      sync.Mutex
}

// SyncPoolOptions  options
type SyncPoolOptions struct {
	NewSyncAdapter func(brokerID, account string) (*SyncAdapter, error)
}

// NewSyncPool create new pool
func NewSyncPool(options *SyncPoolOptions) *SyncPool {
	if options.NewSyncAdapter == nil {
		panic("NewSyncAdapter should not be nil")
	}
	ret := &SyncPool{options: options}
	ret.aMap = make(map[string]*SyncAdapter)
	return ret
}

// GetAdapter get adapter from pool
func (p *SyncPool) GetAdapter(brokerID, account string) (*SyncAdapter, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	key := fmt.Sprintf("%s-%s", brokerID, account)
	a, ok := p.aMap[key]
	if ok && !a.adapter.closed {
		return a, nil
	}
	var err error
	a, err = p.options.NewSyncAdapter(brokerID, account)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, errors.New("Create sync adapter failed")
	}
	p.aMap[key] = a
	return a, nil
}
