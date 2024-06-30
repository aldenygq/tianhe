package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	
	"github.com/gin-gonic/gin"
	"oncall/models"
	"oncall/middleware"
	"oncall/pkg"
	"oncall/config"
	"oncall/service"
	"github.com/aldenygq/toolkits"
)
//校验用户登录状态
func CheckUseLogin(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == ""  {
		middleware.Logger.Errorf("token invalid")
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("token无效"), "")
		return
	}
	result,msg,err := service.CheckUseLogin(accessToken)
	if err != nil {
		middleware.Logger.Errorf("check user %v login status failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	
	ctx.Response(middleware.HTTP_FAIL_CODE, msg, result)
	return
}
//登出
func Logout(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	accessToken := c.GetHeader(middleware.ACCESS_TOKEN)
	if accessToken == ""  {
		middleware.Logger.Errorf("token invalid")
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("token无效"), "")
		return
	}
	ret,err := middleware.ParseToken(accessToken)
	if err != nil {
		middleware.Logger.Errorf("parse token failed:%v\n",err)
		ctx.Response(middleware.HTTP_TOKEN_INVALID, fmt.Sprintf("登出失败,失败原因:%v\n",err), "")
		return
	}
	err = middleware.DelToken(accessToken)
	if err != nil {
		middleware.Logger.Errorf("user %v logout failed:%v\n",ret.UEnName,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "")
		return
	}
	middleware.Logger.Errorf("user :%v logout success",u)
	ctx.Response(middleware.HTTP_SUCCESS_CODE, fmt.Sprintf("登出成功"), "")
	return
}
//登录
func Login(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var (
		param models.ParamLogin
		user *models.Users = &models.Users{}
		//u string
	)
	
	err := ctx.ValidateJson(param)
	if err != nil {
		middleware.Logger.Error("request param invalid")
		return
	}
	
	msg,err := service.Login(c,param)
	if err != nil {
		middleware.Logger.Errorf("modify user %v password failed:%v\n",user.Mobile,err)
		ctx.Response(middleware.HTTP_CHECK_FAILED, fmt.Sprintf("重置密码失败，请联系系统管理员处理!"), "")
		return
	}
	
	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg,nil)
	return
}

//发送短信
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
		middleware.Logger.Errorf("mobile %v invalid:%v\n",param.Mobile,err)
		ctx.Response(middleware.HTTP_MOBILE_INVALID, fmt.Sprintf("手机号不合法"), "")
		return
	}
	msg,err := service.SendSms(param.Mobile)
	if err != nil {
		middleware.Logger.Errorf("send sms to mobile %v failed:%v\n",param.Mobile,err)
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
	return
}
