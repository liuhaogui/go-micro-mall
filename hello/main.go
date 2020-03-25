package main

import (
	"context"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	//"github.com/micro/go-micro/util/log"
	"github.com/liuhaogui/go-micro-mall/common/util/log"
	"github.com/opentracing/opentracing-go"
	"time"

	"github.com/liuhaogui/go-micro-mall/hello/handler"
	//"github.com/liuhaogui/go-micro-mall/hello/subscriber"


	example "github.com/liuhaogui/go-micro-mall/hello/proto/hello"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/liuhaogui/go-micro-mall/common/tracer"
	//"github.com/micro/cli"
	cfgUtil "github.com/liuhaogui/go-micro-mall/common/config/util"
)

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}

const (
	appName = "hello-srv"
)

var (
	appCfg = &cfgUtil.AppCfg{}
)

func init() {
	appCfg = cfgUtil.InitGetAppCfg(appName)
}

func main() {
	log.Info("start hello srv")
	t, io, err := tracer.NewTracer(appCfg.Name,  cfgUtil.GetJaegerAddress())
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	log.Info("cfgUtil.GetConsulAddress() : ",cfgUtil.GetConsulAddress())
	// New Service
	service := micro.NewService(
		micro.Name(appCfg.Name),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(t)),
		micro.Flags(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  cfgUtil.GetConsulAddress(),
		}),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Address(appCfg.Addr()),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.hello", service.Server(), new(subscriber.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
