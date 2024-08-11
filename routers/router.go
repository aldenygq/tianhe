package router

import (
	"net/http"
	"tianhe/config"
	"tianhe/middleware"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_NOT_FOUND_CODE = 404
	HTTP_NO_LOGIN       = 1001
	HTTP_TOKEN_INVALID  = 1002
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Conf.Server.Mode)
	r := gin.New()
	r.Use(middleware.RequestId())
	r.Use(middleware.InitApiLog())
	r.Use(middleware.NoMethodHandler())
	r.Use(gin.Recovery())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{"msg":"The request uri not found"})
	})
	r.Use(middleware.LimitHandler())

	//初始化参数校验
	if err := middleware.TransInit("zh"); err != nil {
		return nil
	}
	apiGroup := r.Group("/tianhe/v1")
	{
		registerUserRouter(apiGroup)
		registerHealthRouter(apiGroup)
		registerLoginRouter(apiGroup)
		registerHostRouter(apiGroup)
		registerSecretRouter(apiGroup)
	}

	return r
}
