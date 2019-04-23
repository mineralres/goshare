# 数据结构

为了扩展方便,goshare采用google的pb框架来作数据结构定义,数据结构都定义在pkg/pb/protos目录下


插件使用了修改版的gogofaster,原因如下:
* 生成的go代码效率更高
* 如果是struct的成员默认使用nullable，不使用pointer方式
* 生成的json tag采用 camalCase ，以方便提供HTTP服务

## 一 常用数据

* ExchangeType 交易所类型枚举
* Symbol 合约代码

## 二 交易数据
* Order 委托报单
* TradeReport 成交报告