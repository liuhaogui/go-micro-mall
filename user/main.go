package main

import (
	cfgUtil "github.com/liuhaogui/go-micro-mall/common/config/util"
	"time"

	"github.com/liuhaogui/go-micro-mall/user/handler"
	user "github.com/liuhaogui/go-micro-mall/user/proto/user"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/tracer"

	db "github.com/liuhaogui/go-micro-mall/user/model"
)

const (
	appName      = "user-srv"
)

var (
	appCfg  = &cfgUtil.AppCfg{}
)

func init()  {
	appCfg = cfgUtil.InitGetAppCfg(appName)
}

func main() {
	// token
	token := &token.Token{}
	var consulAddr string

	// tracer
	t, io, err := tracer.NewTracer(appCfg.Name, cfgUtil.GetJaegerAddress())
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//reg := consul.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = []string{
	//		"127.0.0.1:8500",
	//	}
	//})

	// New Service
	service := grpc.NewService(
		micro.Name(appCfg.Name),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Flags(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  cfgUtil.GetConsulAddress(),
		}),
		micro.Action(func(ctx *cli.Context) {
			token.InitConfig(ctx.String("consul_address"), "micro", "config", "jwt-key", "key")
		}),
		//micro.Registry(reg),
		micro.Address(appCfg.Addr()),
	)

	// Initialise service
	service.Init()
	db.Init(consulAddr)

	// Register Handler
	user.RegisterUserServiceHandler(service.Server(), handler.New(token))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
