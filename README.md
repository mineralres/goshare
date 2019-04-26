## 文档

<a href="https://goshare.cyconst.com" target="_blank">说明文档</a>

## goshare目标
 * goshare是一个开源的golang量化数据工具集合
 * 抓取 新浪,上证官方,东方财富 的历史k线或实时tick
 * 封装相关市场的交易接口
 * 集成量化交易框架
 * 量化交易webui控制台
 * 组建测试集群，充分利用多机多核
 
## Getting Started

### Installation
 ```
 go get -u github.com/mineralres/goshare
```
### Usage
```go
import (
  "github.com/mineralres/goshare"
)

func main() {
  var s goshare.DataSource
  symbol := pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "rb1805"}
  // 获取最新行情
  data, err := s.GetLastTick(&symbol)
  if err != nil {
    panic(err)
  }
}
```

# 数据来源
* 新浪财经
  * 股票最新报价
  * 期货最新报价
  * 50ETF期权最新报价
* 腾讯财经
* 上证所官网
  * 上证A股列表
  * 上证ETF期权列表
  * 分红送股信息
* 东方财富
  * 资金流向信息

# 策略平台

# 交易通道集成
 * [CTP接口]()
 * LTS

# UI

# 交流

![png](https://github.com/mineralres/goshare/blob/master/doc/images/goshare-group.png)

* QQ群 1018983692 
