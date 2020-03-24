package main

import (
	"context"
	"github.com/liuhaogui/go-micro-mall/common/tracer"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/opentracing/opentracing-go"
	"time"

	"github.com/liuhaogui/go-micro-mall/example/handler"
	//"github.com/liuhaogui/go-micro-mall/example/subscriber"

	example "github.com/liuhaogui/go-micro-mall/example/proto/example"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	//"github.com/micro/cli"
)

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}

var name = "go.micro.mall.srv.hello"

func main() {
	t, io, err := tracer.NewTracer(name, "")
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
	service := micro.NewService(
		micro.Name(name),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(t)),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		//micro.Registry(reg),
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
