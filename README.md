## 文档

<a href="https://goshare.cyconst.com" target="_blank">说明文档</a>

## goshare简介
 * goshare是一个开源的golang量化数据工具集合
 * 提供自建数据API服务(生产环境)
 * 三方数据. 抓取新浪,上证官方,东方财富等相关数据源的数据(测试环境)
 * 集成量化策略测试和交易框架
 * 封装相关市场的交易接口
 * webui控制台
 * 自建策略测试集群,提升策略测试效率
 
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

# UI
- UI项目库 [goshare-ui](https://github.com/mineralres/goshare-ui)
- 演示地址 [https://admin.cyconst.com](https://admin.cyconst.com) 

# 数据库
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
 * [CTP接口](https://github.com/mineralres/ctp-binding)
 * LTS

# 交流

![png](https://github.com/mineralres/goshare/blob/master/doc/images/goshare-group.png)

* QQ群 1018983692 
