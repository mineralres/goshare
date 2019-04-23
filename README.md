## goshare
 * [文档链接](https://goshare.cyconst.com)
 * goshare是一个开源的golang量化数据工具集合
 * 抓取 新浪,上证官方,东方财富 的历史k线或实时tick
 * 封装相关市场的交易接口
 * 集成量化交易框架
 * 量化交易webui控制台
 * 组建测试集群，充分利用多机多核
 
<!-- [START getstarted] -->
## Getting Started

### Installation
 ```
 go get -u github.com/mineralres/goshare
```
### Usage
```
import (
  "github.com/mineralres/goshare"
)

func main(){
  // 新浪数据源
  var s goshare.SinaSource
  symbol := pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "rb1805"}
  // 获取历史数据
  data, err := s.GetKData(&symbol, pb.PeriodType_M5, 19990101, 20180307, 1)
  if err != nil {
    panic(err)
  }
}
```

# 数据源

#### 新浪数据

* 期货实时数据
* 股票实时数据

#### 上证交易所官方

* A股列表
* 分红送股
* 50ETF 期权列表 GetSSEStockOptionList

#### 东方财富

* 资金流向  GetRealtimeMoneyTrendList


 ### 交易接口(待实现)
 * CTP
 * LTS
 

# FAQ
* QQ群 249379339 
