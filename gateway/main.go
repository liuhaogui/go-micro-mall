package main

import (
	"log"
	"net"
	"net/http"

	ph "github.com/afex/hystrix-go/hystrix"
	"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/tracer"
	"github.com/liuhaogui/go-micro-mall/common/warapper/auth"
	"github.com/liuhaogui/go-micro-mall/common/warapper/breaker/hystrix"
	"github.com/liuhaogui/go-micro-mall/common/warapper/metrics/prometheus"
	"github.com/liuhaogui/go-micro-mall/common/warapper/tracer/opentracing/stdhttp"

	"github.com/micro/cli"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	opentracing "github.com/opentracing/opentracing-go"
)

const name = "api-gateway"

func init() {
	token := &token.Token{}

	plugin.Register(cors.NewPlugin())

	plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(
			auth.JWTAuthWrapper(token),
		),
		plugin.WithFlag(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  "127.0.0.1:8500",
		}),
		plugin.WithInit(func(ctx *cli.Context) error {
			log.Println(ctx.String("consul_address"))
			token.InitConfig(ctx.String("consul_address"), "micro", "config", "jwt-key", "key")
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

	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	hystrixStreamHandler := ph.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	cmd.Init()

}
