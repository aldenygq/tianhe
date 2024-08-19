package app

import (
	"fmt"

	"tianhe/middleware"
	"tianhe/models"
	"tianhe/service"

	"github.com/aldenygq/toolkits"
	"github.com/gin-gonic/gin"
)

// 校验用户登录状态(uname)
func CheckUseLoginByUname(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	var param models.ParamUserEnName
	err := ctx.Validate(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	token, msg, err := service.CheckUseLoginByUname(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("check user %v login status failed:%v\n",param.EnName,err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, token)
	return
}


// 登出
func Logout(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	uname,err := GetUserByToken(ctx)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf("get user by token failed:%v\n",err)
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return 
	}
	err = middleware.DelToken(uname)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("user logout failed:%v\n", err))
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return
	}
	middleware.LogInfo(ctx.Ctx).Infof(fmt.Sprintf("user logout success"))
	ctx.Response(middleware.HTTP_SUCCESS_CODE, fmt.Sprintf("登出成功"), "")
	return
}

// 登录
func Login(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamLogin

	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid:%v",err))
		return
	}

	data,msg, err := service.Login(ctx.Ctx, param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("modify user %v password failed:%v\n", param.Mobile, err))
		ctx.Response(middleware.HTTP_CHECK_FAILED, fmt.Sprintf("重置密码失败，请联系系统管理员处理!"), data)
		return
	}

	ctx.Response(middleware.HTTP_SUCCESS_CODE, msg, data)
	return
}

// 发送短信

func SendSms(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	//参数校验
	var param models.ParamMobile
	err := ctx.ValidateJson(&param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("request param invalid"))
		return
	}

	if toolkits.CheckMobile(param.Mobile) {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("mobile %v invalid:%v\n", param.Mobile, err))
		ctx.Response(middleware.HTTP_MOBILE_INVALID, fmt.Sprintf("手机号不合法"), "")
		return
	}
	msg, err := service.SendSms(ctx.Ctx,param)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("send sms to mobile %v failed:%v\n", param.Mobile, err))
		ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
		return
	}
	ctx.Response(middleware.HTTP_FAIL_CODE, msg, "")
	return
}



func GetUserByToken(ctx middleware.Context) (string,error) {
	var header models.ParamHeader
	err := ctx.ValidateHeader(&header)
	if err != nil {
		return "",err 
	}
	ret,err := middleware.ParseToken(header.Token)
	if err != nil {
		middleware.LogErr(ctx.Ctx).Errorf(fmt.Sprintf("get user by token failed:%v\n", err))
		ctx.Response(middleware.HTTP_FAIL_CODE, fmt.Sprintf("登出失败,失败原因:%v\n",err), "") 
		return "",err 
	}
	
	return ret.UEnName,nil 
}