package app

import (
	"github.com/gin-gonic/gin"
	"tianhe/middleware"
	"tianhe/models"
	"fmt"
	"tianhe/service"
)

func NodeGroupList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamGetNodeGroup
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.NodeGroupList(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get node group list by cluster id %v failed:%v",param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}
func NodeListByNodeGroup(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamNodeListByNodeGroup
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.NodeListByNodeGroup(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}
/*
func ServiceAccount(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamClusterId
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.ServiceAccount(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get cluster %v service account failed:%v",param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}
*/
func ClusterEvent(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamClusterId
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.ClusterEvent(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get cluster %v event failed:%v",param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func ClusterUserList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamClusterId
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.ClusterUserList(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("%v",err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func DeleteResource(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamReourceYaml
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.DeleteResource(c,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("%v",err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
func CreateConfigMap(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamCreateConfigmap
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	msg,err := service.CreateConfigMap(c,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
func UpdateConfigMap(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamCreateConfigmap
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	msg,err := service.UpdateConfigMap(c,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
func CreateSecret(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamCreateSecret
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	msg,err := service.CreateSecret(c,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
func UpdateSecret(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamCreateSecret
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.UpdateSecret(c,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
func CreateResourceByYaml(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamCreateResourceYaml
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.CreateResourceByYaml(c,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
func WorkloadRollUpdate(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamReourceInfo
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.WorkloadRollUpdate(c,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("workload %v roll restart by cluster failed:%v\n",param.ResourceName,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}


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

func ResourceEvent(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamReourceInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.ResourceEvent(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
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
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get node %v label by cluster %v failed:%v",param.NodeName,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func NodeTaint(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamNodeInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.NodeTaint(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get node %v taint by cluster %v failed:%v",param.NodeName,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func PatchNodeLable(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamPatchNodeLabel
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.PatchNodeLable(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("patch node %v label by cluster %v failed:%v",param.NodeName,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}


func PatchNodeTaint(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamPatchNodeTaint
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.PatchNodeTaint(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("patch node %v taint by cluster %v failed:%v",param.NodeName,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}

func PatchNodeSchedule(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamPatchNodeSchedule
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.PatchNodeSchedule(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("patch node %v schedule rule %v by cluster %v failed:%v",param.NodeName,param.ScheduleRule,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}

func PatchNodeDrain(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamNodeInfo
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	msg,err := service.PatchNodeDrain(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("patch node %v drain  by cluster %v failed:%v",param.NodeName,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}

func PodsInNode(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamNodeInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.PodsInNode(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get pod list by node %v and cluster %v failed:%v",param.NodeName,param.ClusterId,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}
func ResourceYaml(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamReourceYaml
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.ReourceYaml(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func ResourceList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamReourceList
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.ReourceList(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

func ResourceInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	var param models.ParamReourceYaml
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	data,msg,err := service.ResourceInfo(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}
