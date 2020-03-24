package util

import (
	"fmt"
	"github.com/liuhaogui/go-micro-mall/common/config"
)

func GetConsulAddress() string {

	consulCfg := Etcd{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)
}
