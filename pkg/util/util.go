package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"strconv"
	"time"

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

// StringFromGBK2 StringFromGBK2
func StringFromGBK2(src []byte) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
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

// ParseInt ParseInt
func ParseInt(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return i
}

// ParseInt32 ParseInt
func ParseInt32(src string) int32 {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return int32(i)
}

// ParseFloat ParseFloat
func ParseFloat(src string) float64 {
	f, err := strconv.ParseFloat(src, 10)
	if err != nil {
		return 0
	}
	return f
}

// GetMD5 转md5
func GetMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
