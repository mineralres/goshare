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
