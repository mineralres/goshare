package goshare

import (
	"github.com/mineralres/goshare/pkg/history"
	"github.com/mineralres/goshare/pkg/realtime"
)

// Service Service
type Service struct {
	history.HisProvider
	realtime.RProvider
}
