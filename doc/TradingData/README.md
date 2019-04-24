# 交易数据

## 使用方式:

* 直接在代码中调用

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
* 运行独立的goshare server,通过HTTP访问,此种方式便于跨平台跨语言提供数据服务. 或者使用websocket连接接收实时数据推送服务