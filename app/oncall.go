package app

import (
	"tianhe/middleware"
	"tianhe/service"
	//"oncall/service/test"
	//"oncall/tools/resp"
	"github.com/gin-gonic/gin"
	"tianhe/models"
)

//值班默认信息
func DefaultInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamDefaultInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("request param invalid")
		return
	}
	data,msg,err := service.DefaultInfo(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("get oncall rule list failed:",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}


//新增值班规则
func AddOncall(c *gin.Context)  {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamAddOncallRule
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("request param invalid")
		return
	}
	msg,err := service.AddOncall(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("add oncall failed:",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}

//值班任务详情
func OncallInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamOncallInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("request param invalid")
		return
	}
	data,msg,err := service.OncallInfo(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("get oncall rule list failed:",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

//修改值班规则
func ModifyOncallRule(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamModifyOncallRule
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(c).Errorf("request param invalid")
		return
	}
	msg,err := service.ModifyOncallRule(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(c).Error("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}

//修改值班规则状态
func ModifyOncallRuleStatus(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamModifyOncallRuleStatus
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(c).Errorf("request param invalid")
		return
	}
	msg,err := service.ModifyOncallRuleStatus(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(c).Error("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
//值班规则列表
func OncallRules(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamSearch
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("request param invalid")
		return
	}
	rule,msg,err := service.OncallRules(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(c).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, rule)
	return
}
//删除值班规则
func DeleteOncall(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamOncallInfo
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Error("request param invalid")
		return
	}
	msg,err := service.DeleteOncall(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(c).Errorf("%v",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, "")
	return
}
/*
func CurrrentDutyInfos(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamDutyPerson
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	dutyperson,msg,err := service.CurrrentDutyInfos(param)
	if err != nil {
		middleware.Logger.Error("get current duty person failed:",err)
		ctx.Response(HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(HTTP_SUCCESS_CODE, msg, dutyperson)
	return
}
*/
