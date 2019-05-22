package tdxclient

import (
	"bytes"
	"encoding/binary"
	"log"
	"math"

	"github.com/mineralres/goshare/pkg/pb"
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
	log.Println(logpoint, hleax, lheax, lleax, dwEcx, dwEdx, dwEsi, dwEax, tmpEax)

	dbl_xmm6 := 0.0
	dbl_xmm6 = math.Pow(2.0, float64(tmpEax))
	if dwEcx < 0 {
		dbl_xmm6 = 1.0 / dbl_xmm6
	}
	var dbl_xmm4 float64
	if hleax > 0x80 {
		tmpdbl_xmm3 := 0.0
		dwtmpeax := dwEdx + 1
		tmpdbl_xmm3 = math.Pow(2.0, float64(dwtmpeax))
		dbl_xmm0 := float64(math.Pow(2.0, float64(dwEdx)) * 128.0)
		dbl_xmm0 += float64(hleax&0x7f) * float64(tmpdbl_xmm3)
		dbl_xmm4 = (dbl_xmm0)
	} else {
		dbl_xmm0 := 0.0
		if dwEdx >= 0 {
			dbl_xmm0 = math.Pow(2.0, float64(dwEdx)) * float64(hleax)
		} else {
			dbl_xmm0 = (1 / math.Pow(2.0, float64(dwEdx))) * float64(hleax)
		}
		dbl_xmm4 = (dbl_xmm0)
	}
	dbl_xmm3 := math.Pow(2.0, float64(dwEsi)) * float64(lheax)
	dbl_xmm1 := math.Pow(2.0, float64(dwEax)) * float64(lleax)
	if (hleax & 0x80) > 0 {
		dbl_xmm3 *= 2.0
		dbl_xmm1 *= 2.0
	}
	dbl_ret := dbl_xmm6 + float64(dbl_xmm4) + float64(dbl_xmm3) + float64(dbl_xmm1)
	return dbl_ret
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

func toTdxMarket(ex pb.ExchangeType) uint8 {
	switch ex {
	case pb.ExchangeType_SHFE:
		return 30
	case pb.ExchangeType_CZCE:
		return 28
	case pb.ExchangeType_DCE:
		return 29
	case pb.ExchangeType_CFFEX:
		return 47
	}
	return 0
}
