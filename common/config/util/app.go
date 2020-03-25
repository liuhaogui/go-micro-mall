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
		panic(err)
	}

	log.Infof("[initCfg] 配置，cfg：%v", cfg)

	// log init
	esCfg := GetEsCfg()
	log.Info("esCfg: ", esCfg)
	if esCfg.Enabled {
		log.Info("init log es hook start")
		log.EsLogInit(appName, esCfg)
	}

	return
}

func GetConsulAddress() string {
	consulCfg := &Etcd{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)
}

func GetJaegerAddress() string {
	jaegerCfg := &Jaeger{}
	err := config.C().App("jaeger", jaegerCfg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%d", jaegerCfg.Host, jaegerCfg.Port)
}

func GetEsCfg() *ElasticSearch {
	esCfg := &ElasticSearch{}
	err := config.C().App("elasticsearch", esCfg)
	if err != nil {
		panic(err)
	}

	return esCfg
}
