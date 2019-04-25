# getLastTick 

* 从新浪财经相关页面获取的最新盘口价

* API路径 /v1/getLastTick

* 传参数,post方式
```js
{
    "Exchange":4,
    "Code":"601398"
}
```

* 返回数据
```js
{
	"success": true,
	"data": {
		"symbol": {
			"exchange": 4,
			"code": "601398"
		},
		"time": 1556084924, // 更新时间time_t
		"milliseconds": 0, // 毫秒
		"open": 5.82, // 开盘价
		"high": 5.84, // 最高价
		"low": 5.74, // 最低价
		"close": 5.79, // 收盘价
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
