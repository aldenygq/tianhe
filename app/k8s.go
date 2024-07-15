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

func NsInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamCreateNs
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	data,msg,err := service.NsInfo(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get ns %v info by cluster %v failed:%v\n",param.NameSpace,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}
func NsList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamClusterId
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	data,msg,err := service.NsList(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get ns list info by cluster %v failed:%v\n",param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func PodInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamPodInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.PodInfo(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get pod %v info by cluster %v and ns %v failed:%v\n",param.ParamPod,param.ClusterId,param.NameSpace,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func PodList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamCreateNs
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.PodList(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get pod list by cluster %v and ns %v failed:%v\n",param.ClusterId,param.NameSpace,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func PodEvent(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamPodInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.PodEvent(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get pod %v event by cluster %v and ns %v failed:%v\n",param.PodName,param.ClusterId,param.NameSpace,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func PodLog(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamPodInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.PodLog(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get pod %v log by cluster %v and ns %v failed:%v\n",param.PodName,param.ClusterId,param.NameSpace,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func NodeList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamClusterId
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.NodeList(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get node list by cluster %v failed:%v\n",param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}
func NodeInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamNodeInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.NodeInfo(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get node %v info by cluster %v failed:%v",param.ParamNode,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func NodeLable(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamNodeInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.NodeLable(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get node %v label by cluster %v failed:%v",param.ParamNode,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}