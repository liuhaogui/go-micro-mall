package main

import (
	"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/warapper/auth"
	"github.com/liuhaogui/go-micro-mall/common/warapper/breaker/hystrix"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry"
	"log"
	"net"
	"net/http"
	"time"

	ph "github.com/afex/hystrix-go/hystrix"
	//"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/tracer"
	//"github.com/liuhaogui/go-micro-mall/common/warapper/auth"
	//"github.com/liuhaogui/go-micro-mall/common/warapper/breaker/hystrix"
	"github.com/liuhaogui/go-micro-mall/common/warapper/metrics/prometheus"
	"github.com/liuhaogui/go-micro-mall/common/warapper/tracer/opentracing/stdhttp"

	"github.com/micro/cli"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	opentracing "github.com/opentracing/opentracing-go"
)

const app_name = "api_gateway"
const consul_address = "127.0.0.1:8500"
func init() {
	token := &token.Token{}

	plugin.Register(cors.NewPlugin())

	plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(
			auth.JWTAuthWrapper(token),
		),
		plugin.WithInit(func(ctx *cli.Context) error {
			log.Println(consul_address)
			token.InitConfig(consul_address, "micro", "config", "jwt-key", "key")
			return nil
		}),
	))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("breaker"),
		plugin.WithHandler(
			hystrix.BreakerWrapper,
		),
	))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("metrics"),
		plugin.WithHandler(
			prometheus.MetricsWrapper,
		),
	))
}



func main() {
	stdhttp.SetSamplingFrequency(50)

	t, io, err := tracer.NewTracer(app_name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	hystrixStreamHandler := ph.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	//consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			consul_address,
		}
	})

	cmd.Init(
		micro.Name(app_name),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(reg),
		micro.Version("latest"),
	)



}
