# CTP接口

[CTP综合交易平台](http://www.sfit.com.cn/)是上期技术开发的期货柜台系统，提供开放的API接入

go调用CTP接口的方式大概以下三种:


* FTDC协议方式.  CTP接口与服务器之间使用FTDC协议进行通讯，有一些爱好者对协议进行过收集研究:
  * [ftdc](https://github.com/haoziwlh/ftdc)
  * [量化交易从入门到放弃](https://zhuanlan.zhihu.com/p/38214049)
  * [期货交易数据交换协议](http://www.sse.com.cn/lawandrules/regulations/csrcannoun/c/3976298.pdf)

* [goctp](https://github.com/qerio/goctp) 开源库，目前只支持linux平台。缺点是用了cgo编译很慢


* goshare 通过syscall方式调用dll, 此种方式只支持windows平台