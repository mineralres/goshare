package ctp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
	"time"

	proto "github.com/golang/protobuf/proto"
)

// Packet pkt
type Packet struct {
	MsgType    int32
	RequestID  int32
	IsLast     int32
	BodyLength int32
	BodyList   [][]byte
}

// Get1 get 1 param
func (pkt *Packet) Get1(p1 proto.Message) error {
	if len(pkt.BodyList) < 1 {
		log.Println("len(pkt.BodyList) < 1")
		return errors.New("len(pkt.BodyList) < 1")
	}
	if err := proto.Unmarshal(pkt.BodyList[0], p1); err != nil {
		return err
	}
	return nil
}

// Get2 get 2 param
func (pkt *Packet) Get2(p1 proto.Message, p2 proto.Message) error {
	if len(pkt.BodyList) < 2 {
		log.Println("len(pkt.BodyList) < 2")
		return errors.New("len(pkt.BodyList) < 2")
	}
	if err := proto.Unmarshal(pkt.BodyList[0], p1); err != nil {
		return err
	}
	if err := proto.Unmarshal(pkt.BodyList[1], p2); err != nil {
		return err
	}
	return nil
}

func (pkt *Packet) parseBody(body []byte) error {
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

func (pkt *Packet) writeTo(conn net.Conn) error {
	var err error
	if err = binary.Write(conn, binary.LittleEndian, pkt.MsgType); err != nil {
		return err
	}
	if err = binary.Write(conn, binary.LittleEndian, pkt.RequestID); err != nil {
		return err
	}
	if err = binary.Write(conn, binary.LittleEndian, pkt.IsLast); err != nil {
		return err
	}
	pkt.BodyLength = 0
	for i := range pkt.BodyList {
		pkt.BodyLength += int32(len(pkt.BodyList[i]) + 4)
	}
	if err = binary.Write(conn, binary.LittleEndian, pkt.BodyLength); err != nil {
		return err
	}
	for i := range pkt.BodyList {
		if err = binary.Write(conn, binary.LittleEndian, int32(len(pkt.BodyList[i]))); err != nil {
			return err
		}
		if _, err := conn.Write(pkt.BodyList[i]); err != nil {
			return err
		}
	}
	return nil
}

func (pkt *Packet) readFrom(reader *bufio.Reader) error {
	if err := binary.Read(reader, binary.LittleEndian, &pkt.MsgType); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.LittleEndian, &pkt.RequestID); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.LittleEndian, &pkt.IsLast); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.LittleEndian, &pkt.BodyLength); err != nil {
		return err
	}
	body := make([]byte, pkt.BodyLength)
	if _, err := io.ReadFull(reader, body); err != nil {
		return err
	}
	pkt.parseBody(body)
	return nil
}

// Adapter ctp socket adapter
type Adapter struct {
	chOut chan *Packet
}

// NewAdapter create new adapter
func NewAdapter(host string, h func(*Packet)) (*Adapter, error) {
	ret := &Adapter{}
	ret.chOut = make(chan *Packet, 100)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(conn)
	// read messages
	go func() {
		for {
			conn.SetReadDeadline(time.Now().Add(time.Second * 60))
			pkt := new(Packet)
			err = pkt.readFrom(reader)
			if err != nil {
				return
			}
			if h != nil {
				h(pkt)
			}
		}
	}()
	// write函数
	go func() {
		for {
			select {
			// heartbeat
			case <-time.After(20 * time.Second):
				pkt := &Packet{MsgType: 10000}
				conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
				err := pkt.writeTo(conn)
				if err != nil {
					log.Println(err)
					return
				}
			case p := <-ret.chOut:
				conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
				err := p.writeTo(conn)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}()
	return ret, nil
}

// Post post request
func (adapter *Adapter) Post(msgType int32, req proto.Message, requestID int32) {
	pkt := &Packet{MsgType: msgType, RequestID: requestID}
	if req != nil {
		d, _ := proto.Marshal(req)
		pkt.BodyList = append(pkt.BodyList, d)
	}
	adapter.chOut <- pkt
}
