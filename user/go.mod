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
	github.com/liuhaogui/go-micro-mall v0.0.9

	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/go-plugins/config/source/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.1.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	gopkg.in/olivere/elastic.v5 v5.0.84 // indirect
	gopkg.in/sohlich/elogrus.v2 v2.0.2 // indirect
)
