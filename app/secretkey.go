package app

import (
	"github.com/gin-gonic/gin"
	"tianhe/middleware"
	"tianhe/models"
	"fmt"
	"tianhe/service"
)

func AddCloudSecret(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamAddCloudSecret
	uname,err := GetUserByToken(ctx)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("get user by token failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return 
	}
	param.Creator = uname 
	err = ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.AddCloudSecret(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,"")
	return
}