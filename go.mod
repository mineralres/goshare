module github.com/mineralres/goshare

go 1.12

require (
	git.cyconst.com/mineralres/alps v0.0.0-20190425074432-60249d321486
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/gin-contrib/static v0.0.0-20190301062546-ed515893e96b
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.0
	github.com/ledongthuc/pdf v0.0.0-20190215042515-a147dfdf8062
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/go-micro v1.1.0
	github.com/micro/util v0.2.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/syndtr/goleveldb v1.0.0
	golang.org/x/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 // indirect
	golang.org/x/text v0.3.1
	google.golang.org/genproto v0.0.0-20190418145605-e7d98fc518a7
	google.golang.org/grpc v1.20.1
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
