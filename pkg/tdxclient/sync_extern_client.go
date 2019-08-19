package tdxclient

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
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

// SyncExternClient SyncExternClient
type SyncExternClient struct {
	mu             sync.RWMutex
	ready          bool
	host           string
	conn           net.Conn
	marketList     []ExtRspQryMarket
	referenceCount int
}

// NewSyncExternClient create new sync extern client
func NewSyncExternClient(host string, timeout time.Duration) (*SyncExternClient, error) {
	c := &SyncExternClient{host: host}
	conn, err := net.DialTimeout("tcp", c.host, timeout)
	if err != nil {
		log.Printf("[%s] 连接失败 [%v]", c.host, err)
		return c, err
	}
	_, err = conn.Write(exSetupCmd1)
	if err != nil {
		return c, err
	}
	c.conn = conn
	resp, err := read(conn)
	if err != nil {
		return c, err
	}
	c.ready = true
	log.Printf("[%s]连接成功,funcID:[0x%x], %v", c.host, resp.h.F2, err)
	l, err := c.GetMarketList()
	if err != nil {
		return c, err
		c.ready = false
	}
	c.marketList = l
	return c, nil
}

// GetInstrumentCount 合约数量
func (c *SyncExternClient) GetInstrumentCount() (int, error) {
	if !c.ready {
		return 0, errors.New("unready")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	// 前四个字段等于F2 小端
	c.conn.Write(exGetInstrumentBarCountCmd)
	resp, err := read(c.conn)
	if resp.h.F2 != 0x66480301 {
		return 0, errors.New("resp.h.F2 != 0x66480301")
	}
	// 股票数量
	var rsp RspGetInstrumentCount
	err = unmarshal(resp.body[19:], &rsp)
	if err != nil {
		return 0, err
	}
	// log.Printf("[%s]查询合约数量[%d]", c.host, rsp.Count)
	return int(rsp.Count), nil
}

// GetMarketList 查询市场列表
func (c *SyncExternClient) GetMarketList() ([]ExtRspQryMarket, error) {
	if !c.ready {
		return nil, errors.New("unready")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(exGetMarketsCmd)
	if err != nil {
		return nil, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return nil, err
	}
	if resp.h.F2 != 1766326801 {
		log.Println("resp.h.F2", resp.h.F2)
		return nil, errors.New("resp.h.F2 != 1766326801")
	}
	var num uint16
	var l []ExtRspQryMarket
	reader := bytes.NewReader(resp.body)
	binary.Read(reader, binary.LittleEndian, &num)
	pos := 2
	for i := 0; i < int(num); i++ {
		var item ExtRspQryMarket
		err := unmarshal(resp.body[pos:], &item)
		item.Name = util.Decode(item.Name)
		item.ShortName = util.Decode(item.ShortName)
		if err != nil {
			return nil, err
		}
		if item.Name != "" {
			l = append(l, item)
		}
		// log.Println("扩展行情查询市场", item, num, i)
		pos += 64
	}
	return l, nil
}

// GetInstrumentInfo 合约信息
func (c *SyncExternClient) GetInstrumentInfo(start uint32, count uint16) ([]RspGetInstrumentInfo, error) {
	log.Println("GetInstrumentInfo start", start, count)
	c.mu.Lock()
	defer c.mu.Unlock()
	var req ReqGetInstrumentInfo
	req.Start = start
	req.Count = count
	cmd := []byte{0x01, 0x04, 0x48, 0x67, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0xf5, 0x23}
	cmd = append(cmd, marshal(req)...)
	_, err := c.conn.Write(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return nil, err
	}
	var l []RspGetInstrumentInfo
	var req2 ReqGetInstrumentInfo
	unmarshal(resp.body, &req2)
	pos := 6
	for i := 0; i < int(req2.Count); i++ {
		var rsp RspGetInstrumentInfo
		err := unmarshal(resp.body[pos:], &rsp)
		rsp.Code = util.Decode(rsp.Code)
		rsp.Name = util.Decode(rsp.Name)
		rsp.F1 = ""
		rsp.F2 = ""
		rsp.Discription = util.Decode(rsp.Discription)
		l = append(l, rsp)
		log.Println(rsp, err, req2.Count, i)
		pos += 64
	}
	return l, nil
}

// GetLastTick 查询行情
func (c *SyncExternClient) GetLastTick(ex, symbol string) (*pb.MarketDataSnapshot, error) {
	if ex == "INE" {
		return nil, errors.New("当前TDX的SC合约是用SHFE交易所代码")
	}
	var req ReqGetInstrumentQuote
	req.Code = strings.ToUpper(symbol)
	req.Market = ToTdxMarket(ex)

	cmd := []byte{0x01, 0x01, 0x08, 0x02, 0x02, 0x01, 0x0c, 0x00, 0x0c, 0x00, 0xfa, 0x23}
	cmd = append(cmd, marshal(req)...)
	_, err := c.conn.Write(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return nil, err
	}

	var rsp RspGetInstrumentQuote
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		return nil, err
	}
	var ret pb.MarketDataSnapshot
	switch rsp.Market {
	case 28:
		ret.Exchange = "CZCE"
	case 29:
		ret.Exchange = "DCE"
	case 30:
		ret.Exchange = "SHFE"
	case 47:
		ret.Exchange = "CFFEX"
	}
	ret.Exchange = ex
	ret.Symbol = symbol
	ret.PreClose = float64(rsp.PreClose)
	ret.Open = float64(rsp.Open)
	ret.High = float64(rsp.High)
	ret.Low = float64(rsp.Low)
	ret.Close = float64(rsp.Price)
	ret.Price = ret.Close
	ret.Position = rsp.Position
	ret.Volume = rsp.Volume
	ob := &pb.OrderBook{Ask: float64(rsp.Ask1), Bid: float64(rsp.Bid1), AskVolume: rsp.AskVolume1, BidVolume: rsp.BidVolume1}
	ret.Depths = append(ret.Depths, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask2), Bid: float64(rsp.Bid2), AskVolume: (rsp.AskVolume2), BidVolume: (rsp.BidVolume2)}
	ret.Depths = append(ret.Depths, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask3), Bid: float64(rsp.Bid3), AskVolume: (rsp.AskVolume3), BidVolume: (rsp.BidVolume3)}
	ret.Depths = append(ret.Depths, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask4), Bid: float64(rsp.Bid4), AskVolume: (rsp.AskVolume4), BidVolume: (rsp.BidVolume4)}
	ret.Depths = append(ret.Depths, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask5), Bid: float64(rsp.Bid5), AskVolume: (rsp.AskVolume5), BidVolume: (rsp.BidVolume5)}
	ret.Depths = append(ret.Depths, ob)
	ret.Time = time.Now().Unix()
	if ret.Open == 0 && ret.Price == 0 {
		return &ret, errors.New("InvalidSymbol")
	}
	return &ret, nil
}

// GetInstrumentBars 查询K线
func (c *SyncExternClient) GetInstrumentBars(req *ReqGetInstrumentBars) ([]*TdxKline, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var err error
	var ret []*TdxKline
	if req.Code == "" {
		err = errors.New("invvalid code")
		return ret, err
	}
	if req.Category > TdxKlineType_DAILY && req.Category != TdxKlineType_EXHQ_1MIN {
		err = errors.New("unsported period")
		log.Println(err)
		return ret, err
	}
	cmd := []byte{0x01, 0x01, 0x08, 0x6a, 0x01, 0x01, 0x16, 0x00, 0x16, 0x00, 0xff, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.conn.SetDeadline(time.Now().Add(time.Second * 15))
	_, err = c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	// 前18个字节没解析
	body := resp.body[18:]
	var num uint16
	reader := bytes.NewReader(body)
	binary.Read(reader, binary.LittleEndian, &num)
	pos := 2
	for i := 0; i < int(num); i++ {
		y, m, d, h, min := getDateTime(req.Category, body[pos:])
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
		ret = append(ret, &kline)
	}
	return ret, nil
}

// GetMinuteTimeData 查询分时数据
func (c *SyncExternClient) GetMinuteTimeData(market uint8, code string) ([]*RspGetMinuteTimeData, error) {
	var ret []*RspGetMinuteTimeData
	req := ReqGetMinuteTimeData{Market: market, Code: code}
	cmd := []byte{0x01, 0x07, 0x08, 0x00, 0x01, 0x01, 0x0c, 0x00, 0x0c, 0x00, 0x0b, 0x24}
	cmd = append(cmd, marshal(req)...)
	_, err := c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}
	var rsp RspGetMinuteTimeDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	pos := 12
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetMinuteTimeData
		err = unmarshal(resp.body[pos:], &item)
		// hour := (item.Time / 60)
		// minute := (item.Time % 60)
		// log.Println("分时", hour, minute, item)
		pos += 18
		ret = append(ret, &item)
	}
	return ret, nil
}

// GetHistoryMinuteTimeData 查询历史分时
func (c *SyncExternClient) GetHistoryMinuteTimeData(market uint8, code string, date uint32) ([]*RspGetMinuteTimeData, error) {
	var ret []*RspGetMinuteTimeData
	req := ReqGetHistoryMinuteTimeData{Market: market, Code: code, Date: date}
	cmd := []byte{0x01, 0x01, 0x30, 0x00, 0x01, 0x01, 0x10, 0x00, 0x10, 0x00, 0x0c, 0x24}
	cmd = append(cmd, marshal(req)...)
	_, err := c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}
	var rsp RspGetHistoryMinuteTimeDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	pos := 20
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetMinuteTimeData
		err = unmarshal(resp.body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		log.Println("分时", hour, minute, item)
		pos += 18
	}
	return ret, nil
}

// GetTransactionData 查询分笔数据
func (c *SyncExternClient) GetTransactionData(market uint8, code string, start int32, count uint16) ([]*RspGetTransactionData, error) {
	var ret []*RspGetTransactionData
	req := ReqGetTransactionData{Market: market, Code: code, Start: start, Count: count}
	cmd := []byte{0x01, 0x01, 0x08, 0x00, 0x03, 0x01, 0x12, 0x00, 0x12, 0x00, 0xfc, 0x23}
	cmd = append(cmd, marshal(req)...)
	_, err := c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}
	var rsp RspGetTransactionDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	log.Println("resp: ", resp)
	pos := 16
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetTransactionData
		err = unmarshal(resp.body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		second := (item.Direction % 10000)
		if second > 59 {
			second = 0
		}
		log.Println("分笔成交", hour, minute, second, item, i, rsp.Count)
		pos += 16
	}
	return ret, nil
}

// GetHistoryTransactionData 查询历史分笔成交
func (c *SyncExternClient) GetHistoryTransactionData(date uint32, market uint8, code string, start int32, count uint16) ([]*RspGetTransactionData, error) {
	var ret []*RspGetTransactionData
	req := ReqGetHistoryTransactionData{Date: date, Market: market, Code: code, Start: start, Count: count}
	cmd := []byte{0x01, 0x02, 0x30, 0x00, 0x02, 0x01, 0x16, 0x00, 0x16, 0x00, 0x06, 0x24}
	cmd = append(cmd, marshal(req)...)
	_, err := c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}
	var rsp RspGetTransactionDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	pos := 16
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetTransactionData
		err = unmarshal(resp.body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		second := (item.Direction % 10000)
		if second > 59 {
			second = 0
		}
		log.Println("历史分笔成交", hour, minute, second, item, i, rsp.Count)
		pos += 16
	}
	return ret, nil
}
