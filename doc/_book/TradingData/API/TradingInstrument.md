# 2.2.1 交易合约

# 上证所50ETF期权列表

* 数据来源于上证所官网期权列表

* API路径/v1/sseStockOptionList

* 参考链接 https://goshare.cyconst.com/v1/sseStockOptionList

* 参数无

* 返回值
```js
{
	"list": [{
		"exercisePrice": "2.500",
		"updateVersion": "否",
		"optionType": "E",
		"dailyPriceUpLimit": "0.7593",
		"timeSave": "2019-04-23",
		"DELISTFlag": "否",
		"startDate": "20190311",
		"expireDate": "20190424",
		"contractUnit": "10000",
		"callOrPut": "认购",
		"lmtOrdMaxFloor": "30",
		"deliveryDate": "20190425",
		"changeFlag": "否",
		"mktOrdMaxFloor": "10",
		"underlyingType": "EBS",
		"dailyPriceDownLimit": "0.0001",
		"roundLot": "1",
		"securityClosePX": "0.4592",
		"settlPrice": "0.4630",
		"contractSymbol": "50ETF购4月2500",
		"num": "1",
		"contractID": "510050C1904M02500",
		"marginRatioParam1": "12",
		"marginRatioParam2": "7",
		"lmtOrdMinFloor": "1",
		"mktOrdMinFloor": "1",
		"endDate": "20190424",
		"priceLimitType": "N",
		"exerciseDate": "20190424",
		"marginUnit": "8185.6",
		"securityID": "10001771",
		"securityNameByID": "50ETF(510050)",
		"contractFlag": "否",
		"underlyingClosePX": "2.963"
	}, {
		"exercisePrice": "2.550",
		"updateVersion": "否",
		"optionType": "E",
		"dailyPriceUpLimit": "0.7093",
		"timeSave": "2019-04-23",
		"DELISTFlag": "否",
		"startDate": "20190228",
		"expireDate": "20190424",
		"contractUnit": "10000",
		"callOrPut": "认购",
		"lmtOrdMaxFloor": "30",
		"deliveryDate": "20190425",
		"changeFlag": "否",
		"mktOrdMaxFloor": "10",
		"underlyingType": "EBS",
		"dailyPriceDownLimit": "0.0001",
		"roundLot": "1",
		"securityClosePX": "0.4098",
		"settlPrice": "0.4130",
		"contractSymbol": "50ETF购4月2550",
		"num": "2",
		"contractID": "510050C1904M02550",
		"marginRatioParam1": "12",
		"marginRatioParam2": "7",
		"lmtOrdMinFloor": "1",
		"mktOrdMinFloor": "1",
		"endDate": "20190424",
		"priceLimitType": "N",
		"exerciseDate": "20190424",
		"marginUnit": "7685.6",
		"securityID": "10001751",
		"securityNameByID": "50ETF(510050)",
		"contractFlag": "否",
		"underlyingClosePX": "2.963"
	}]
}
```