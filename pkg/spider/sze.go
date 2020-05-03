package spider

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	pb "github.com/mineralres/protos/src/go/goshare"
)

// SZE sze
type SZE struct {
}

// StockList instrument list
func (sze *SZE) StockList(toFillPriceInfo bool) ([]*pb.Instrument, error) {
	excelFileName := fmt.Sprintf("SZE_STOCK_LIST_%s.xlsx", time.Now().Format("20060102"))
	if _, err := os.Stat(excelFileName); os.IsNotExist(err) {
		const url = `http://www.szse.cn/api/report/ShowReport?SHOWTYPE=xlsx&CATALOGID=1110&TABKEY=tab1&random=0.45654626027265355`
		log.Println("下载文件", excelFileName)
		err := downloadFile(url, excelFileName, "")
		if err != nil {
			panic(err)
		}
		log.Println("下载完成")
	}

	f, err := excelize.OpenFile(excelFileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := f.GetRows("A股列表")
	var list []*pb.Instrument
	for _, row := range rows {
		if len(row) <= 7 {
			log.Println(row)
			continue
		}
		if row[0] == "公司代码" {
			continue
		}
		var item pb.Instrument
		item.Exchange = "SZE"
		item.Symbol = row[5]
		item.Symbol = strings.TrimSpace(item.Symbol)
		opendate := row[7]
		opendate = strings.Replace(opendate, "-", "", -1)
		d, err := strconv.Atoi(opendate)
		if err != nil || d == 0 {
			log.Println(err, row)
			continue
		}
		item.OpenDate = int32(d)

		item.Product = "SZA"
		item.PriceTick = 0.01
		item.Multiple = 1
		item.Name = row[1]
		item.UpdateTime = time.Now().Unix()
		item.ProductType = int32(pb.ProductType_STOCK)
		item.IsCloseTodayAllowed = false
		item.MaxLimitOrderVolume = 1000000
		item.MaxMarketOrderVolume = 1000000
		item.MinBuyVolume = 100
		item.MinLimitOrderVolume = 100
		item.MinMarketOrderVolume = 100
		item.MinSellVolume = 100
		item.IsTrading = true

		if strings.Contains(item.Name, "ST") || strings.Contains(item.Name, "*ST") {
			item.ProductClass = "ST"
		}

		ud, _ := strconv.Atoi(time.Now().Format("20060102"))
		item.TradingDay = int32(ud)

		list = append(list, &item)
	}
	log.Println("深圳股票数量", len(list))
	if toFillPriceInfo {
		fillPriceInfo("SZE", list)
	}
	return list, nil
}

func showProcess(prefix string, current, total int, currentSymbol string) {
	log.Printf("%s: 当前进度[%d/%d, %s]", prefix, current, total, currentSymbol)
}

func fillPriceInfo(ex string, list []*pb.Instrument) {
	log.Printf("填充股票价格[%d]", len(list))
	var sina Sina
	for _, inst := range list {
		mds, err := sina.GetLastTick(ex, inst.Symbol)
		if err != nil {
			continue
		}
		if mds.UpperLimit > 0 {
			inst.UpperLimit = mds.UpperLimit
		}
		if mds.LowerLimit > 0 {
			inst.LowerLimit = mds.LowerLimit
		}
		// inst.PrePosition = int32(mds.PreClose * float64(mds.Volume) / 100)
		// showProcess("填充股票价格", i, len(list), inst.Symbol)
	}
}
