package base

import (
	"bytes"
	"io/ioutil"
	"strings"

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
