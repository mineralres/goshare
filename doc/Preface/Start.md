## 开始使用

### 安装
 ```
 go get -u github.com/mineralres/goshare
```
### 使用
* 直接在代码中调用goshare

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


* 运行独立的goshare服务

在cmd/goshare下运行命令
```
go build
```
生成的goshare可执行文件是一个服务端程序，可独立运行。对外提供HTTP和websocket服务