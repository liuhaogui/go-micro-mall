package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/liuhaogui/go-micro-mall/order/handler"
	"github.com/liuhaogui/go-micro-mall/order/subscriber"

	user "github.com/liuhaogui/go-micro-mall/order/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("micro.svc.order.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.svc.order.service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
