## 开始使用

### 下载安装golang
* go版本>=1.12.4
* [下载安装go最新版本](https://golang.google.cn/dl/)

### 安装
 ```
 go get -u github.com/mineralres/goshare
```
### 使用
* 直接在代码中调用goshare

```go
import (
  "github.com/mineralres/goshare"
)

func main() {
  var s goshare.DataSource
  symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "601398"}
  // 获取最新行情
  data, err := s.GetLastTick(&symbol)
  if err != nil {
    panic(err)
  }
}
```


* 运行API数据服务

在 cmd/server 下运行命令
```
go build -mod vendor
```
生成的 server 可执行文件是一个服务端程序，可独立运行。对外提供HTTP和websocket服务.
此程序默认使用19030端口，如果有360等防护软件提示，请允许，否则浏览器无法访问
```
./server
```
API示例: http://localhost:19030/v1/apiTest 

或者访问goshare官方数据API服务:https://goshare.cyconst.com/v1/apiTest

返回数据是JSON格式，601398(工商银行)的最新报价:

```json
{
    "success":true,
    "data":{
        "symbol":{
            "exchange":4,
            "code":"601398"
        },
        "time":1556004602,
        "milliseconds":0,
        "open":5.84,
        "high":5.85,
        "low":5.79,
        "close":5.81,
        "volume":1831041,
        "amount":1064240000,
        "position":0,
        "price":5.81,
        "preClose":5.83,
        "preSettlementPrice":0,
        "prePosition":0,
        "settlementPrice":0,
        "upperLimitPrice":6.41,
        "lowerLimitPrice":5.25,
        "preDelta":0,
        "delta":0,
        "averagePrice":0,
        "tradingDay":20190423,
        "orderBookList":[
            {
                "ask":5.81,
                "askVolume":21880,
                "bid":5.8,
                "bidVolume":14124
            },
            {
                "ask":5.82,
                "askVolume":26584,
                "bid":5.79,
                "bidVolume":53546
            },
            {
                "ask":5.83,
                "askVolume":33468,
                "bid":5.78,
                "bidVolume":54019
            },
            {
                "ask":5.84,
                "askVolume":41990,
                "bid":5.77,
                "bidVolume":12324
            },
            {
                "ask":5.85,
                "askVolume":31150,
                "bid":5.76,
                "bidVolume":23009
            }
        ],
        "name":"",
        "exercisePrice":0
    },
    "msg":"",
    "code":0
}
```