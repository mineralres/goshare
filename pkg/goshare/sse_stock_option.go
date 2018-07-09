package goshare

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/mineralres/goshare/pkg/base"
	"github.com/mineralres/goshare/pkg/pb"
)

func getURLContent(url, referer string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	if referer != "" {
		req.Header.Set("Referer", referer)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}

// GetSSEStockOptionList 获取上证所网站的 50ETF个股期权列表
func (s *Service) GetSSEStockOptionList() ([]pb.SSEStockOption, error) {
	const url = "http://query.sse.com.cn/commonQuery.do?jsonCallBack=jsonpCallback77327&isPagination=true&expireDate=&securityId=&sqlId=SSE_ZQPZ_YSP_GGQQZSXT_XXPL_DRHY_SEARCH_L&pageHelp.pageSize=10000&pageHelp.pageNo=1&pageHelp.beginPage=1&pageHelp.cacheSize=1&pageHelp.endPage=5&_=1531102881526"
	str, err := getURLContent(url, "http://www.sse.com.cn/assortment/options/disclo/preinfo/")
	if err != nil {
		return nil, err
	}
	var rsp struct {
		ActionErrors []int  `json:"actionErrors"`
		Locale       string `json:"locale"`
		IsPagination string `json:"isPagination"`
		PageHelp     struct {
			BeginPage int `json:"beginPage"`
			CacheSize int `json:"cacheSize"`
			Data      []struct {
				EXERCISE_PRICE        string `json:"EXERCISE_PRICE"` // 行权价
				UPDATE_VERSION        string `json:"UPDATE_VERSION"` //
				OPTION_TYPE           string `json:"OPTION_TYPE"`
				DAILY_PRICE_UPLIMIT   string `json:"DAILY_PRICE_UPLIMIT"` // 涨停价
				TIMESAVE              string `json:"TIMESAVE"`
				DELISTFLAG            string `json:"DELISTFLAG"`
				START_DATE            string `json:"START_DATE"`
				EXPIRE_DATE           string `json:"EXPIRE_DATE"`
				CONTRACT_UNIT         string `json:"CONTRACT_UNIT"`
				CALL_OR_PUT           string `json"CALL_OR_PUT"`
				LMTORD_MAXFLOOR       string `json:"LMTORD_MAXFLOOR"`
				DELIVERY_DATE         string `json:"DELIVERY_DATE"`
				CHANGEFLAG            string `json:"CHANGEFLAG"`
				MKTORD_MAXFLOOR       string `json:"MKTORD_MAXFLOOR"`
				UNDERLYING_TYPE       string `json:"UNDERLYING_TYPE"`
				DAILY_PRICE_DOWNLIMIT string `json:"DAILY_PRICE_DOWNLIMIT"`
				ROUND_LOT             string `json:"ROUND_LOT"`
				SECURITY_CLOSEPX      string `json:"SECURITY_CLOSEPX"`
				SETTL_PRICE           string `json:"SETTL_PRICE"`
				CONTRACT_SYMBOL       string `json:"CONTRACT_SYMBOL"`
				NUM                   string `json:"NUM"`
				CONTRACT_ID           string `json:"CONTRACT_ID"`
				MARGIN_RATIO_PARAM1   string `json:"MARGIN_RATIO_PARAM1"`
				MARGIN_RATIO_PARAM2   string `json:"MARGIN_RATIO_PARAM2"`
				LMTORD_MINFLOOR       string `json:"LMTORD_MINFLOOR"`
				MKTORD_MINFLOOR       string `json:"MKTORD_MINFLOOR"`
				END_DATE              string `json:"END_DATE"`
				PRICE_LIMIT_TYPE      string `json:"PRICE_LIMIT_TYPE"`
				EXERCISE_DATE         string `json:"EXERCISE_DATE"`
				MARGIN_UNIT           string `json:"MARGIN_UNIT"`
				SECURITY_ID           string `json:"SECURITY_ID"`
				SECURITYNAMEBYID      string `json:"SECURITYNAMEBYID"`
				CONTRACTFLAG          string `json:"CONTRACTFLAG"`
				UNDERLYING_CLOSEPX    string `json:"UNDERLYING_CLOSEPX"`
			} `json:"data"`
		} `json:"pageHelp"`
	}
	start := strings.Index(str, "(")
	str = str[start+1 : len(str)-1]
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		return nil, err
	}
	var ret []pb.SSEStockOption
	for i := range rsp.PageHelp.Data {
		d := &rsp.PageHelp.Data[i]
		// log.Printf("合约编码[%s] 合约交易代码[%s] 合约简称[%s] 标的券名称及代码[%s] 类型[%s] 行权价[%s] 合约单位[%s] 期权行权日[%s] 行权交收日[%s] 到期日[%s] 新挂[%s] 涨停价[%s] 跌停价[%s] 前结算价[%s] 调整[%s]",
		// 	d.SECURITY_ID, d.CONTRACT_ID, d.CONTRACT_SYMBOL, d.SECURITYNAMEBYID, d.CALL_OR_PUT, d.EXERCISE_PRICE, d.CONTRACT_UNIT, d.EXERCISE_DATE,
		// 	d.DELIVERY_DATE, d.EXPIRE_DATE, d.CHANGEFLAG, d.DAILY_PRICE_UPLIMIT, d.DAILY_PRICE_DOWNLIMIT, d.SETTL_PRICE, d.CHANGEFLAG)
		var op pb.SSEStockOption
		op.ExercisePrice = d.EXERCISE_PRICE
		op.UpdateVersion = d.UPDATE_VERSION
		op.OptionType = d.OPTION_TYPE
		op.DailyPriceUpLimit = d.DAILY_PRICE_UPLIMIT
		op.TimeSave = d.TIMESAVE
		op.DELIST_Flag = d.DELISTFLAG
		op.StartDate = d.START_DATE
		op.ExpireDate = d.EXPIRE_DATE
		op.ContractUnit = d.CONTRACT_UNIT
		op.CallOrPut = d.CALL_OR_PUT
		op.LmtOrdMaxFloor = d.LMTORD_MAXFLOOR
		op.DeliveryDate = d.DELIVERY_DATE
		op.ChangeFlag = d.CHANGEFLAG
		op.MktOrdMaxFloor = d.MKTORD_MAXFLOOR
		op.UnderlyingClosePX = d.UNDERLYING_CLOSEPX
		op.UnderlyingType = d.UNDERLYING_TYPE
		op.DailyPriceDownLimit = d.DAILY_PRICE_DOWNLIMIT
		op.RoundLot = d.ROUND_LOT
		op.SecurityClosePX = d.SECURITY_CLOSEPX
		op.SettlPrice = d.SETTL_PRICE
		op.ContractSymbol = d.CONTRACT_SYMBOL
		op.Num = d.NUM
		op.ContractID = d.CONTRACT_ID
		op.MarginRatioParam1 = d.MARGIN_RATIO_PARAM1
		op.MarginRatioParam2 = d.MARGIN_RATIO_PARAM2
		op.LmtOrdMinFloor = d.LMTORD_MINFLOOR
		op.MktOrdMinFloor = d.MKTORD_MINFLOOR
		op.EndDate = d.END_DATE
		op.PriceLimitType = d.PRICE_LIMIT_TYPE
		op.ExerciseDate = d.EXERCISE_DATE
		op.MarginUnit = d.MARGIN_UNIT
		op.SecurityID = d.SECURITY_ID
		op.SecurityNameByID = d.SECURITYNAMEBYID
		op.ContractFlag = d.CONTRACTFLAG
		op.UnderlyingClosePX = d.UNDERLYING_CLOSEPX
		ret = append(ret, op)
	}
	return ret, nil
}

// GetSSEStockOptionTick 取所有行情
func (s *Service) GetSSEStockOptionTick(symbols []pb.Symbol) ([]pb.MarketDataSnapshot, error) {
	rets := []pb.MarketDataSnapshot{}
	all := "http://hq.sinajs.cn/list="
	for _, value := range symbols {
		all = all + "CON_OP_" + value.Code + ","
	}
	resp, err := http.Get(all)
	if err != nil {
		return nil, errors.New("ErrGetIndex")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	tickArr1 := strings.Split(string(body), ";")
	for _, v := range tickArr1 {
		tickArr := strings.Split(v, ",")
		ret := pb.MarketDataSnapshot{}
		if err == nil && len(tickArr) >= 42 {
			symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: tickArr[0][19:27]}
			ret.Symbol = symbol
			ret.Price = base.ParseFloat(tickArr[2])
			ret.Close = ret.Price
			ret.Open = base.ParseFloat(tickArr[9])
			ret.High = base.ParseFloat(tickArr[39])
			ret.Low = base.ParseFloat(tickArr[40])
			ret.Volume = (base.ParseFloat(tickArr[41]))
			ret.Amount = float64(base.ParseInt(tickArr[42]))
			ret.UpperLimitPrice = base.ParseFloat(tickArr[10])
			ret.LowerLimitPrice = base.ParseFloat(tickArr[11])
			var ob5 pb.OrderBook
			ob5.BidVolume = base.ParseFloat(tickArr[12])
			ob5.Bid = base.ParseFloat(tickArr[13])
			ob5.AskVolume = base.ParseFloat(tickArr[30])
			ob5.Ask = base.ParseFloat(tickArr[31])
			var ob4 pb.OrderBook
			ob4.BidVolume = base.ParseFloat(tickArr[14])
			ob4.Bid = base.ParseFloat(tickArr[15])
			ob4.AskVolume = base.ParseFloat(tickArr[28])
			ob4.Ask = base.ParseFloat(tickArr[29])
			var ob3 pb.OrderBook
			ob3.BidVolume = base.ParseFloat(tickArr[16])
			ob3.Bid = base.ParseFloat(tickArr[17])
			ob3.AskVolume = base.ParseFloat(tickArr[26])
			ob3.Ask = base.ParseFloat(tickArr[27])
			var ob2 pb.OrderBook
			ob2.BidVolume = base.ParseFloat(tickArr[18])
			ob2.Bid = base.ParseFloat(tickArr[19])
			ob2.AskVolume = base.ParseFloat(tickArr[24])
			ob2.Ask = base.ParseFloat(tickArr[25])
			var ob1 pb.OrderBook
			ob1.BidVolume = base.ParseFloat(tickArr[20])
			ob1.Bid = base.ParseFloat(tickArr[21])
			ob1.AskVolume = base.ParseFloat(tickArr[22])
			ob1.Ask = base.ParseFloat(tickArr[23])
			rets = append(rets, ret)
		}
	}
	return rets, nil
}
