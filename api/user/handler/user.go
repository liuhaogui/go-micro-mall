package handler

import (
	//"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/liuhaogui/go-micro-mall/common/token"
	"github.com/liuhaogui/go-micro-mall/common/warapper/tracer/opentracing/gin2micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"net/http"

	helloS "github.com/liuhaogui/go-micro-mall/hello/proto/hello"
	userS "github.com/liuhaogui/go-micro-mall/user/proto/user"
)

const name = "go.micro.api.user"

// UserAPIService 服务
type UserAPIService struct {
	jwt    *token.Token
	helloC helloS.ExampleService
	userC  userS.UserService
	//pub    micro.Publisher
}

// New UserAPIService
func New(client client.Client, token *token.Token) *UserAPIService {
	return &UserAPIService{
		jwt:    token,
		helloC: helloS.NewExampleService("", client),
		userC:  userS.NewUserService("", client),
		//pub:    pub,
	}
}

// Anything 测试demo，调用hello服务和user两个服务
func (s *UserAPIService) Anything(c *gin.Context) {
	log.Log("Received Say.Anything API request")

	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Log("get context err")
	}

	//serviceClient := helloS.NewExampleService("go.micro.srv.example", client.DefaultClient)
	//res, err := serviceClient.Call(ctx, &helloS.Request{Name: "xuxu"})
	res, err := s.helloC.Call(ctx, &helloS.Request{Name: "xuxu"})
	if err != nil {
		log.Log("call error : ", err)
		//	c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Log(res)

	//s.pub.Publish(context.TODO(), &helloS.Message{Say: "你好"})

	// userres, err := s.userC.Ping(ctx, &userS.Request{})
	// if err != nil {
	// 	log.Log(err)
	// 	c.AbortWithError(http.StatusInternalServerError, err)
	// 	return
	// }
	// log.Log(userres)

	c.JSON(http.StatusOK, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

// Create 新建一个用户
/**
{
	"name":"xx",
	"email": "123.@qq.com",
	"tel":"tel1",
	"password":"d"
}
*/
func (s *UserAPIService) Create(c *gin.Context) {

	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Log("get context err")
	}
	var user userS.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("JWT decode failed"))
		return
	}

	_, err := s.userC.Create(ctx, &user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}
