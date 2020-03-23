module github.com/liuhaogui/go-micro-mall/gateway

go 1.12

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-log/log v0.1.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.12
	github.com/liuhaogui/go-micro-mall v0.0.0-20200323124405-d73228717069 // indirect
	github.com/liuhaogui/go-micro-mall/user v0.0.0-20200323042051-1919988536d2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro v1.16.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.2.1
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/uber/jaeger-client-go v2.17.0+incompatible
)

replace (
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d
	github.com/micro/go-plugins v1.5.1 => github.com/micro/go-plugins v1.2.0

)
