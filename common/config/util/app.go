package util

import (
	"fmt"
	"github.com/liuhaogui/go-micro-mall/common/basic"
	"github.com/liuhaogui/go-micro-mall/common/config"
	"github.com/micro/go-plugins/config/source/grpc"
	"strconv"

	comCfg "github.com/liuhaogui/go-micro-mall/common/config"
	"github.com/micro/go-micro/util/log"
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

	log.Logf("[initCfg] 配置，cfg：%v", cfg)
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
