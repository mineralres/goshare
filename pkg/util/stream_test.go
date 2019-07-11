package util

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Test_i(t *testing.T) {
	sig := make(chan bool)
	s := NewStream("stream", func(t int32, d []byte, isLast bool) {
		if isLast {
			return
		}
		if len(d) != 10 {
			log.Println(t, d, isLast)
			panic("")
		}
		log.Println(t, d, isLast)
	})
	b := make([]byte, 10)
	s.Write(99, b)
	s.Flush()
	s.Close()
	s = NewStream("stream", func(t int32, d []byte, isLast bool) {
		if isLast {
			sig <- true
			return
		}
		if len(d) != 10 {
			log.Println(t, d, isLast)
			panic("")
		}
		log.Println(t, d, isLast)
	})
	<-sig
}
