package util

import (
	"fmt"
	"github.com/liuhaogui/go-micro-mall/common/basic"
	"github.com/liuhaogui/go-micro-mall/common/config"
	"github.com/micro/go-plugins/config/source/grpc"
	"strconv"

	comCfg "github.com/liuhaogui/go-micro-mall/common/config"
	"github.com/liuhaogui/go-micro-mall/common/util/log"
)

type AppCfg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (a *AppCfg) Addr() string {
	return a.Address + ":" + strconv.Itoa(a.Port)
}

func InitGetAppCfg(appName string) (cfg *AppCfg) {
	cfg = &AppCfg{}
	source := grpc.NewSource(
		grpc.WithAddress(comCfg.Config_srv_address),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		log.Errorf("load app[%s] config error  %s ", appName, err)
		panic(err)
	}

	log.Infof("[initCfg] 配置，cfg：%v", cfg)

	// log init
	esCfg := GetEsCfg()
	log.Info("esCfg: ", esCfg)
	if esCfg.Enabled {
		log.Info("init log es hook start")
		EsLogInit(appName, esCfg)
	}

	return
}

func GetConsulAddress() string {
	consulCfg := &Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		log.Error("load consul config error ", err)
		panic(err)
	}

	return fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)
}

func GetJaegerAddress() string {
	jaegerCfg := &Jaeger{}
	err := config.C().App("jaeger", jaegerCfg)
	if err != nil {
		log.Error("load jaeger config error ", err)
		panic(err)
	}

	return fmt.Sprintf("%s:%d", jaegerCfg.Host, jaegerCfg.Port)
}

func GetEsCfg() *ElasticSearch {
	esCfg := &ElasticSearch{}
	err := config.C().App("elasticsearch", esCfg)
	if err != nil {
		log.Error("load elasticsearch config error ", err)
		panic(err)
	}

	return esCfg
}

func GetHystrixCfg() *Hystrix {
	hystrixCfg := &Hystrix{}
	err := config.C().App("hystrix", hystrixCfg)
	log.Info("hystrixCfg load config : ", hystrixCfg)
	if err != nil {
		log.Error("load hystrix config error ", err)
		return &Hystrix{
			DefaultTimeout:               1000,
			DefaultMaxConcurrent:         10,
			DefaultVolumeThreshold:       20,
			DefaultSleepWindow:           5000,
			DefaultErrorPercentThreshold: 50,
		}
	}

	return hystrixCfg
}
