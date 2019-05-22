package tdxclient

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"net"
	"time"

	"github.com/mineralres/goshare/pkg/util"
)

// QuoteClient 行情接口
type QuoteClient struct {
	chOut      chan []byte
	retryTimes int
}

// ConnectAndJoin (*ftdc.Trader) 连接
func (c *QuoteClient) ConnectAndJoin(ctx context.Context, address string, h func(), autoReconnect bool) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println("连接失败", err)
		return
	}
	defer func() {
		conn.Close()
		log.Println("通达信连接断开")
		if autoReconnect {
			time.Sleep(time.Second)
			c.retryTimes++
			go c.ConnectAndJoin(ctx, address, h, autoReconnect)
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
				c.ReqQryStockCount(0)
			case <-ctx.Done():
				return
			case d := <-c.chOut:
				if d == nil || len(d) == 0 {
					continue
				}
				n, err := conn.Write(d)
				if err != nil {
					conn.Close()
				}
				// log.Println("发送成功", n, err, d)
				if n < len(d) {
					panic("n < len(d)")
				}
			}
		}
	}()
	c.setup()
	h()
	// 启动读
	c.read(conn)
}

func (c *QuoteClient) setup() {
	pkgSetupCmd1 := []byte{0x0c, 0x02, 0x18, 0x93, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x01}
	pkgSetupCmd2 := []byte{0x0c, 0x02, 0x18, 0x94, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x02}
	pkgSetupCmd3 := []byte{0x0c, 0x03, 0x18, 0x99, 0x00, 0x01, 0x20, 0x00, 0x20, 0x00, 0xdb, 0x0f, 0xd5,
		0xd0, 0xc9, 0xcc, 0xd6, 0xa4, 0xa8, 0xaf, 0x00, 0x00, 0x00, 0x8f, 0xc2, 0x25, 0x40, 0x13, 0x00, 0x00,
		0xd5, 0x00, 0xc9, 0xcc, 0xbd, 0xf0, 0xd7, 0xea, 0x00, 0x00, 0x00, 0x02}
	c.send(pkgSetupCmd1)
	c.send(pkgSetupCmd2)
	c.send(pkgSetupCmd3)

}

func (c *QuoteClient) read(conn net.Conn) {
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

func (c *QuoteClient) send(pkg []byte) {
	c.chOut <- pkg
}

func (c *QuoteClient) handleMessage(h *header, body []byte) {
	b, _ := json.Marshal(h)
	log.Printf("收到消息 %s len(body)=%d", string(b), len(body))
	switch h.F2 {
	case 0x6c180c0c:
		// 股票数量
		var rsp RspQryStockCount
		err := unmarshal(body, &rsp)
		if err != nil {
			panic(err)
		}
		log.Printf("查询股票数量[%d]", rsp.Count)
	case 0x6418011c:
		var num uint16
		reader := bytes.NewReader(body)
		binary.Read(reader, binary.LittleEndian, &num)
		pos := 2
		for i := 0; i < int(num); i++ {
			var item RspQrySecurity
			err := unmarshal(body[pos:], &item)
			item.Name = util.Decode(item.Name)
			preClose := GetVolume(item.PreCloseRaw)
			if err != nil {
				panic(err)
			}
			log.Println("查询证券", item, num, i, preClose)
			pos += 29
		}

		log.Println("查询证券列表返回")
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

	}
}

// ReqQryStockCount 查询股票数量
func (c *QuoteClient) ReqQryStockCount(market uint16) {
	// 前四个字段等于F2 小端
	cmd := []byte{0x0c, 0x0c, 0x18, 0x6c, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0x4e, 0x04, 0x00, 0x00, 0x75, 0xc7, 0x33, 0x01}
	c.chOut <- cmd
}

// ReqGetSecurityList 查询证券列表
func (c *QuoteClient) ReqGetSecurityList(market, start uint16) {
	buf := &bytes.Buffer{}
	cmd := []byte{0x1c, 0x01, 0x18, 0x64, 0x01, 0x01, 0x06, 0x00, 0x06, 0x00, 0x50, 0x04}
	buf.Write(cmd)
	binary.Write(buf, binary.LittleEndian, &market)
	binary.Write(buf, binary.LittleEndian, &start)
	c.chOut <- buf.Bytes()
}

// ExtReqGetMarkets 获取扩展市场数据
func (c *QuoteClient) ExtReqGetMarkets() {
	cmd := []byte{0x01, 0x02, 0x48, 0x69, 0x00, 0x01, 0x02, 0x00, 0x02, 0x00, 0xf4, 0x23}
	c.chOut <- cmd
}
