package util

import (
	"log"
	"testing"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Test_ldb(t *testing.T) {
	l, err := NewLdbCache()
	if err != nil {
		panic(err)
	}
	var inst pb.Instrument
	err = l.SetInstrument(&inst)
	if err == nil {
		panic("should be empty symbol panic")
	}
	inst.Exchange = "SHFE"
	inst.Symbol = "ru2009"
	err = l.SetInstrument(&inst)
	if err != nil {
		panic(err)
	}
	i2, err := l.GetInstrument(inst.Exchange, inst.Symbol)
	if err != nil {
		panic(err)
	}
	if i2.Exchange != inst.Exchange || i2.Symbol != inst.Symbol {
		panic("")
	}
}
