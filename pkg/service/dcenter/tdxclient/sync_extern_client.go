package tdxclient

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/util"
)

// SyncExternClient SyncExternClient
type SyncExternClient struct {
	sync.RWMutex
	ready      bool
	address    string
	conn       net.Conn
	marketList []ExtRspQryMarket
}

func makeSyncExternClient(address string) *SyncExternClient {
	return &SyncExternClient{address: address}
}

// connect
func (c *SyncExternClient) init() {
	connect := func() {
		c.Lock()
		defer c.Unlock()
		if c.ready {
			return
		}
		c.ready = false
		conn, err := net.DialTimeout("tcp", c.address, time.Second*3)
		if err != nil {
			log.Printf("[%s] 连接失败 [%v]", c.address, err)
			return
		}
		n, err := conn.Write(exSetupCmd1)
		if err != nil || n != len(exSetupCmd1) {
			c.ready = false
			return
		}
		c.conn = conn
		resp, err := c.read()
		if err != nil {
			c.ready = false
		}
		c.ready = true
		log.Printf("[%s]连接成功,funcID:[0x%x], %v", c.address, resp.h.F2, err)
	}
	connect()
	l, err := c.GetMarketList()
	if err != nil {
		c.ready = false
		log.Println(err)
	}
	c.marketList = l
	c.GetInstrumentQuote(&ReqGetInstrumentQuote{30, "SC1906"})
}

func (c *SyncExternClient) read() (*tdxResponse, error) {
	conn := c.conn
	conn.SetReadDeadline(time.Now().Add(time.Second * 15))
	var h header
	headerBuf := make([]byte, 16)
	readed := 0
	for {
		n, err := conn.Read(headerBuf[readed:])
		if err != nil {
			c.ready = false
			log.Println(err)
			return nil, err
		}
		readed += n
		if readed >= len(headerBuf) {
			err = unmarshal(headerBuf, &h)
			if err != nil {
				log.Println("invalid ex tdx header")
				c.ready = false
				return nil, err
			}
			break
		}
	}
	body := make([]byte, h.ZipSize)
	readed = 0
	for {
		n, err := conn.Read(body[readed:])
		if err != nil {
			c.ready = false
			return nil, err
		}
		readed += n
		if readed >= int(h.ZipSize) {
			break
		}
	}
	if h.ZipSize != h.UnzipSize {
		// log.Println("需要解压")
		r, err := zlib.NewReader(bytes.NewReader(body))
		if err != nil {
			log.Println(err, len(body), h, readed)
			c.ready = false
			return nil, err
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
	c.ready = true
	return &tdxResponse{h: h, body: body}, nil
}

func (c *SyncExternClient) write(b []byte) (int, error) {
	n, err := c.conn.Write(b)
	if err != nil {
		c.ready = false
		return n, err
	}
	if n != len(b) {
		panic("n != len(b)")
	}
	return n, err
}

// GetInstrumentCount 合约数量
func (c *SyncExternClient) GetInstrumentCount() (int, error) {
	if !c.ready {
		return 0, errors.New("unready")
	}
	c.Lock()
	defer c.Unlock()
	// 前四个字段等于F2 小端
	c.write(exGetInstrumentBarCountCmd)
	resp, err := c.read()
	if resp.h.F2 != 0x66480301 {
		return 0, errors.New("resp.h.F2 != 0x66480301")
	}
	// 股票数量
	var rsp RspGetInstrumentCount
	err = unmarshal(resp.body[19:], &rsp)
	if err != nil {
		return 0, err
	}
	// log.Printf("[%s]查询合约数量[%d]", c.address, rsp.Count)
	return int(rsp.Count), nil
}

// GetMarketList 查询市场列表
func (c *SyncExternClient) GetMarketList() ([]ExtRspQryMarket, error) {
	if !c.ready {
		return nil, errors.New("unready")
	}
	c.Lock()
	defer c.Unlock()
	n, err := c.write(exGetMarketsCmd)
	if err != nil || n != len(exGetMarketsCmd) {
		return nil, err
	}
	resp, err := c.read()
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
	var req ReqGetInstrumentInfo
	req.Start = start
	req.Count = count
	cmd := []byte{0x01, 0x04, 0x48, 0x67, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0xf5, 0x23}
	cmd = append(cmd, marshal(req)...)
	n, err := c.write(cmd)
	if err != nil || n != len(cmd) {
		return nil, err
	}
	resp, err := c.read()
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

// GetInstrumentQuote 查询行情
func (c *SyncExternClient) GetInstrumentQuote(req *ReqGetInstrumentQuote) (*pb.MarketDataSnapshot, error) {
	cmd := []byte{0x01, 0x01, 0x08, 0x02, 0x02, 0x01, 0x0c, 0x00, 0x0c, 0x00, 0xfa, 0x23}
	cmd = append(cmd, marshal(req)...)
	_, err := c.write(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := c.read()

	var rsp RspGetInstrumentQuote
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		return nil, err
	}
	var ret pb.MarketDataSnapshot
	ret.Symbol = &pb.Symbol{}
	switch rsp.Market {
	case 28:
		ret.Symbol.Exchange = pb.ExchangeType_CZCE
	case 29:
		ret.Symbol.Exchange = pb.ExchangeType_DCE
	case 30:
		ret.Symbol.Exchange = pb.ExchangeType_SHFE
	case 47:
		ret.Symbol.Exchange = pb.ExchangeType_CFFEX
	}
	ret.Symbol.Code = rsp.Code
	ret.PreClose = float64(rsp.PreClose)
	ret.Open = float64(rsp.Open)
	ret.High = float64(rsp.High)
	ret.Low = float64(rsp.Low)
	ret.Close = float64(rsp.Price)
	ret.Price = ret.Close
	ret.Position = rsp.Position
	ret.Volume = rsp.Volume
	ob := &pb.OrderBook{Ask: float64(rsp.Ask1), Bid: float64(rsp.Bid1), AskVolume: rsp.AskVolume1, BidVolume: rsp.BidVolume1}
	ret.OrderBookList = append(ret.OrderBookList, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask2), Bid: float64(rsp.Bid2), AskVolume: (rsp.AskVolume2), BidVolume: (rsp.BidVolume2)}
	ret.OrderBookList = append(ret.OrderBookList, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask3), Bid: float64(rsp.Bid3), AskVolume: (rsp.AskVolume3), BidVolume: (rsp.BidVolume3)}
	ret.OrderBookList = append(ret.OrderBookList, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask4), Bid: float64(rsp.Bid4), AskVolume: (rsp.AskVolume4), BidVolume: (rsp.BidVolume4)}
	ret.OrderBookList = append(ret.OrderBookList, ob)
	ob = &pb.OrderBook{Ask: float64(rsp.Ask5), Bid: float64(rsp.Bid5), AskVolume: (rsp.AskVolume5), BidVolume: (rsp.BidVolume5)}
	ret.OrderBookList = append(ret.OrderBookList, ob)
	ret.Time = time.Now().Unix()
	if ret.Open == 0 && ret.Price == 0 {
		return &ret, errors.New("InvalidSymbol")
	}
	// b, err := json.Marshal(rsp)
	// log.Println("onRspGetInstrumentQuote", err, string(b), len(resp.body))
	return &ret, nil
}
