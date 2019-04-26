## 开始使用

### 下载安装golang
* go版本>=1.12.4
* [下载安装go最新版本](https://golang.google.cn/dl/)

### 安装
 ```
 go get -u github.com/mineralres/goshare
```



#### 运行API数据服务

在 cmd/server 下运行命令
```
go build -mod vendor
```
生成的 server 可执行文件是一个服务端程序，可独立运行。对外提供HTTP和websocket服务.
此程序默认使用18080端口，如果有360等防护软件提示，请允许，否则浏览器无法访问
```
./server
```
API示例:

* GET 方式取601398(工商银行)的最新行情 http://localhost:18080/v1/lastTick/SSE/601398

* POST方式取601398(工商银行)的最新行情 http://localhost:18080/v1/lastTick 传参数:
```js
{
    "exchange": "SSE",
    "code": "601398"
}
```

* goshare官方API测试:https://goshare.cyconst.com/v1/lastTick/SSE/601398

返回结果:

```js
{
	"symbol": {
		"exchange": "SSE",
		"code": "601398"
	},
	"time": "1556177404",
	"open": 5.79,
	"high": 5.8,
	"low": 5.71,
	"close": 5.73,
	"volume": 2259720,
	"amount": 1300360000,
	"price": 5.73,
	"preClose": 5.81,
	"upperLimitPrice": 6.39,
	"lowerLimitPrice": 5.23,
	"tradingDay": 20190425,
	"orderBookList": [{
		"ask": 5.73,
		"askVolume": 15128,
		"bid": 5.72,
		"bidVolume": 14673
	}, {
		"ask": 5.74,
		"askVolume": 16797,
		"bid": 5.71,
		"bidVolume": 44838
	}, {
		"ask": 5.75,
		"askVolume": 11945,
		"bid": 5.7,
		"bidVolume": 71088
	}, {
		"ask": 5.76,
		"askVolume": 19536,
		"bid": 5.69,
		"bidVolume": 10624
	}, {
		"ask": 5.77,
		"askVolume": 7861,
		"bid": 5.68,
		"bidVolume": 13338
	}]
}
```

#### 直接调用goshare

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
