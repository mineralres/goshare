package ctp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

type packet struct {
	MsgType    int32
	RequestID  int32
	IsLast     int32
	BodyLength int32
	BodyList   [][]byte
}

func (pkt *packet) parseBody(body []byte) error {
	if len(body) < 4 {
		return nil
	}
	var l int32
	binary.Read(bytes.NewReader(body), binary.LittleEndian, &l)
	if l+4 > int32(len(body)) {
		return nil
	}
	pkt.BodyList = append(pkt.BodyList, body[4:4+l])
	if len(pkt.BodyList) > 10 {
		return nil
	}
	return pkt.parseBody(body[4+l:])
}

// Adapter ctp socket adapter
type Adapter struct {
	chOut chan *packet
	h     func(*packet)
}

// NewAdapter create new adapter
func NewAdapter(host string, h func(*packet)) *Adapter {
	adapter := &Adapter{h: h}
	adapter.chOut = make(chan *packet, 100)
	adapter.bind(host)
	return adapter
}

// Bind bind to ctp hub gateway
func (adapter *Adapter) bind(host string) error {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err
	}
	// 读函数
	reader := bufio.NewReader(conn)
	go func() {
		defer func() {
			conn.Close()
		}()
		for {
			// read
			pkt := new(packet)
			if err := binary.Read(reader, binary.LittleEndian, &pkt.MsgType); err != nil {
				log.Println(err)
				return
			}
			if err := binary.Read(reader, binary.LittleEndian, &pkt.RequestID); err != nil {
				log.Println(err)
				return
			}
			if err := binary.Read(reader, binary.LittleEndian, &pkt.IsLast); err != nil {
				log.Println(err)
				return
			}
			if err := binary.Read(reader, binary.LittleEndian, &pkt.BodyLength); err != nil {
				log.Println(err)
				return
			}
			body := make([]byte, pkt.BodyLength)
			if _, err := io.ReadFull(reader, body); err != nil {
				log.Println(err)
				return
			}
			// 读取成功处理packet
			err := pkt.parseBody(body)
			if err == nil {
				adapter.h(pkt)
			} else {
				log.Println("parse body error:", err)
				return
			}
		}
	}()

	// write函数
	go func() {
		for p := range adapter.chOut {
			// log.Println("send p ", p.MsgType)
			if err = binary.Write(conn, binary.LittleEndian, p.MsgType); err != nil {
				log.Println(err)
				return
			}
			if err = binary.Write(conn, binary.LittleEndian, p.RequestID); err != nil {
				log.Println(err)
				return
			}
			if err = binary.Write(conn, binary.LittleEndian, p.IsLast); err != nil {
				log.Println(err)
				return
			}
			p.BodyLength = 0
			for i := range p.BodyList {
				p.BodyLength += int32(len(p.BodyList[i]) + 4)
			}
			if err = binary.Write(conn, binary.LittleEndian, p.BodyLength); err != nil {
				log.Println(err)
				return
			}
			for i := range p.BodyList {
				if err = binary.Write(conn, binary.LittleEndian, int32(len(p.BodyList[i]))); err != nil {
					log.Println(err)
					return
				}
				if _, err := conn.Write(p.BodyList[i]); err != nil {
					log.Println(err, p.BodyList[i])
					return
				}
			}
		}
	}()
	return nil
}

// Send send request
func (adapter *Adapter) Send(msgType int32, d1 []byte, requestID int32) {
	pkt := &packet{MsgType: msgType, RequestID: requestID}
	if d1 != nil {
		pkt.BodyList = append(pkt.BodyList, d1)
	}
	adapter.chOut <- pkt
}
