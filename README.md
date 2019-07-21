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
* pkg/pb  protobuf生成的文件
* pkg/spider 抓取一些网站的数据
* pkg/tdxclient TDX数据接口
* pkg/util 实用函数

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
