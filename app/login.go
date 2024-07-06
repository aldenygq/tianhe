package app

import (
	"fmt"

	"oncall/middleware"
	"oncall/models"
	"oncall/service"

	//"github.com/aldenygq/toolkits"
	"github.com/gin-gonic/gin"
)

// 校验用户登录状态(uname)
func CheckUseLoginByUname(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamUserEnName
	err := ctx.Validate(&param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}

	token, msg, err := service.CheckUseLoginByUname(param)
	if err != nil {
		middleware.Logger.Errorf("check user %v login status failed:%v\n",param.EnName,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, token)
	return
}

// 校验用户登录状态(token)
func CheckUseLoginByUname(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamUserEnName
	err := ctx.Validate(&param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}

	token, msg, err := service.CheckUseLoginByUname(param)
	if err != nil {
		middleware.Logger.Errorf("check user %v login status failed:%v\n",param.EnName,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, token)
	return
}

// 登出
func Logout(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == "" {
		middleware.Logger.Errorf("token invalid")
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("token无效"), "")
		return
	}
	ret,err:= middleware.ParseToken(accessToken)
	if err != nil {
		middleware.Logger.Errorf("user logout failed:%v\n", err)
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return
	}
	err = middleware.DelToken(ret.UEnName)
	if err != nil {
		middleware.Logger.Errorf("user logout failed:%v\n", err)
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return
	}
	middleware.Logger.Errorf("user logout success")
	ctx.Response(middleware.HTTP_SUCCESS_CODE, fmt.Sprintf("登出成功"), "")
	return
}

// 登录
func Login(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var (
		param models.ParamLogin
	)

	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.Logger.Error("request param invalid:%v",err)
		return
	}

	data,msg, err := service.Login(c, param)
	if err != nil {
		middleware.Logger.Errorf("modify user %v password failed:%v\n", param.Mobile, err)
		ctx.Response(middleware.HTTP_CHECK_FAILED, fmt.Sprintf("重置密码失败，请联系系统管理员处理!"), data)
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

// 发送短信
/*
func SendSms(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamMobile
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}

	if toolkits.CheckMobile(param.Mobile) {
		middleware.Logger.Errorf("mobile %v invalid:%v\n", param.Mobile, err)
		ctx.Response(middleware.HTTP_MOBILE_INVALID, fmt.Sprintf("手机号不合法"), "")
		return
	}
	msg, err := service.SendSms(param.Mobile)
	if err != nil {
		middleware.Logger.Errorf("send sms to mobile %v failed:%v\n", param.Mobile, err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
	return
}
*/
