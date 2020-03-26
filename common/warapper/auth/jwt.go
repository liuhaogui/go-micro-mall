package auth

import (
	"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/util/log"
	"github.com/micro/micro/plugin"
	"net/http"
	"strings"
)

// JWTAuthWrapper JWT鉴权Wrapper
func JWTAuthWrapper(token *token.Token) plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("auth plugin received: ", r.URL.Path)
			// TODO 从配置中心动态获取白名单URL
			if r.URL.Path == "/user/login" || r.URL.Path == "/user/register" || r.URL.Path == "/user/test" ||
				r.URL.Path == "/metrics" || strings.HasPrefix(r.URL.Path, "/swagger") {
				log.Info("pass ", r.URL.Path)
				h.ServeHTTP(w, r)
				return
			}

			tokenstr := r.Header.Get("Authorization")
			userFromToken, e := token.Decode(tokenstr)

			if e != nil {
				log.Error("Unauthorized : ", r.URL.Path)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			log.Info("User Name : ", userFromToken.UserName)
			r.Header.Set("X-Example-Username", userFromToken.UserName)
			h.ServeHTTP(w, r)
		})
	}
}
