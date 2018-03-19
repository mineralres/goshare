package base

import "errors"

var (
	// ErrTimeout 超时
	ErrTimeout = errors.New("超时")
	// ErrUnsupported ErrUnsupported
	ErrUnsupported = errors.New("不支持的协议")
)
