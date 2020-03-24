package main

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/liuhaogui/go-micro-mall/api/user/handler"
	"github.com/liuhaogui/go-micro-mall/common/tracer"
	"github.com/liuhaogui/go-micro-mall/common/warapper/tracer/opentracing/gin2micro"
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/web"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/opentracing/opentracing-go"
	"time"

	"github.com/liuhaogui/go-micro-mall/common/token"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

const name = "go.micro.mall.api.user"
const consul_address = "127.0.0.1:8500"

func main() {
	// token
	token := &token.Token{}

	// tracer
	gin2micro.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//consul
	//reg := consul.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = []string{
	//		consul_address,
	//	}
	//})

	service := web.NewService(
		web.Name(name),
		web.Version("lastest"),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
		web.MicroService(grpc.NewService()),
		web.Flags(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  consul_address,
		}),
		web.Action(func(ctx *cli.Context) {
			token.InitConfig(consul_address, "micro", "config", "jwt-key", "key")
		}),
		//web.Registry(reg),
		web.Address(":8081"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	hystrix.DefaultTimeout = 5000

	sClient := hystrixplugin.NewClientWrapper()(service.Options().Service.Client())
	sClient.Init(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
		client.Retries(3),
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Log(req.Method(), retryCount, " client retry")
			return true, nil
		}),

	)

	//pub := micro.NewPublisher("/test", sClient)

	apiService := handler.New(sClient, token)
	router := gin.Default()
	r := router.Group("/user")
	r.Use(gin2micro.TracerWrapper)
	r.GET("/test", apiService.Anything)
	r.POST("/register", apiService.Create)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
