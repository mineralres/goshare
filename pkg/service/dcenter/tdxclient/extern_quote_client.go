package tdxclient

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/mineralres/goshare/pkg/util"
)

var (
	exGetMarketsCmd            = []byte{0x01, 0x02, 0x48, 0x69, 0x00, 0x01, 0x02, 0x00, 0x02, 0x00, 0xf4, 0x23} // 查询市场列表
	exGetInstrumentBarCountCmd = []byte{0x01, 0x03, 0x48, 0x66, 0x00, 0x01, 0x02, 0x00, 0x02, 0x00, 0xf0, 0x23} // 查询合约数量
	exSetupCmd1                = []byte{0x01, 0x01, 0x48, 0x65, 0x00, 0x01, 0x52, 0x00, 0x52, 0x00, 0x54, 0x24, // 初始化连接
		0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32,
		0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5,
		0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d,
		0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0xcc, 0xe1, 0x6d, 0xff, 0xd5, 0xba, 0x3f, 0xb8,
		0xcb, 0xc5, 0x7a, 0x05, 0x4f, 0x77, 0x48, 0xea}
)

type requestType struct {
	klineType   TdxKlineType
	requestID   int64
	handlerFunc interface{}
}

type tdxResponse struct {
	h    header
	body []byte
}

// ExternQuoteClient 行情接口
// 因为没有测试出返回数据包里requestID对应的字段，所以请求与响应之间不好做到一一对应，因为异步模式暂时不启用
// 使用tdx接口初衷也不是速度，而是基于数据质量和方便性考虑的
type ExternQuoteClient struct {
	chOut            chan []byte
	requestQueue     []requestType
	requestQueueLock sync.RWMutex
	ready            bool
	retryTimes       int
	server           string
}

// Ready 是否准备好
func (c *ExternQuoteClient) Ready() bool {
	return c.ready
}

// Server Server
func (c *ExternQuoteClient) Server() string {
	return c.server
}

// ConnectAndJoin 连接
func (c *ExternQuoteClient) ConnectAndJoin(address string, hf func(), autoReconnect bool) {
	if c.retryTimes > 20 {
		log.Printf("[%s] 经过[%d]次重试不成功停止", address, c.retryTimes)
		return
	}
	c.server = address
	c.requestQueueLock.Lock()
	c.requestQueue = c.requestQueue[:0]
	c.requestQueueLock.Unlock()
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("[%s] 连接失败 [%v]", address, err)
		return
	}
	defer func() {
		c.ready = false
		conn.Close()
		log.Printf("通达信连接断开 [%t]", autoReconnect)
		if autoReconnect {
			time.Sleep(time.Second)
			c.retryTimes++
			go c.ConnectAndJoin(address, hf, autoReconnect)
		}
	}()
	c.chOut = make(chan []byte, 100)
	timer := time.NewTicker(time.Second * 60)
	// 启动写
	go func() {
		for {
			select {
			case <-timer.C:
				// 发送心跳包 用查询股票数量当心跳
				c.GetInstrumentCount()
			case d := <-c.chOut:
				if d == nil {
					continue
				}
				if len(d) == 0 {
					// 关闭连接
					conn.Close()
				}
				n, err := conn.Write(d)
				if err != nil {
					conn.Close()
					return
				}
				// log.Println("发送成功", n, err, d)
				if n < len(d) {
					// panic("n < len(d)")
					return
				}
			}
		}
	}()
	c.setup()
	c.ready = true
	hf()
	// 启动读
	c.read(conn)
}

// Close 关机
func (c *ExternQuoteClient) Close() {
	c.chOut <- make([]byte, 0)
}

func (c *ExternQuoteClient) setup() {
	c.chOut <- exSetupCmd1
}

func (c *ExternQuoteClient) read(conn net.Conn) {
	defer conn.Close()
	for {
		headerBuf := make([]byte, 16)
		n, err := conn.Read(headerBuf)
		if err != nil {
			log.Println(err)
			break
		}
		// log.Println(n, err)
		if n != len(headerBuf) {
			log.Println(err)
			continue
		}
		var h header
		err = unmarshal(headerBuf, &h)
		if err != nil {
			panic(err)
		}
		body := make([]byte, h.ZipSize)
		var readed uint16
		for {
			n, err := conn.Read(body[readed:])
			if err != nil {
				break
			}
			readed += uint16(n)
			if readed >= h.ZipSize {
				break
			}
		}
		if h.ZipSize == h.UnzipSize {
			// log.Println("不需要解压")
		} else {
			// log.Println("需要解压")
			r, err := zlib.NewReader(bytes.NewReader(body))
			if err != nil {
				log.Println(err, len(body), h, readed)
				panic(err)
			}
			unzipBuf := &bytes.Buffer{}
			io.Copy(unzipBuf, r)
			r.Close()
			if unzipBuf.Len() != int(h.UnzipSize) {
				log.Printf("unzipBuf.Len(%d) != h.unzipSize(%d)", unzipBuf.Len(), h.UnzipSize)
				panic("unzipBuf.Len() != h.unzipSize")
			}
			body = unzipBuf.Bytes()
			// log.Printf("解压成功 srclen[%d] destlen[%d]", len(body), unzipBuf.Len())
		}
		c.handleMessage(&h, body)
	}
}

func (c *ExternQuoteClient) send(pkg []byte) {
	c.chOut <- pkg
}

func (c *ExternQuoteClient) handleMessage(h *header, body []byte) {
	b, _ := json.Marshal(h)
	log.Printf("收到消息 %s len(body)=%d", string(b), len(body))
	switch h.F2 {
	case 0x66480301:
		// 股票数量
		var rsp RspGetInstrumentCount
		err := unmarshal(body[19:], &rsp)
		if err != nil {
			panic(err)
		}
		log.Printf("查询合约数量[%d]", rsp.Count)
	case 1766326801:
		var num uint16
		reader := bytes.NewReader(body)
		binary.Read(reader, binary.LittleEndian, &num)
		pos := 2
		for i := 0; i < int(num); i++ {
			var item ExtRspQryMarket
			err := unmarshal(body[pos:], &item)
			item.Name = util.Decode(item.Name)
			item.ShortName = util.Decode(item.ShortName)
			if err != nil {
				panic(err)
			}
			log.Println("扩展行情查询市场", item, num, i)
			pos += 64
		}
	case 1778909457:
		c.onRspGetInstrumentBars(h, body)
	case 1778909441:
		log.Println("查询失败", string(body))
		// 查询失败.
		c.requestQueueLock.Lock()
		if len(c.requestQueue) == 0 {
			panic("len(c.requestQueue) == 0")
		}
		req := c.requestQueue[0]
		c.requestQueue = c.requestQueue[1:]
		c.requestQueueLock.Unlock()
		if req.handlerFunc != nil {
			h := req.handlerFunc.(func([]TdxKline, int64))
			h(nil, req.requestID)
		}
	case 34078993:
		c.onRspGetInstrumentQuote(h, body)
	case 1732772881:
		c.onRspGetInstrumentInfo(h, body)
	case 526097:
		if h.F3 == 604700673 {
			c.onRspGetMinuteTimeData(h, body)
		}
	case 3146001:
		if h.F3 == 604766209 {
			// 历史分时
			c.onRspGetHistoryMinuteTimeData(h, body)
		}
	case 524561:
		c.onRspGetTransactionData(h, body)
	case 3146257:
		c.onRspGetHistoryTransactionData(h, body)
	}
}

// GetInstrumentCount 查询股票数量
func (c *ExternQuoteClient) GetInstrumentCount() {
	// 前四个字段等于F2 小端
	// cmd := []byte{0x01, 0x03, 0x48, 0x66, 0x00, 0x01, 0x02, 0x00, 0x02, 0x00, 0xf0, 0x23}
	c.chOut <- exGetInstrumentBarCountCmd
}

// ReqGetMarkets 获取扩展市场数据
func (c *ExternQuoteClient) ReqGetMarkets() {
	c.chOut <- exGetMarketsCmd
}

// GetInstrumentBars 查询K线
func (c *ExternQuoteClient) GetInstrumentBars(req *ReqGetInstrumentBars, requestID int64, h func([]TdxKline, int64)) error {
	var err error
	defer func() {
		if err != nil && h != nil {
			h(nil, requestID)
		}
	}()
	if req.Code == "" {
		err = errors.New("invvalid code")
		return err
	}
	if req.Category > TdxKlineType_DAILY && req.Category != TdxKlineType_EXHQ_1MIN {
		err = errors.New("unsported period")
		log.Println(err)
		return err
	}
	c.requestQueueLock.Lock()
	defer c.requestQueueLock.Unlock()
	c.requestQueue = append(c.requestQueue, requestType{klineType: req.Category, requestID: requestID, handlerFunc: h})

	cmd := []byte{0x01, 0x01, 0x08, 0x6a, 0x01, 0x01, 0x16, 0x00, 0x16, 0x00, 0xff, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.chOut <- cmd
	return err
}

func (c *ExternQuoteClient) onRspGetInstrumentBars(h *header, body []byte) {
	c.requestQueueLock.Lock()
	if len(c.requestQueue) == 0 {
		panic("len(c.requestQueue) == 0")
	}
	req := c.requestQueue[0]
	c.requestQueue = c.requestQueue[1:]
	c.requestQueueLock.Unlock()
	// 前18个字节没解析
	body = body[18:]
	var num uint16
	reader := bytes.NewReader(body)
	binary.Read(reader, binary.LittleEndian, &num)
	pos := 2
	var ks []TdxKline
	for i := 0; i < int(num); i++ {
		y, m, d, h, min := getDateTime(req.klineType, body[pos:])
		var kline TdxKline
		err := unmarshal(body[pos+4:], &kline)
		if err != nil {
			panic(err)
		}
		v, err := time.Parse("20060102 15:04", fmt.Sprintf("%04d%02d%02d %02d:%02d", y, m, d, h, min))
		// log.Println("扩展行情查询K线", y, m, d, h, min, kline, num, i, err, v)
		if err == nil {
			kline.Time = v.Unix() - 3600*8
		}
		pos += 32
		ks = append(ks, kline)
	}
	if req.handlerFunc != nil {
		h := req.handlerFunc.(func([]TdxKline, int64))
		h(ks, req.requestID)
	}

}

// GetInstrumentQuote 查询行情
func (c *ExternQuoteClient) GetInstrumentQuote(req *ReqGetInstrumentQuote) {
	cmd := []byte{0x01, 0x01, 0x08, 0x02, 0x02, 0x01, 0x0c, 0x00, 0x0c, 0x00, 0xfa, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.chOut <- cmd
}

func (c *ExternQuoteClient) onRspGetInstrumentQuote(h *header, body []byte) {
	var rsp RspGetInstrumentQuote
	err := unmarshal(body, &rsp)
	b, err := json.Marshal(rsp)
	log.Println("onRspGetInstrumentQuote", err, string(b))
}

// GetInstrumentInfo 查询合约信息
func (c *ExternQuoteClient) GetInstrumentInfo(start uint32, count uint16) {
	var req ReqGetInstrumentInfo
	req.Start = start
	req.Count = count
	cmd := []byte{0x01, 0x04, 0x48, 0x67, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0xf5, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.chOut <- cmd
}

func (c *ExternQuoteClient) onRspGetInstrumentInfo(h *header, body []byte) {
	var req ReqGetInstrumentInfo
	unmarshal(body, &req)
	pos := 6
	for i := 0; i < int(req.Count); i++ {
		var rsp RspGetInstrumentInfo
		err := unmarshal(body[pos:], &rsp)
		rsp.Code = util.Decode(rsp.Code)
		rsp.Name = util.Decode(rsp.Name)
		rsp.F1 = ""
		rsp.F2 = ""
		rsp.Discription = util.Decode(rsp.Discription)
		log.Println(rsp, err)
		pos += 64
	}
}

// GetInstrumentQuoteList 报价表
func (c *ExternQuoteClient) GetInstrumentQuoteList(market uint8, category uint8, start uint32, count uint16) {

}

// GetMinuteTimeData 查询分时数据
func (c *ExternQuoteClient) GetMinuteTimeData(market uint8, code string) {
	req := ReqGetMinuteTimeData{Market: market, Code: code}
	cmd := []byte{0x01, 0x07, 0x08, 0x00, 0x01, 0x01, 0x0c, 0x00, 0x0c, 0x00, 0x0b, 0x24}
	cmd = append(cmd, marshal(req)...)
	c.chOut <- cmd
}

func (c *ExternQuoteClient) onRspGetMinuteTimeData(h *header, body []byte) {
	var rsp RspGetMinuteTimeDataHeader
	err := unmarshal(body, &rsp)
	if err != nil {
		log.Println(err)
		return
	}
	pos := 12
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetMinuteTimeData
		err = unmarshal(body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		log.Println("分时", hour, minute, item)
		pos += 18
	}

}

// GetHistoryMinuteTimeData 查询历史分时
func (c *ExternQuoteClient) GetHistoryMinuteTimeData(market uint8, code string, date uint32) {
	req := ReqGetHistoryMinuteTimeData{Market: market, Code: code, Date: date}
	cmd := []byte{0x01, 0x01, 0x30, 0x00, 0x01, 0x01, 0x10, 0x00, 0x10, 0x00, 0x0c, 0x24}
	cmd = append(cmd, marshal(req)...)
	c.chOut <- cmd
}

func (c *ExternQuoteClient) onRspGetHistoryMinuteTimeData(h *header, body []byte) {
	var rsp RspGetHistoryMinuteTimeDataHeader
	err := unmarshal(body, &rsp)
	if err != nil {
		log.Println(err)
		return
	}
	pos := 20
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetMinuteTimeData
		err = unmarshal(body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		log.Println("分时", hour, minute, item)
		pos += 18
	}
}

// GetTransactionData 查询分笔数据
func (c *ExternQuoteClient) GetTransactionData(market uint8, code string, start int32, count uint16) {
	req := ReqGetTransactionData{Market: market, Code: code, Start: start, Count: count}
	cmd := []byte{0x01, 0x01, 0x08, 0x00, 0x03, 0x01, 0x12, 0x00, 0x12, 0x00, 0xfc, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.chOut <- cmd
}

func (c *ExternQuoteClient) onRspGetTransactionData(h *header, body []byte) {
	var rsp RspGetTransactionDataHeader
	err := unmarshal(body, &rsp)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("header", rsp, h, len(body))
	pos := 16
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetTransactionData
		err = unmarshal(body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		second := (item.Direction % 10000)
		if second > 59 {
			second = 0
		}
		log.Println("分笔成交", hour, minute, second, item, i, rsp.Count)
		pos += 16
	}
}

// GetHistoryTransactionData 查询历史分笔成交
func (c *ExternQuoteClient) GetHistoryTransactionData(date uint32, market uint8, code string, start int32, count uint16) {
	req := ReqGetHistoryTransactionData{Date: date, Market: market, Code: code, Start: start, Count: count}
	cmd := []byte{0x01, 0x02, 0x30, 0x00, 0x02, 0x01, 0x16, 0x00, 0x16, 0x00, 0x06, 0x24}
	cmd = append(cmd, marshal(req)...)
	c.chOut <- cmd
}

func (c *ExternQuoteClient) onRspGetHistoryTransactionData(h *header, body []byte) {
	var rsp RspGetTransactionDataHeader
	err := unmarshal(body, &rsp)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("header", rsp, h, len(body))
	pos := 16
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetTransactionData
		err = unmarshal(body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		second := (item.Direction % 10000)
		if second > 59 {
			second = 0
		}
		log.Println("历史分笔成交", hour, minute, second, item, i, rsp.Count)
		pos += 16
	}
}
