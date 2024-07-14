package app

import (
	"github.com/gin-gonic/gin"
	"tianhe/middleware"
	"tianhe/models"
	"fmt"
	"tianhe/service"
)

func RegisterCluster(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamRegisterCluster
	uname,err := GetUserByToken(ctx)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("get user by token failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return 
	}
	err = ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	param.Creator = uname
	msg, err := service.RegisterCluster(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("register k8s cluster %v failed:%v\n",param.ClusterName,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}

func CreateNs(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamCreateNs
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	msg,err := service.CreateNs(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("k8s cluster %v create ns %v failed:%v\n",param.ClusterId,param.NameSpace,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
func ClusterList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	/*
	var param models.ParamCreateNs
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	*/

	data,msg,err := service.ClusterList(ctx.Ctx)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get cliuster list failed:%v\n",err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}