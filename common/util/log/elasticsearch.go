package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
	elogrus "gopkg.in/sohlich/elogrus.v2"
	cfgUtil "github.com/liuhaogui/go-micro-mall/common/config/util"
	"time"
)

func EsLogInit(appname string, esCfg *cfgUtil.ElasticSearch) {
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
	AddHook(hook)

}
