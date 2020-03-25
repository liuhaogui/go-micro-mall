package util

import (
	"fmt"
	"github.com/liuhaogui/go-micro-mall/common/util/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
	elogrus "gopkg.in/sohlich/elogrus.v2"
	"time"
)

type ElasticSearch struct {
	Enabled  bool   `json:"enabled"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func EsLogInit(appname string, esCfg *ElasticSearch) {
	esUrl := fmt.Sprintf("%s://%s:%d", esCfg.Protocol, esCfg.Host, esCfg.Port)
	client, err := elastic.NewClient(elastic.SetURL(esUrl))
	if err != nil {
		Error("initLog elastic.NewClient error", err)
	}
	esIndex := fmt.Sprintf("%s-%s", appname, time.Now().Format("2006-01-02"))
	hook, err := elogrus.NewElasticHook(client, esCfg.Host, logrus.DebugLevel, esIndex)
	if err != nil {
		Error("initLog NewElasticHook error", err)
	}
	log.AddHook(hook)
}
