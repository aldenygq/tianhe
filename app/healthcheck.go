package app

import (
	"github.com/gin-gonic/gin"
	"tianhe/middleware"
)

func HealthCheck(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	ctx.Response(200, "", "")
	return
}