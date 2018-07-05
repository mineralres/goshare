package base

import (
	"strconv"
)

// ParseInt ParseInt
func ParseInt(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return i
}

// ParseFloat ParseFloat
func ParseFloat(src string) float64 {
	f, err := strconv.ParseFloat(src, 10)
	if err != nil {
		return 0
	}
	return f
}

// HTTPResponse 回复
type HTTPResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Code    int32       `json:"code"`
}
