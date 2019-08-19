## 文档

* <a href="https://goshare.cyconst.com/doc" target="_blank">说明文档</a>
* <a href="https://goshare.cyconst.com/preview" target="_blank">演示</a> 

goshare是一个开源的golang量化数据工具集合。(注意:项目当前处于快速迭代阶段，不考虑向前兼容)

## goshare简介
  goshare 的最终目标是提供一个产品级策略平台,设计过程中考虑提供以下特性:

* 以微服务的方式设计架构和提供服务

* 集成专业和免费三方数据源,提供相对统一的访问接口

* 灵活的策略测试方式

* 集成多机联合测试功能

* 集成交易通道(CTP期货)

* 尽可能自动化常规的运维工作

* 方便使用的UI, [goshare-ui](https://github.com/mineralres/goshare-ui)

## 目录简介

### pkg

* pkg/api 尝试集成常用统一访问接口


* pkg/hub 集成交易通道
  * adapter.go 配合[ctp-binding](https://github.com/mineralres/ctp-binding) 的一个CTP交易客户端组件
  * demo.go  一个A股和中国期货模拟撮合的组件
  * pool.go 简易的adapter池
  * subscriber.go 与[ctp-binding](https://github.com/mineralres/ctp-binding)配合使用的订阅行情的组件
  * sync.go 一个同步访问的adapter封装


* pkg/pb  protobuf生成的文件


* pkg/spider 抓取一些网站的数据
  * east_money.go 东方财富相关
  * sina.go 新浪财经相关. 获取期货股票最新报价等
  * sse.go  上证所官网. 获取50ETF期权列表
  * xueqiu.go 雪球网. 如24小时内热度排名top10


* pkg/tdxclient TDX数据接口
  * sync_extern_client.go 扩展行情(期货，外盘等)
  * sync_quote_client.go 普通行情(上海,深圳股票)


* pkg/util 实用功能.
  * ldb_cache.go 简单的k/v缓存，用来缓存K线数组，最新报价, 合约信息等
  * send_main.go 发送邮件
  * stream.go 流式存储及回放. 类似redis的appendonly功能
  * tiny_gateway.go 一个简易的HTTP API网关

### cmd

* cmd/goshare  goshare演示项目
* cmd/util/monitor  一个简单的监控股票价格并发送邮件的程序

## 使用

```shell
go get -u github.com/mineralres/goshare
cd goshare/cmd/goshare
go build -mod vendor
./goshare
```
server需要监听一些端口提供http服务，如果有360等防护软件提示，请允许，否则浏览器无法访问

# 策略平台

# 关联项目
 * [CTP网关](https://github.com/mineralres/ctp-binding)
 * [UI](https://github.com/mineralres/goshare-ui)

# 交流

![png](https://github.com/mineralres/goshare/blob/master/doc/images/goshare-group.png)

* QQ群 1018983692 
