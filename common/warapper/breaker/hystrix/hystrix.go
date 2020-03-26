package hystrix

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	cfgUtil "github.com/liuhaogui/go-micro-mall/common/config/util"
	status_code "github.com/liuhaogui/go-micro-mall/common/http"
	"github.com/liuhaogui/go-micro-mall/common/util/log"
	"net/http"
)

func InitCfg() {
	cfg := cfgUtil.GetHystrixCfg()
	log.Infof("hystrix init config %s ", cfg)
	hystrix.DefaultTimeout = cfg.DefaultTimeout
	hystrix.DefaultMaxConcurrent = cfg.DefaultMaxConcurrent
	hystrix.DefaultVolumeThreshold = cfg.DefaultVolumeThreshold
	hystrix.DefaultSleepWindow = cfg.DefaultSleepWindow
	hystrix.DefaultErrorPercentThreshold = cfg.DefaultErrorPercentThreshold
}

// BreakerWrapper hystrix breaker
func BreakerWrapper(h http.Handler) http.Handler {
	InitCfg()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI
		log.Info(name)
		err := hystrix.Do(name, func() error {
			sct := &status_code.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
			h.ServeHTTP(sct.WrappedResponseWriter(), r)

			if sct.Status >= http.StatusBadRequest {
				str := fmt.Sprintf("status code %d ", sct.Status)
				log.Error(str,name)
				return errors.New(str)
			}
			return nil
		}, func(e error) error {
			//if e == hystrix.ErrCircuitOpen {
				//w.WriteHeader(http.StatusBadRequest)
				errResp := `{"msg":"Server error, please try again later", "error":"%s", "success":false}`
				errResp = fmt.Sprintf(errResp, e)

				w.Write([]byte(errResp))
			//}

			return e
		})
		if err != nil {
			log.Error("hystrix breaker err: ", err)
			return
		}
	})
}
