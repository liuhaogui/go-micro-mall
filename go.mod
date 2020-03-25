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

	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	google.golang.org/grpc v1.28.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/olivere/elastic.v5 v5.0.84
	gopkg.in/sohlich/elogrus.v2 v2.0.2
)
