package app

import (
	//"encoding/json"
	"fmt"

	"tianhe/middleware"
	"tianhe/models"
	"tianhe/pkg"
	"tianhe/service"

	"github.com/gin-gonic/gin"
)
//用户设置登录有效时长,用户在登录的情况在才可以设置，第一次登录时有效期为默认值
func SetTokenExpire(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamSetUserTokenExpire
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == "" {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("token invalid"))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	middleware.LogInfo(ctx.Ctx).Infof(fmt.Sprintf("token:",accessToken))
	msg,err := service.SetTokenExpire(ctx.Ctx,accessToken,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("set token expire failed:%v\n",err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,"")
	return
}
//用户注册
func UserRegister(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	requestid,_ := c.Get("X-Request-Id")
	middleware.LogInfo(ctx.Ctx).Infof(fmt.Sprintf("requestid:%v",requestid))
	//参数校验
	var param models.ParamUserRegister
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid"))
		return
	}
	if !pkg.CheckMobile(param.Mobile) {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("mobile %v invalid:%v",param.Mobile,err))
		ctx.Response(middleware.HTTP_MOBILE_INVALID, fmt.Sprintf("手机号不合法"), "")
		return
	}
	msg,err := service.UserRegister(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("user %v register failed:%v",param.EnName,err))
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
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("token invalid"))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	middleware.LogInfo(ctx.Ctx).Infof(fmt.Sprintf("token:%v",accessToken))
	data,msg,err := service.UserInfo(ctx.Ctx,accessToken)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get user info failed:%v\n",err))
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
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("token invalid"))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	//参数校验
	var param models.ParamModifyUserPassword
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid"))
		return
	}
	
	msg,err := service.ModifyPassword(ctx.Ctx,accessToken,param.PassWord)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get user info failed:%v\n",err))
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
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid"))
		return
	}
	if !pkg.CheckMobile(param.Mobile) {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("mobile %v invalid:%v\n",param.Mobile,err))
		ctx.Response(middleware.HTTP_MOBILE_INVALID, fmt.Sprintf("手机号不合法"), "")
		return
	}
	msg,err := service.ForgotPassword(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("modify user %v password failed:%v\n",param.Mobile,err))
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
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid"))
		return
	}
	msg,err := service.ModifyUserStatus(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("modify user %v password failed:%v\n",param.EnName,err))
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
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid"))
		return
	}
	data,msg,err := service.UserList(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get user list failed:%v\n",err))
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
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid"))
		return
	}
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == "" {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("token invalid"))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	ret,err := middleware.ParseToken(accessToken)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("token invalid"))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户登录状态异常"), nil)
		return
	}
	if ret.UEnName != param.EnName {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("login user:%v,modify user:%v not match",ret.UEnName,param.EnName))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户登录状态异常"), nil)
		return
	}
	msg,err := service.ModifyUserInfo(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("modify user %v info failed:%v\n",param.EnName,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}

func Unregister(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == "" {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("token invalid"))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户未登录"), nil)
		return
	}
	msg,err := service.Unregister(ctx.Ctx,accessToken)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("%v",err))
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("用户注销失败"), nil)
		return
	}
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}