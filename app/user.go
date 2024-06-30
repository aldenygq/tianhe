package app

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
	"oncall/middleware"
	"oncall/models"
	"oncall/pkg"
	"oncall/service"
)
//用户注册
func UserRegister(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamUserRegister
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	if !pkg.CheckMobile(param.Mobile) {
		middleware.Logger.Errorf("mobile %v invalid:%v\n",param.Mobile,err)
		ctx.Response(middleware.HTTP_MOBILE_INVALID, fmt.Sprintf("手机号不合法"), "")
		return
	}
	msg,err := service.UserRegister(param)
	if err != nil {
		middleware.Logger.Errorf("user %v register failed:%v\n",param.CnName,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,"")
	return
}

//获取用户信息
func UserInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == "" {
		middleware.Logger.Errorf("token invalid")
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	data,msg,err := service.UserInfo(accessToken)
	if err != nil {
		middleware.Logger.Errorf("get user info failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,data)
	return
}
//修改密码
func ModifyPassword(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == "" {
		middleware.Logger.Errorf("token invalid")
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	//参数校验
	var param models.ParamModifyUserPassword
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	
	msg,err := service.ModifyPassword(accessToken,param.PassWord)
	if err != nil {
		middleware.Logger.Errorf("get user info failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}

//忘记密码
func ForgotPassword(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var (
		param models.ParamForgotPassword
	)
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	if !pkg.CheckMobile(param.Mobile) {
		middleware.Logger.Errorf("mobile %v invalid:%v\n",param.Mobile,err)
		ctx.Response(middleware.HTTP_MOBILE_INVALID, fmt.Sprintf("手机号不合法"), "")
		return
	}
	msg,err := service.ForgotPassword(param)
	if err != nil {
		middleware.Logger.Errorf("modify user %v password failed:%v\n",user.Mobile,err)
		ctx.Response(middleware.HTTP_CHECK_FAILED, msg, "")
		return
	}
	
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}
//启用/禁用用户
func ModifyUserStatus(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var (
		param models.ParamModifyUserStatus
	)
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	msg,err := service.ModifyUserStatus(param)
	if err != nil {
		middleware.Logger.Errorf("modify user %v password failed:%v\n",user.Mobile,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}

//删除用户
func DeleteUser(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var (
		param models.ParamUserEnName
	)
	err := ctx.Validate(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	msg,err := service.DeleteUser(param)
	if err != nil {
		middleware.Logger.Errorf("delete user %v failed:%v\n",param.EnName,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}

//用户列表
func UserList(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var (
		param models.ParamUserList
	)
	err := ctx.Validate(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	data,msg,err := service.UserList(param)
	if err != nil {
		middleware.Logger.Errorf("get user list failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,data)
	return
}

func ModifyUserInfo(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var (
		param models.ParamModifyUserInfo
	)
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == "" {
		middleware.Logger.Errorf("token invalid")
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	ret,err := middleware.ParseToken(accessToken)
	if err != nil {
		middleware.Logger.Errorf("token invalid")
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户登录状态异常"), nil)
		return
	}
	if ret.UEnName != param.EnName {
		middleware.Logger.Errorf("login user:%v,modify user:%v not match",ret.UEnName,param.EnName)
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户登录状态异常"), nil)
		return
	}
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	msg,err := service.ModifyUserInfo(param)
	if err != nil {
		middleware.Logger.Errorf("modify user %v info failed:%v\n",param.EnName,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}
