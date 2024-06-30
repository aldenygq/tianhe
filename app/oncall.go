package app

import (
	"oncall/middleware"
	"oncall/service"
	//"oncall/service/test"
	//"oncall/tools/resp"
	"github.com/gin-gonic/gin"
	"oncall/models"
)
func DefaultInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamDefaultInfo
	err := ctx.Validate(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	data,msg,err := service.DefaultInfo(param)
	if err != nil {
		middleware.Logger.Error("get oncall rule list failed:",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}


func AddOncall(c *gin.Context)  {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamAddOncallRule
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	msg,err := service.AddOncall(param)
	if err != nil {
		middleware.Logger.Error("%v oncall failed:",param.,err)
		ctx.Response(HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(HTTP_SUCCESS_CODE, msg, "")
	return
}

func ModifyOncall(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamModifyOncallRule
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	msg,err := service.ModifyOncall(param)
	if err != nil {
		middleware.Logger.Error("%v oncall failed:",param.,err)
		ctx.Response(HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(HTTP_SUCCESS_CODE, msg, "")
	return
}
func OncallRules(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamSearch
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	rule,msg,err := service.OncallRules(param)
	if err != nil {
		middleware.Logger.Error("get oncall rule list failed:",err)
		ctx.Response(HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(HTTP_SUCCESS_CODE, msg, rule)
	return
}

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

