package tdxclient

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"sync"
	"time"

	"github.com/mineralres/goshare/pkg/util"
)

// SyncQuoteClient SyncQuoteClient
type SyncQuoteClient struct {
	sync.RWMutex
	ready          bool
	host           string
	conn           net.Conn
	referenceCount int
}

// NewSyncQuoteClient create sync quote client
func NewSyncQuoteClient(host string, timeout time.Duration) (*SyncQuoteClient, error) {
	c := &SyncQuoteClient{host: host}
	conn, err := net.DialTimeout("tcp", host, timeout)
	if err != nil {
		log.Printf("[%s] 连接失败 [%v]", host, err)
		return c, err
	}
	pkgSetupCmd1 := []byte{0x0c, 0x02, 0x18, 0x93, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x01}
	pkgSetupCmd2 := []byte{0x0c, 0x02, 0x18, 0x94, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x02}
	pkgSetupCmd3 := []byte{0x0c, 0x03, 0x18, 0x99, 0x00, 0x01, 0x20, 0x00, 0x20, 0x00, 0xdb, 0x0f, 0xd5,
		0xd0, 0xc9, 0xcc, 0xd6, 0xa4, 0xa8, 0xaf, 0x00, 0x00, 0x00, 0x8f, 0xc2, 0x25, 0x40, 0x13, 0x00, 0x00,
		0xd5, 0x00, 0xc9, 0xcc, 0xbd, 0xf0, 0xd7, 0xea, 0x00, 0x00, 0x00, 0x02}
	_, err = conn.Write(pkgSetupCmd1)
	resp, err := read(conn)
	if err != nil {
		return c, err
	}

	_, err = conn.Write(pkgSetupCmd2)
	resp, err = read(conn)
	if err != nil {
		return c, err
	}

	_, err = conn.Write(pkgSetupCmd3)
	resp, err = read(conn)
	if err != nil {
		return c, err
	}
	c.conn = conn
	c.ready = true
	log.Printf("[%s]连接成功,funcID:[0x%x], %v", host, resp.h.F2, err)
	return c, nil
}

// ReqQryStockCount 查询股票数量
func (c *SyncQuoteClient) ReqQryStockCount() (uint16, error) {
	cmd := []byte{0x0c, 0x0c, 0x18, 0x6c, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0x4e, 0x04, 0x00, 0x00, 0x75, 0xc7, 0x33, 0x01}
	_, err := c.conn.Write(cmd)
	if err != nil {
		return 0, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return 0, err
	}
	var rsp RspQryStockCount
	err = unmarshal(resp.body, &rsp)
	return rsp.Count, err
}

// ReqGetSecurityList 查询证券列表
func (c *SyncQuoteClient) ReqGetSecurityList(market, start uint16) ([]*RspQrySecurity, error) {
	var ret []*RspQrySecurity
	buf := &bytes.Buffer{}
	// pkg = bytearray.fromhex(u'0c 01 18 64 01 01 06 00 06 00 50 04')
	cmd := []byte{0x1c, 0x01, 0x18, 0x64, 0x01, 0x01, 0x06, 0x00, 0x06, 0x00, 0x50, 0x04}
	buf.Write(cmd)
	binary.Write(buf, binary.LittleEndian, &market)
	binary.Write(buf, binary.LittleEndian, &start)
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}

	var num uint16
	reader := bytes.NewReader(resp.body)
	binary.Read(reader, binary.LittleEndian, &num)
	pos := 2
	for i := 0; i < int(num); i++ {
		var item RspQrySecurity
		err := unmarshal(resp.body[pos:], &item)
		if err == nil {
			item.Name = util.Decode(item.Name)
		}
		// preClose := GetVolume(item.PreCloseRaw)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println("查询证券", item, num, i, preClose)
		pos += 29
		ret = append(ret, &item)
	}
	return ret, nil
}

// ReqGetSecurityQuotes quotes
func (c *SyncQuoteClient) ReqGetSecurityQuotes(stockList []*ReqGetInstrumentQuote) ([]*SecurityQuote, error) {
	var ret []*SecurityQuote
	c.conn.SetDeadline(time.Now().Add(time.Second * 10))
	stockLen := len(stockList)
	cmd := []byte{}
	var req struct {
		F1          uint16
		F2          uint32
		PkgDataLen  uint16
		PkgDataLen2 uint16
		F3          uint32
		F4          uint32
		F5          uint16
		StockLen    uint16
	}
	// pkg_header = struct.pack("<HIHHIIHH", *values)
	req.F1 = 0x10c
	req.F2 = 0x02006320
	req.PkgDataLen = uint16(stockLen*7 + 12)
	req.PkgDataLen2 = req.PkgDataLen
	req.F3 = 0x5053e
	req.StockLen = uint16(stockLen)
	cmd = append(cmd, marshal(&req)...)
	for _, req := range stockList {
		cmd = append(cmd, req.Market)
		if len(req.Code) > 6 {
			cmd = append(cmd, req.Code[:6]...)
		} else if len(req.Code) <= 6 {
			cmd = append(cmd, req.Code...)
			for i := 0; i < 6-len(req.Code); i++ {
				cmd = append(cmd, 0)
			}
		}
	}
	// log.Println(cmd)
	_, err := c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}
	var num uint16
	pos := 2
	d := resp.body
	reader := bytes.NewReader(d[pos:])
	binary.Read(reader, binary.LittleEndian, &num)
	pos += 2

	for i := 0; i < int(num); i++ {
		var item SecurityQuote
		item.Market = d[pos]
		item.Code = string(d[pos+1 : pos+7])
		pos += 9
		item.Price, pos = getPrice(d, pos)
		item.LastCloseDiff, pos = getPrice(d, pos)
		item.OpenDiff, pos = getPrice(d, pos)
		item.HighDiff, pos = getPrice(d, pos)
		item.LowDiff, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		item.Vol, pos = getPrice(d, pos)
		item.CurVol, pos = getPrice(d, pos)
		var rawAmount uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawAmount)
		pos += 4
		item.RawAmount = rawAmount
		item.Amount = GetVolume(rawAmount)
		item.SVol, pos = getPrice(d, pos)
		item.BVol, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)

		item.Bid1, pos = getPrice(d, pos)
		item.Ask1, pos = getPrice(d, pos)
		item.BidVol1, pos = getPrice(d, pos)
		item.AskVol1, pos = getPrice(d, pos)

		item.Bid2, pos = getPrice(d, pos)
		item.Ask2, pos = getPrice(d, pos)
		item.BidVol2, pos = getPrice(d, pos)
		item.AskVol2, pos = getPrice(d, pos)

		item.Bid3, pos = getPrice(d, pos)
		item.Ask3, pos = getPrice(d, pos)
		item.BidVol3, pos = getPrice(d, pos)
		item.AskVol3, pos = getPrice(d, pos)

		item.Bid4, pos = getPrice(d, pos)
		item.Ask4, pos = getPrice(d, pos)
		item.BidVol4, pos = getPrice(d, pos)
		item.AskVol4, pos = getPrice(d, pos)

		item.Bid5, pos = getPrice(d, pos)
		item.Ask5, pos = getPrice(d, pos)
		item.BidVol5, pos = getPrice(d, pos)
		item.AskVol5, pos = getPrice(d, pos)

		// reversed_bytes4 = struct.unpack("<H", d[pos:pos+2])
		pos += 2
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		// (reversed_bytes9, active2) = struct.unpack("<hH", d[pos: pos + 4])
		pos += 4

		ret = append(ret, &item)
	}
	return ret, nil
}

// ReqGetSecurityBars get kline
func (c *SyncQuoteClient) ReqGetSecurityBars(category, market uint16, code string, start, count uint16) ([]*SecurityBar, error) {
	var ret []*SecurityBar
	var req struct {
		F1       uint16
		F2       uint32
		F3       uint16
		F4       uint16
		F5       uint16
		Market   uint16
		Code     string `xlen:"6"`
		Category uint16
		F6       uint16
		Start    uint16
		Count    uint16
		F7       uint32
		F8       uint32
		F9       uint16
	}
	req.F1 = 0x10c
	req.F2 = 0x01016408
	req.F3 = 0x1c
	req.F4 = 0x1c
	req.F5 = 0x052d
	req.F6 = 1
	req.Category = category
	req.Market = market
	req.Code = code
	req.Start = start
	req.Count = count
	_, err := c.conn.Write(marshal(&req))
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}
	d := resp.body
	var num uint16
	binary.Read(bytes.NewReader(d[0:2]), binary.LittleEndian, &num)
	pos := 2
	var preDiffBase int
	for i := 0; i < int(num); i++ {
		var item SecurityBar
		item.Year, item.Mon, item.Day, item.Hour, item.Minute = getDateTime(TdxKlineType(category), d[pos:])
		pos += 4
		item.PriceOpenDiff, pos = getPrice(d, pos)
		item.PriceCloseDiff, pos = getPrice(d, pos)

		item.PriceHighDiff, pos = getPrice(d, pos)
		item.PriceLowDiff, pos = getPrice(d, pos)

		var rawVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawVol)
		vol := GetVolume(rawVol)
		pos += 4
		item.RawVol = rawVol
		item.Vol = int(vol)

		var rawDBVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawDBVol)
		dbvol := GetVolume(rawDBVol)
		pos += 4
		item.RawDBVol = rawDBVol
		item.DBVol = int(dbvol)

		item.Open = calPrice1000(item.PriceOpenDiff, preDiffBase)

		item.PriceOpenDiff = item.PriceOpenDiff + preDiffBase

		item.Close = calPrice1000(item.PriceOpenDiff, item.PriceCloseDiff)
		item.High = calPrice1000(item.PriceOpenDiff, item.PriceHighDiff)
		item.Low = calPrice1000(item.PriceOpenDiff, item.PriceLowDiff)

		preDiffBase = item.PriceOpenDiff + item.PriceCloseDiff

		ret = append(ret, &item)
	}
	return ret, nil
}

// ReqGetIndexBars req get index bar
func (c *SyncQuoteClient) ReqGetIndexBars(category, market uint16, code string, start, count uint16) ([]*SecurityBar, error) {
	var ret []*SecurityBar
	var req struct {
		F1       uint16
		F2       uint32
		F3       uint16
		F4       uint16
		F5       uint16
		Market   uint16
		Code     string `xlen:"6"`
		Category uint16
		F6       uint16
		Start    uint16
		Count    uint16
		F7       uint32
		F8       uint32
		F9       uint16
	}
	req.F1 = 0x10c
	req.F2 = 0x01016408
	req.F3 = 0x1c
	req.F4 = 0x1c
	req.F5 = 0x052d
	req.F6 = 1
	req.Category = category
	req.Market = market
	req.Code = code
	req.Start = start
	req.Count = count
	_, err := c.conn.Write(marshal(&req))
	if err != nil {
		return ret, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return ret, err
	}
	d := resp.body
	var num uint16
	binary.Read(bytes.NewReader(d[0:2]), binary.LittleEndian, &num)
	pos := 2
	var preDiffBase int
	for i := 0; i < int(num); i++ {
		var item SecurityBar
		item.Year, item.Mon, item.Day, item.Hour, item.Minute = getDateTime(TdxKlineType(category), d[pos:])
		pos += 4
		item.PriceOpenDiff, pos = getPrice(d, pos)
		item.PriceCloseDiff, pos = getPrice(d, pos)

		item.PriceHighDiff, pos = getPrice(d, pos)
		item.PriceLowDiff, pos = getPrice(d, pos)

		var rawVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawVol)
		vol := GetVolume(rawVol)
		pos += 4
		item.RawVol = rawVol
		item.Vol = int(vol)

		var rawDBVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawDBVol)
		dbvol := GetVolume(rawDBVol)
		pos += 4
		item.RawDBVol = rawDBVol
		item.DBVol = int(dbvol)

		binary.Read(bytes.NewReader(d[pos:pos+2]), binary.LittleEndian, &item.UpCount)
		binary.Read(bytes.NewReader(d[pos+2:pos+4]), binary.LittleEndian, &item.DownCount)
		pos += 4

		item.Open = calPrice1000(item.PriceOpenDiff, preDiffBase)

		item.PriceOpenDiff = item.PriceOpenDiff + preDiffBase

		item.Close = calPrice1000(item.PriceOpenDiff, item.PriceCloseDiff)
		item.High = calPrice1000(item.PriceOpenDiff, item.PriceHighDiff)
		item.Low = calPrice1000(item.PriceOpenDiff, item.PriceLowDiff)

		preDiffBase = item.PriceOpenDiff + item.PriceCloseDiff

		ret = append(ret, &item)

	}
	return ret, nil
}

// ReqGetMinuteTimeData get minute time data
func (c *SyncQuoteClient) ReqGetMinuteTimeData(market uint16, code string) ([]*RspGetMinuteTimeData, error) {
	var ret []*RspGetMinuteTimeData
	if len(code) != 6 {
		return nil, errors.New("code len should be 6")
	}
	cmd := []byte{0x0c, 0x1b, 0x08, 0x00, 0x01, 0x01, 0x0e, 0x00, 0x0e, 0x00, 0x1d, 0x05}
	buf := &bytes.Buffer{}
	buf.Write(cmd)
	binary.Write(buf, binary.LittleEndian, &market)
	buf.Write([]byte(code))
	buf.Write([]byte{0, 0, 0, 0})
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	resp, err := read(c.conn)
	if err != nil {
		return nil, err
	}
	d := resp.body
	var num uint16
	binary.Read(bytes.NewReader(d[0:2]), binary.LittleEndian, &num)
	log.Println(num, len(resp.body))
	pos := 4
	var lastPrice int
	for i := 0; i < int(num); i++ {
		var item RspGetMinuteTimeData
		rawPrice, pos := getPrice(d, pos)
		_, pos = getPrice(d, pos)
		vol, pos := getPrice(d, pos)
		lastPrice += rawPrice
		item.Time = uint16(lastPrice)
		item.Volume = uint32(vol)
		item.Price = float32(rawPrice)
		ret = append(ret, &item)
	}
	return ret, nil
}
