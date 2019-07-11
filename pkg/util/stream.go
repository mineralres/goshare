package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type msg struct {
	t  int32
	d  []byte
	cb interface{}
}

// Stream stream
type Stream struct {
	directory string
	writeChan chan *msg
}

type handler func(t int32, d []byte, isLast bool)

// NewStream create new stream
func NewStream(directory string, h handler) *Stream {
	s := &Stream{}
	s.directory = directory
	os.MkdirAll(directory, os.ModePerm)
	s.writeChan = make(chan *msg, 100)
	go s.loadStream(h)
	return s
}

// LoadStream startup loading
func (s *Stream) loadStream(h handler) error {
	streamID := s.loadStreamID()
	file, err := os.Open(s.getStreamFileName(streamID))
	if err != nil {
		id := s.loadStreamID()
		go s.streamWriter(id)
		h(0, nil, true)
		return err
	}
	b := make([]byte, 100000)
	c := 0
	if err != nil {
		log.Println("open stream error", err)
	} else {
		log.Printf("open stream %s", s.getStreamFileName(streamID))
		if err != nil {
			panic(err)
		}
		for err == nil && file != nil {
			var t int32
			err = binary.Read(file, binary.LittleEndian, &t)
			if err != nil {
				break
			}
			var sz int32
			err = binary.Read(file, binary.LittleEndian, &sz)
			if sz == 0 || err != nil {
				break
			}
			if int(sz) > len(b) {
				b = make([]byte, sz)
			}
			n, err := file.Read(b[:sz])
			if err == nil && n == int(sz) {
				h(t, b[:sz], false)
				c++
			} else {
				err = errors.New("invalid size to sz")
			}
		}
	}
	log.Printf("loadStream completed, get msg count: %d, streamID:%d", c, streamID)
	id := s.loadStreamID()
	file.Close()
	go s.streamWriter(id)
	h(0, nil, true)
	return nil
}

func (s *Stream) getStreamFileName(id int) string {
	return fmt.Sprintf("%s/%d", s.directory, id)
}

func (s *Stream) loadStreamID() int {
	b, err := ioutil.ReadFile(s.directory + "/stream-id.txt")
	if err == nil {
		id, err := strconv.Atoi(string(b))
		if err == nil {
			return id
		}
	}
	return 0
}

func (s *Stream) writeStreamID(id int) {
	os.MkdirAll(s.directory, os.ModePerm)
	err := ioutil.WriteFile(s.directory+"/stream-id.txt", []byte(fmt.Sprintf("%d", id)), os.ModePerm)
	if err != nil {
		log.Println("writeStreamID error ", err)
	}
}

// Write write one message
func (s *Stream) Write(t int32, d []byte) {
	s.writeChan <- &msg{t: t, d: d}
}

// Close the stream
func (s *Stream) Close() {
	s.writeChan <- &msg{t: -2, d: nil}
}

// Reset switch to new stream
func (s *Stream) Reset() {
	s.writeChan <- &msg{t: -1}
}

// Flush steram
func (s *Stream) Flush() {
	sig := make(chan bool)
	s.writeChan <- &msg{t: -3, cb: func() {
		sig <- true
	}}
	<-sig
}

func (s *Stream) streamWriter(streamID int) {
	log.Println("StreamWriter started ", streamID)
	file, err := os.OpenFile(s.getStreamFileName(streamID), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	defer file.Close()
	if err == nil {
		for msg := range s.writeChan {
			if msg.t == -2 {
				log.Println("Close the stream")
				return
			}
			if msg.t == -1 {
				streamID++
				s.writeStreamID(streamID)
				log.Println("Reset stream to ", streamID)
				go s.streamWriter(streamID)
				return
			}
			if msg.t == -3 {
				if msg.cb != nil {
					log.Println("flush stream")
					msg.cb.(func())()
				}
				continue
			}
			binary.Write(file, binary.LittleEndian, &msg.t)
			sz := int32(len(msg.d))
			binary.Write(file, binary.LittleEndian, &sz)
			file.Write(msg.d)
		}
	} else {
		fmt.Println("open stream file  error:", err, " streamID:", streamID)
	}
}
