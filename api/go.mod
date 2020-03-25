module github.com/liuhaogui/go-micro-mall/api

go 1.13

replace (
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d
	github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.5.1 => github.com/micro/go-plugins v1.2.0
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-gonic/gin v1.6.1
	github.com/liuhaogui/go-micro-mall v0.0.1
	github.com/liuhaogui/go-micro-mall/common v0.0.0-20200325054819-642eaddd7d6c
	github.com/liuhaogui/go-micro-mall/user v0.0.0-20200323131939-9a71bea8fde9
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
)
