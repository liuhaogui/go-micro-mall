package main

import (
	"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/util/log"
	"github.com/liuhaogui/go-micro-mall/common/warapper/auth"
	"github.com/liuhaogui/go-micro-mall/common/warapper/breaker/hystrix"
	"net"
	"net/http"
	"time"

	//"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/tracer"
	//"github.com/liuhaogui/go-micro-mall/common/warapper/auth"
	//"github.com/liuhaogui/go-micro-mall/common/warapper/breaker/hystrix"
	"github.com/liuhaogui/go-micro-mall/common/warapper/metrics/prometheus"
	"github.com/liuhaogui/go-micro-mall/common/warapper/tracer/opentracing/stdhttp"

	ph "github.com/afex/hystrix-go/hystrix"
	cfgUtil "github.com/liuhaogui/go-micro-mall/common/config/util"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	opentracing "github.com/opentracing/opentracing-go"
)

const appName = "api-gateway"

var (
	appCfg = &cfgUtil.AppCfg{}
)

func init() {
	appCfg = cfgUtil.InitGetAppCfg(appName)
}

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
			Value:  cfgUtil.GetConsulAddress(),
		}),
		plugin.WithInit(func(ctx *cli.Context) error {
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

	t, io, err := tracer.NewTracer(appCfg.Name, cfgUtil.GetJaegerAddress())
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	hystrixStreamHandler := ph.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	cmd.Init(
		micro.Name(appCfg.Name),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Address(appCfg.Addr()),
		//micro.Version("latest"),
	)

}
