package tdxclient

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
)

// GetVolume 转码
func GetVolume(ivol uint32) float64 {
	logpoint := ivol >> (8 * 3)
	hleax := (ivol >> (8 * 2)) & 0xff // [2]
	lheax := (ivol >> 8) & 0xff       // [1]
	lleax := ivol & 0xff              // [0]

	dwEcx := int32(logpoint*2 - 0x7f)
	dwEdx := int32(logpoint*2 - 0x86)
	dwEsi := int32(logpoint*2 - 0x8e)
	dwEax := int32(logpoint*2 - 0x96)

	tmpEax := dwEcx
	if dwEcx < 0 {
		tmpEax = -dwEcx
	} else {
		tmpEax = dwEcx
	}
	// log.Println(logpoint, hleax, lheax, lleax, dwEcx, dwEdx, dwEsi, dwEax, tmpEax)

	dblXmm6 := 0.0
	dblXmm6 = math.Pow(2.0, float64(tmpEax))
	if dwEcx < 0 {
		dblXmm6 = 1.0 / dblXmm6
	}
	var dblXmm4 float64
	if hleax > 0x80 {
		tmpdblXmm3 := 0.0
		dwtmpeax := dwEdx + 1
		tmpdblXmm3 = math.Pow(2.0, float64(dwtmpeax))
		dblXmm0 := float64(math.Pow(2.0, float64(dwEdx)) * 128.0)
		dblXmm0 += float64(hleax&0x7f) * float64(tmpdblXmm3)
		dblXmm4 = (dblXmm0)
	} else {
		dblXmm0 := 0.0
		if dwEdx >= 0 {
			dblXmm0 = math.Pow(2.0, float64(dwEdx)) * float64(hleax)
		} else {
			dblXmm0 = (1 / math.Pow(2.0, float64(dwEdx))) * float64(hleax)
		}
		dblXmm4 = (dblXmm0)
	}
	dblXmm3 := math.Pow(2.0, float64(dwEsi)) * float64(lheax)
	dblXmm1 := math.Pow(2.0, float64(dwEax)) * float64(lleax)
	if (hleax & 0x80) > 0 {
		dblXmm3 *= 2.0
		dblXmm1 *= 2.0
	}
	return dblXmm6 + float64(dblXmm4) + float64(dblXmm3) + float64(dblXmm1)
}

func getDateTime(category TdxKlineType, buffer []byte) (int, int, int, int, int) {
	year := 0
	month := 0
	day := 0
	hour := 15
	minute := 0
	reader := bytes.NewReader(buffer)
	if category < 4 || category == 7 || category == 8 {
		var zipday, tminutes uint16
		binary.Read(reader, binary.LittleEndian, &zipday)
		binary.Read(reader, binary.LittleEndian, &tminutes)
		year = int(zipday>>11) + 2004
		month = int((zipday % 2048) / 100)
		day = int(zipday%2048) % 100

		hour = int(tminutes / 60)
		minute = int(tminutes) % 60
	} else {
		var zipday uint32
		binary.Read(reader, binary.LittleEndian, &zipday)

		year = int(zipday / 10000)
		month = int((zipday % 10000) / 100)
		day = int(zipday) % 100
	}
	return year, month, day, hour, minute
}

func toTdxMarket(ex string) uint8 {
	switch ex {
	case "SHFE":
		return 30
	case "CZCE":
		return 28
	case "DCE":
		return 29
	case "CFFEX":
		return 47
	}
	return 0
}

func read(conn net.Conn) (*tdxResponse, error) {
	conn.SetReadDeadline(time.Now().Add(time.Second * 15))
	var h header
	headerBuf := make([]byte, 16)
	readed, err := io.ReadFull(conn, headerBuf)
	if err != nil {
		return nil, err
	}
	err = unmarshal(headerBuf, &h)
	if err != nil {
		log.Println("invalid ex tdx header")
		return nil, err
	}

	body := make([]byte, h.ZipSize)
	readed, err = io.ReadFull(conn, body)
	if err != nil {
		return nil, err
	}
	if h.ZipSize != h.UnzipSize {
		// log.Println("需要解压")
		r, err := zlib.NewReader(bytes.NewReader(body))
		if err != nil {
			log.Println(err, len(body), h, readed)
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
	return &tdxResponse{h: h, body: body}, nil
}

func getPrice(data []byte, pos int) (int, int) {
	var posByte uint32
	posByte = 6
	bdata := data[pos]
	var intdata int
	intdata = int(bdata & 0x3f)
	sign := false
	if bdata&0x40 > 0 {
		sign = true
	}
	if bdata&0x80 > 0 {
		for {
			pos++
			bdata = data[pos]
			intdata += int(bdata&0x7f) << posByte
			posByte += 7
			// log.Println(posByte, bdata, intdata, sign, pos, bdata&0x80)
			if bdata&0x80 > 0 {
				// pass
			} else {
				break
			}
		}
	}
	pos++
	if sign {
		intdata = -intdata
	}
	return intdata, pos
}

func calPrice1000(basePrice, diff int) float64 {
	return float64(basePrice+diff) / 1000
}

// ToKlineArr map to gspb kline series
func ToKlineArr(l []*SecurityBar) ([]*pb.Kline, error) {
	var ret []*pb.Kline
	for _, sb := range l {
		k := &pb.Kline{}
		k.Open = sb.Open
		k.High = sb.High
		k.Low = sb.Low
		k.Close = sb.Close
		k.Volume = int32(sb.Vol)
		k.Amount = float64(sb.DBVol)
		t, err := time.Parse("20060102 15:04", fmt.Sprintf("%d%02d%02d %02d:%02d", sb.Year, sb.Mon, sb.Day, sb.Hour, sb.Minute))
		if err == nil {
			k.Time = t.Unix()
		}
		ret = append(ret, k)
	}
	return ret, nil
}

func toTdxPeriod(src pb.PeriodType) TdxKlineType {
	switch src {
	case pb.PeriodType_M1:
		return TdxKlineType_EXHQ_1MIN
	case pb.PeriodType_M5:
		return TdxKlineType_5MIN
	case pb.PeriodType_M15:
		return TdxKlineType_15MIN
	case pb.PeriodType_M30:
		return TdxKlineType_30MIN
	case pb.PeriodType_H1:
		return TdxKlineType_1HOUR
	case pb.PeriodType_D1:
		return TdxKlineType_DAILY
	}
	return TdxKlineType_YEARLY
	// TdxKlineType_WEEKLY    = 5
	// TdxKlineType_MONTHLY   = 6
	// TdxKlineType_EXHQ_1MIN = 7
	// TdxKlineType_RI_K      = 9
	// TdxKlineType_3MONTH    = 10
	// TdxKlineType_YEARLY    = 11
}

func toDepthsMarketDataArr(l []*SecurityQuote) ([]*pb.MarketDataSnapshot, error) {
	var ret []*pb.MarketDataSnapshot
	for _, q := range l {
		var s pb.MarketDataSnapshot
		s.Price = float64(q.Price) / 100
		s.Open = float64(q.OpenDiff+q.Price) / 100
		s.High = float64(q.HighDiff+q.Price) / 100
		s.Low = float64(q.LowDiff+q.Price) / 100
		s.Close = float64(q.LastCloseDiff+q.Price) / 100
		s.Volume = int32(q.Vol)
		s.VolumeDelta = int32(q.CurVol)
		s.Amount = q.Amount

		ob := &pb.OrderBook{}
		ob.Ask = float64(q.Price+q.Ask1) / 100
		ob.Bid = float64(q.Price+q.Bid1) / 100
		ob.AskVolume = int32(q.AskVol1)
		ob.BidVolume = int32(q.BidVol1)
		s.Depths = append(s.Depths, ob)

		ob = &pb.OrderBook{}
		ob.Ask = float64(q.Price+q.Ask2) / 100
		ob.Bid = float64(q.Price+q.Bid2) / 100
		ob.AskVolume = int32(q.AskVol2)
		ob.BidVolume = int32(q.BidVol2)
		s.Depths = append(s.Depths, ob)

		ob = &pb.OrderBook{}
		ob.Ask = float64(q.Price+q.Ask3) / 100
		ob.Bid = float64(q.Price+q.Bid3) / 100
		ob.AskVolume = int32(q.AskVol3)
		ob.BidVolume = int32(q.BidVol3)
		s.Depths = append(s.Depths, ob)

		ob = &pb.OrderBook{}
		ob.Ask = float64(q.Price+q.Ask4) / 100
		ob.Bid = float64(q.Price+q.Bid4) / 100
		ob.AskVolume = int32(q.AskVol4)
		ob.BidVolume = int32(q.BidVol4)
		s.Depths = append(s.Depths, ob)

		ob = &pb.OrderBook{}
		ob.Ask = float64(q.Price+q.Ask5) / 100
		ob.Bid = float64(q.Price+q.Bid5) / 100
		ob.AskVolume = int32(q.AskVol5)
		ob.BidVolume = int32(q.BidVol5)
		s.Depths = append(s.Depths, ob)

		ret = append(ret, &s)
	}
	return ret, nil
}
