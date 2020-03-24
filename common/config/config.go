package config

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"sync"
)

var (
	m      sync.RWMutex
	inited bool

	// 默认配置器
	c = &configurator{}
)

type configurator struct {
	conf    config.Config
	appName string
}

func C() Configurator {
	return c
}

func (c *configurator) App(name string, config interface{}) (err error) {

	v := c.conf.Get(name)
	if v != nil {
		err = v.Scan(config)
	} else {
		err = fmt.Errorf("[App] 配置不存在，err：%s", name)
	}

	return
}




func (c *configurator) init(ops Options) (err error) {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[init] 配置已经初始化过")
		return
	}

	c.conf = config.NewConfig()
	c.appName = ops.AppName

	// 加载配置
	err = c.conf.Load(ops.Sources...)
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		log.Logf("[init] 侦听配置变动 ...")

		// 开始侦听变动事件
		watcher, err := c.conf.Watch()
		if err != nil {
			log.Fatal(err)
		}

		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatal(err)
			}

			log.Logf("[init] 侦听配置变动: %v", string(v.Bytes()))
		}
	}()

	// 标记已经初始化
	inited = true
	return
}

func Init(opts ...Option) {

	ops := Options{}
	for _, o := range opts {
		o(&ops)
	}

	c = &configurator{}

	c.init(ops)
}
