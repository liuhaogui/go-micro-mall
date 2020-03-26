module github.com/liuhaogui/go-micro-mall/user

go 1.13

replace (
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

	github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.5.1 => github.com/micro/go-plugins v1.2.0
)

require (
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/liuhaogui/go-micro-mall v0.0.28

	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/go-plugins/config/source/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.1.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)
