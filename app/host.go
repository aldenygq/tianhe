package app

import (
	"github.com/gin-gonic/gin"
	"tianhe/middleware"
	"tianhe/models"
	"fmt"
	"tianhe/service"
)

func AddHost(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamAddHost
	uname,err := GetUserByToken(ctx)
	if err != nil {
		middleware.Log(ctx.Ctx).Errorf("get user by token failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return 
	}
	err = ctx.ValidateJson(&param)
	if err != nil {
		middleware.Log(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	param.Creator = uname
	msg, err := service.AddHost(ctx.Ctx,param)
	if err != nil {
		middleware.Log(ctx.Ctx).Errorf(fmt.Sprintf("add host %v failed:%v\n",param.HostIp,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}


func DelHost(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamDelHost
	//uname,err := GetUserByToken(ctx)
	//if err != nil {
	//	middleware.Log(ctx.Ctx).Errorf("get user by token failed:%v\n",err)
	//	ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
	//	return 
	//}
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.Log(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	//param.Creator = uname
	msg, err := service.DelHost(ctx.Ctx,param)
	if err != nil {
		middleware.Log(ctx.Ctx).Errorf(fmt.Sprintf("delete host %v failed:%v\n",param.HostId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}

func HostInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamHostInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.Log(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	//param.Creator = uname
	data,msg, err := service.HostInfo(ctx.Ctx,param)
	if err != nil {
		middleware.Log(ctx.Ctx).Errorf(fmt.Sprintf("get host %v info failed:%v\n",param.HostId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}