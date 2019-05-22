package tdxclient

import (
	"bytes"
	"compress/zlib"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

// SyncClient SyncClient
type SyncClient struct {
	sync.RWMutex
	ready   bool
	address string
	conn    net.Conn
}

func makeSyncClient(address string) *SyncClient {
	return &SyncClient{address: address}
}

func (c *SyncClient) init() {
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
		pkgSetupCmd1 := []byte{0x0c, 0x02, 0x18, 0x93, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x01}
		pkgSetupCmd2 := []byte{0x0c, 0x02, 0x18, 0x94, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x02}
		pkgSetupCmd3 := []byte{0x0c, 0x03, 0x18, 0x99, 0x00, 0x01, 0x20, 0x00, 0x20, 0x00, 0xdb, 0x0f, 0xd5,
			0xd0, 0xc9, 0xcc, 0xd6, 0xa4, 0xa8, 0xaf, 0x00, 0x00, 0x00, 0x8f, 0xc2, 0x25, 0x40, 0x13, 0x00, 0x00,
			0xd5, 0x00, 0xc9, 0xcc, 0xbd, 0xf0, 0xd7, 0xea, 0x00, 0x00, 0x00, 0x02}
		n, err := conn.Write(pkgSetupCmd1)
		n, err = conn.Write(pkgSetupCmd2)
		n, err = conn.Write(pkgSetupCmd3)
		if err != nil || n == 0 {
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
}

func (c *SyncClient) read() (*tdxResponse, error) {
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
