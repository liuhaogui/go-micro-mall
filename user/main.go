package main

import (
	"time"

	"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/tracer"

	"github.com/liuhaogui/go-micro-mall/user/handler"
	user "github.com/liuhaogui/go-micro-mall/user/proto/user"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	opentracing "github.com/opentracing/opentracing-go"

)

const name = "go.micro.svc.user"

func main() {
	// token
	token := &token.Token{}
	var consulAddr string

	// tracer
	t, io, err := tracer.NewTracer(name, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// New Service
	service := grpc.NewService(
		micro.Name(name),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Flags(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  "127.0.0.1:8500",
		}),
		micro.Action(func(ctx *cli.Context) {
			consulAddr = ctx.String("consul_address")
			token.InitConfig(consulAddr, "micro", "config", "jwt-key", "key")
		}),
	)

	// Initialise service
	service.Init()

	db.Init(consulAddr)

	// Register Handler
	//user.RegisterUserHandler(service.Server(), new(handler.User))
	user.RegisterUserServiceHandler(service.Server(), handler.New(token))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("micro.svc.user.service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
