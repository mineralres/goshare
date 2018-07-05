package goshare

// func init() {
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
// }

// // TestGetLastTick TestGetLastTick
// func TestGetLastTick2(t *testing.T) {
// 	var p Service
// 	symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "600000"}
// 	md, err := p.GetLastTick(&symbol)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if len(md.OrderBookList) == 0 {
// 		t.Error("获取行情盘口深度为空")
// 	}
// 	log.Printf("Tick[%s], Open[%.2f], High[%.2f], Low[%.2f], Close[%.2f]", md.Symbol.Code, md.Open, md.High, md.Low, md.Close)
// }

// func TestIndexTick2(t *testing.T) {
// 	//测试获取sina各种指数
// 	log.Printf("测试获取sina各种指数")

// 	m_index := map[string]string{
// 		"道琼斯指数": "int_dji",
// 		"上证指数":  "sh000001",
// 		"纳斯达克":  "int_nasdaq",
// 		"恒生指数":  "int_hangseng",
// 		"日经指数":  "b_TWSE",
// 		"新加坡指数": "b_FSSTI",
// 	}
// 	var p Service
// 	for key, views := range m_index {
// 		symbol := pb.Symbol{Exchange: pb.ExchangeType_INDEX, Code: views}
// 		md, err := p.GetLastTick(&symbol)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		if (md.Close) <= 0 {
// 			t.Error("获取行情为空")
// 		}
// 		md.Symbol.Code = key
// 		log.Printf("Tick[%s],Close[%.2f]", md.Symbol.Code, md.Close)
// 	}

// }

// func TestMainFutureTick(t *testing.T) {
// 	var p Service
// 	arr, err := p.GetMainFutureLastTick(pb.ExchangeType_SHFE)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if len(arr) == 0 {
// 		t.Error("取排名数据失败")
// 	}
// }
