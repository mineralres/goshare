package base

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/mineralres/goshare/pkg/pb"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// StringFromGBK 转换GBK
func StringFromGBK(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}

// MakeSymbol 转symbol
func MakeSymbol(s string) pb.Symbol {
	var ret pb.Symbol
	items := strings.Split(s, "_")
	if len(items) == 2 {
		ex := strings.ToUpper(items[0])
		v, ok := pb.ExchangeType_value[ex]
		if ok {
			ret.Exchange = pb.ExchangeType(v)
		}
		ret.Code = items[1]
	}
	return ret
}

// ParseBeijingTime 解析北京时间
func ParseBeijingTime(layout, value string) int64 {
	loc, err := time.LoadLocation("Asia/Chongqing") // 北京时间
	if err == nil {
		tx, err := time.ParseInLocation(layout, value, loc)
		if err == nil {
			return tx.Unix()
		}
		return 0
	}
	tx, err := time.Parse(layout, value)
	if err == nil {
		return (tx.Unix() - 8*3600)
	}
	return 0
}

func Encode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	if err == nil {
		dst = string(data)
	}
	return
}

// Decode 转UTF8
func Decode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}
