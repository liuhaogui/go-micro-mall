module github.com/liuhaogui/go-micro-mall

go 1.13

replace (
	github.com/gin-gonic/gin v1.6.0 => github.com/gin-gonic/gin v1.4.0
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d
	github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.5.1 => github.com/micro/go-plugins v1.2.0
)

require (
	github.com/golang/protobuf v1.3.3
	github.com/liuhaogui/go-micro-mall/common v0.0.0-20200325054819-642eaddd7d6c
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	google.golang.org/grpc v1.28.0 // indirect
)
