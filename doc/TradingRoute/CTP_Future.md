# CTP接口

[CTP综合交易平台](http://www.sfit.com.cn/)是上期技术开发的期货柜台系统，提供开放的API接入

## go调用CTP接口的方式

大概以下三种:

* FTDC协议方式.  CTP接口与服务器之间使用FTDC协议进行通讯，有一些爱好者对协议进行过收集研究:
  * [ftdc](https://github.com/haoziwlh/ftdc)
  * [量化交易从入门到放弃](https://zhuanlan.zhihu.com/p/38214049)
  * [期货交易数据交换协议](http://www.sse.com.cn/lawandrules/regulations/csrcannoun/c/3976298.pdf)

* [goctp](https://github.com/qerio/goctp) 开源库，目前只支持linux平台。缺点是用了cgo编译很慢


* goshare 通过syscall方式调用dll, 此种方式只支持windows平台

## goshare测试环境

此测试环境为仿[simnow](http://www.simnow.com.cn/)测试平台

* 交易前置机地址 test.cyconst.com:41205

* BrokerID: ta

* [账号注册](https://test.cyconst.com/trade/futures.html?i=80104)

* 成交规则:
    * 1、期货交易按照交易所公布的买一卖一价对价成交；

    * 2、买入时：如果委托价大于等于卖一价，则成交，成交价为委托价、卖一价、最新价三价取中，如果委托价小于卖一价，不能成交，等待更优的行情才能成交；

    * 3、卖出时：如果委托价小于等于买一价，则成交，成交价为委托价、买一价、最新价三价取中，如果委托价大于买一价，不能成交，等待更优的行情才能成交。