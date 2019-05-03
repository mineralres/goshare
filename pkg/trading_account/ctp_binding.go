package tradingaccount

import (
	"log"
	"sync/atomic"
	"syscall"
	"unsafe"
)

const (
	bufferSize = 1024 * 1024
)

var (
	dll = syscall.MustLoadDLL("C:\\develop\\github.com\\mineralres\\ctp-binding\\build\\binding\\Release\\binding.dll")
	// dll             = syscall.MustLoadDLL("bindings.dll")
	ctpTradeCall         = dll.MustFindProc("CTP_Trade_Call")
	ctpTradePopupMessage = dll.MustFindProc("CTP_Trade_PopMessage")
	ctpMDCall            = dll.MustFindProc("CTP_MD_Call")
	ctpMDPopupMessage    = dll.MustFindProc("CTP_MD_PopMessage")
)

type session struct {
	goapi      uintptr
	gospi      int64
	cppapi     uintptr
	cppspi     uintptr
	inType     uint64
	inData     unsafe.Pointer
	inDataLen  uint64
	inParam1   int64
	outType    uint64
	outData    unsafe.Pointer
	outDataLen uint64
}

var (
	started           = false
	HandlerList       = make([]Handler, 1000)
	HandlerListOffset int64
	mdSpiList         = make([]MarketDataSpi, 1000)
	mdSpiListOffset   int64
)

func insertHandler(spi Handler) int64 {
	offset := atomic.LoadInt64(&HandlerListOffset)
	if offset > 900 {
		panic("offset > 900")
	}
	HandlerList[offset] = spi
	atomic.AddInt64(&HandlerListOffset, 1)
	return offset
}

func insertMarketDataSpi(spi MarketDataSpi) int64 {
	offset := atomic.LoadInt64(&mdSpiListOffset)
	if offset > 900 {
		panic("offset > 900")
	}
	mdSpiList[offset] = spi
	atomic.AddInt64(&mdSpiListOffset, 1)
	return offset
}

func init() {
	if started {
		panic("aleady started")
	}
	log.Println("unsafe.Sizeof(session)", unsafe.Sizeof(session{}))
	go func() {
		var s session
		buffer := make([]byte, bufferSize)
		s.outData = unsafe.Pointer(&buffer[0])
		for {
			ctpTradePopupMessage.Call(uintptr(unsafe.Pointer(&s)))
			onTraderPopupMessage(&s)
		}
	}()
	go func() {
		var s session
		buffer := make([]byte, bufferSize)
		s.outData = unsafe.Pointer(&buffer[0])
		for {
			ctpMDPopupMessage.Call(uintptr(unsafe.Pointer(&s)))
			onMDPopupMessage(&s)
		}
	}()
}
