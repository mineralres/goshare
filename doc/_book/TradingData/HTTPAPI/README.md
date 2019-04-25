# HTTP接口


goshare server提供HTTP API，返回的数据为JSON格式.
```js
{
	"success": true, // 调用是否成功
	"data": null, // API返回的具体数据
	"msg": "", // 提示性消息
	"code": 0 // 返回的错误代码,成功的情况下一般为0
}
```

官方提供的数据API链接为: https://goshare.cyconst.com/v1/apiTest

* API路径为/v1/apiTest
* v1为API版本
* /apiTest为具体的数据接口
* 对于需要传参数的接口，分post和get两种方式，多数API只提供post方式，特别情况下会提供一些get形式的API.

# Websocket接口
用于推送数据

## API 列表

* 最新行情 [GetLastTick](./GetLastTick.md#)
* 上证50ETF期权列表 [sseStockOptionList](./SSEStockOptionList.md#)