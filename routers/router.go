package router

import (
	//"fmt"
	"tianhe/config"
	"tianhe/middleware"

	//"oncall/routers/testRouter"
	"github.com/gin-gonic/gin"
)

const (
	HTTP_NOT_FOUND_CODE = 404
	HTTP_NO_LOGIN       = 1001
	HTTP_TOKEN_INVALID  = 1002
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Conf.Server.Mode)
	r := gin.Default()
	r.Use(middleware.RequestId())
	r.Use(middleware.Log())
	//初始化参数校验
	if err := middleware.TransInit("zh"); err != nil {
		return nil
	}

	InitRegisterRoute(r)

	return r
}
func InitRegisterRoute(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/tianhe")

	//值班规格管理
	//oncall := g.Group("/oncall")
	//registerOncallRouter(oncall)
	//用户管理
	user := g.Group("/user")
	registerUserRouter(user)

	health := g.Group("/checkhealth")
	registerHealthRouter(health)

	auth := g.Group("/auth")
	registerLoginRouter(auth)

	host := g.Group("host")
	registerHostRouter(host)
	//key := g.Group("/key")
	//registerKeyRouter(key)
	// 404处理
	r.Use(middleware.NoRoute)
	return g
}
