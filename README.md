## 文档

<a href="https://goshare.cyconst.com/doc" target="_blank">说明文档</a>

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


## 运行

```shell
go get -u github.com/mineralres/goshare
cd goshare/cmd/goshare
go build -mod vendor
./goshare
```
server需要监听一些端口提供http服务，如果有360等防护软件提示，请允许，否则浏览器无法访问

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
 * [CTP网关](https://github.com/mineralres/ctp-binding)
 * LTS

# 交流

![png](https://github.com/mineralres/goshare/blob/master/doc/images/goshare-group.png)

* QQ群 1018983692 
