module github.com/liuhaogui/go-micro-mall/common

go 1.13

replace (
	github.com/gin-gonic/gin v1.6.0 => github.com/gin-gonic/gin v1.4.0
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d
	github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.5.1 => github.com/micro/go-plugins v1.2.0
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-gonic/gin v1.6.0
	github.com/liuhaogui/go-micro-mall v0.0.0-20200323124405-d73228717069
	github.com/liuhaogui/go-micro-mall/user v0.0.0-20200323042051-1919988536d2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/uber/jaeger-client-go v2.22.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
