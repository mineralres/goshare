/v1/getLastTick
* API路径 /v1/getLastTick
* 传参数
```json
{
    "Exchange":4,
    "Code":"601398"
}
```
* 返回数据
```json
{
	"success": true,
	"data": {
		"symbol": {
			"exchange": 4,
			"code": "601398"
		},
		"time": 1556084924,
		"milliseconds": 0,
		"open": 5.82,
		"high": 5.84,
		"low": 5.74,
		"close": 5.79,
		"volume": 1675013,
		"amount": 971280000,
		"position": 0,
		"price": 5.79,
		"preClose": 5.81,
		"tradingDay": 20190424,
		"orderBookList": [{
			"ask": 5.8,
			"askVolume": 4043,
			"bid": 5.79,
			"bidVolume": 6998
		}],
		"name": "",
		"exercisePrice": 0
	},
	"msg": "",
	"code": 0
}
```
