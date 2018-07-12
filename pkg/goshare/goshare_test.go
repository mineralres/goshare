package goshare

import (
	"log"
	"testing"
	//"net/http"
	//"io/ioutil"
	//"strings"
	"github.com/mineralres/goshare/pkg/pb"
)

var s Service

func TestKData(t *testing.T) {
	symbol := pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "rb1805"}
	k, err := s.GetKData(&symbol, pb.PeriodType_M5, 19990101, 20180307, 1)
	if err != nil {
		t.Error(err)
	}
	if len(k.List) == 0 {
		t.Error("GetKData error")
	}
	// log.Println()
}

func TestGetLastTick(t *testing.T) {
	symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "600000"}
	md, err := s.GetLastTick(&symbol)
	if err != nil {
		t.Error(err)
	}
	if len(md.OrderBookList) == 0 {
		t.Error("获取行情盘口深度为空")
	}
	if md.Open == 0 {
		t.Error("md.Open == 0")
	}
	// log.Printf("Tick[%s], Open[%.2f], High[%.2f], Low[%.2f], Close[%.2f]", md.Symbol.Code, md.Open, md.High, md.Low, md.Close)
}

func TestIndexTick(t *testing.T) {
	return
	//测试获取sina各种指数
	log.Printf("测试获取sina各种指数")
	m_index := map[string]string{
		"道琼斯指数": "int_dji",
		"上证指数":  "sh000001",
		"纳斯达克":  "int_nasdaq",
		"恒生指数":  "int_hangseng",
		"日经指数":  "b_TWSE",
		"新加坡指数": "b_FSSTI",
	}
	for key, views := range m_index {
		symbol := pb.Symbol{Exchange: pb.ExchangeType_INDEX, Code: views}
		md, err := s.GetLastTick(&symbol)
		if err != nil {
			t.Error(err)
		}
		if (md.Close) <= 0 {
			t.Error("获取行情为空")
		}
		md.Symbol.Code = key
		log.Printf("Tick[%s],Close[%.2f]", md.Symbol.Code, md.Close)
	}
}

func TestOptionSSETick(t *testing.T) {
	return
	// 获取sina50etf期权的合约列表：
	// 同一个月份看涨和看跌: OP_UP_5100501804   OP_DOWN_5100501804
	sym := "OP_DOWN_5100501807"
	syms := GetSina50EtfSym(sym)
	for _, value := range syms {
		//log.Printf("Index: %d  Value: %s\n", index, value)
		symbol := pb.Symbol{Exchange: pb.ExchangeType_OPTION_SSE, Code: value}
		md, err := s.GetLastTick(&symbol)
		if err != nil {
			t.Error(err)
		}
		if (md.Close) <= 0 {
			t.Error("获取行情为空")
		}
		log.Printf("Tick[%s], Open[%.2f], High[%.2f], Low[%.2f], Close[%.2f]", md.Symbol.Code, md.Open, md.High, md.Low, md.Close)
	}
}

func TestRealtimeMoneyTrend(t *testing.T) {
	return
	size := 10
	l, err := s.GetRealtimeMoneyTrendList(size)
	if err != nil {
		t.Error(err)
	}
	if l == nil {
		t.Fatal("获取实时资金流向结果为空")
	}
	if len(l.List) == 0 {
		t.Error("获取实时资金流向结果为空")
	}
	if len(l.List) != size {
		t.Error("获取实时资金流向结果条数与预期不符合")
	}
}

func TestOptionSSEKline(t *testing.T) {
	return
	// 获取sina50etf期权的合约列表：
	// 同一个月份看涨和看跌: OP_UP_5100501804   OP_DOWN_5100501804
	log.Printf("测试获取期权k线")
	log.Printf("查询期权的合约k线，只有日k线")
	sym := "OP_DOWN_5100501807"
	syms := GetSina50EtfSym(sym)
	for _, value := range syms {
		log.Printf(" Value: %s\n", value)
		symbol := pb.Symbol{Exchange: pb.ExchangeType_OPTION_SSE, Code: value}
		k, err := s.GetKData(&symbol, pb.PeriodType_D1, 19990101, 20180707, 1)
		if err != nil {
			t.Error(err)
		}
		if len(k.List) == 0 {
			t.Error("GetKData error")
		}
		log.Printf("get result size : %d\n", len(k.List))
		// break
	}
}

func TestOption(t *testing.T) {
	log.Printf("测试获取期权数据：50etf对应的sina期权合约代码，详情，tick，kline")

	//期权标的+到期日
	sym := "OP_DOWN_5100501808"

	syms := GetSina50EtfSym(sym)
	//CON_OP_10001345
	//t型报价表
	all := "http://hq.sinajs.cn/list="
	for _, value := range syms {
		log.Printf(" sina 期权合约代码为: %s\n", value)
		all = all + value + ","
		// get tick
		symbol := pb.Symbol{Exchange: pb.ExchangeType_OPTION_SSE, Code: value}
		md, err := s.GetLastTick(&symbol)
		if err != nil {
			t.Error(err)
		}
		if (md.Close) <= 0 {
			t.Error("获取行情为空")
		}
		log.Printf("Tick[%s], Open[%.2f], High[%.2f], Low[%.2f], Close[%.2f]", md.Symbol.Code, md.Open, md.High, md.Low, md.Close)
		//get kline
		k, err := s.GetKData(&symbol, pb.PeriodType_D1, 19990101, 20180707, 1)
		if err != nil {
			t.Error(err)
		}
		log.Printf("获取k线数量为 size : %d\n", len(k.List))
		//break
	}

	//批量t型表的获取
	log.Printf("根据50etf期权到期月份，直接获取tick T型报价数据")
	allTick, _ := s.GetOptionSinaTick("1808")
	for _, v := range allTick {
		log.Printf("Tick[%s], Close[%.4f],preClose[%.4f]", v.Symbol.Code, v.Close, v.PreClose)
		// log.Println(v.OrderBookList[0])
		// log.Println(v)
		// log.Println(v.Position)
		//log.Printf("Tick[%s], ask[%.2f]，bid[%.2f]，Close[%.2f]", v.Symbol.Code, v.OrderBookList[0].Ask, v.OrderBookList[0].Bid, v.Close)
	}
	log.Printf("根据50etf期权到期月份，直接获取tick T型报价数据----2")
	allTick1, _ := s.GetOptionSinaTickMarket("1808")
	for _, val := range allTick1 {
		log.Printf("执行价[%.2f],name为%s,执行价[%.2f],name为%s,call 为%s,put 为%s", val.CallTk.ExercisePrice, val.CallTk.Name, val.PutTk.ExercisePrice, val.PutTk.Name, val.CallTk.Symbol.Code, val.PutTk.Symbol.Code)
	}

}

func TestGetSSEStockOptionList(t *testing.T) {
	return
	ret, err := s.GetSSEStockOptionList()
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) == 0 {
		t.Fatalf("上证股票期权列表为空")
	}

	var symbols []pb.Symbol
	for i := range ret {
		symbols = append(symbols, pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: ret[i].SecurityID})
	}
	mds, err := s.GetSSEStockOptionTick(symbols)
	log.Println(mds, err)
}
