package router

import (
	"github.com/gin-gonic/gin"
	"oncall/app"
	"oncall/middleware"
)
/*
func registerOncallRouter(oncall *gin.RouterGroup) {
	t := oncall.Group("/v1")
	{
		//增加值班策略默认信息
		t.GET("/defaultinfo", app.DefaultInfo).Use(middleware.Auth)
		//增加值班策略
		t.POST("/addoncall", app.AddOncall).Use(middleware.Auth)
		//修改值班策略
		t.POST("/modifyoncall", app.ModifyOncall).Use(middleware.Auth)
		//获取值班规则列表
		t.POST("/oncallrulls",app.OncallRules).Use(middleware.Auth)
		// 获取当前值班信息
		t.POST("/currentdutyinfos",app.CurrrentDutyInfos)
	}
}
*/

//用户中心
func registerUserRouter(user *gin.RouterGroup) {
	u := user.Group("/v1")
	{
		//用户注册
		u.POST("/register", app.UserRegister)
		//忘记密码
		u.POST("/forgotPassword", app.ForgotPassword)
		//发送密码
		//u.POST("/sendsms", app.SendSms)

		u.Use(middleware.Auth())
		//用户信息(需要认证)
		u.GET("/info", app.UserInfo)
		//用户设置登录token有效期
		u.POST("/setTokenExpire", app.SetTokenExpire)
		//修改密码
		u.POST("/modifypassword", app.ModifyPassword)
		//启用/禁用/删除用户
		u.POST("/modifystatus", app.ModifyUserStatus)
		//用户列表
		u.GET("/list", app.UserList)
		//用户注销
		u.POST("/unregister", app.Unregister)
		//修改用户信息
		u.POST("/modifyuserinfo",app.ModifyUserInfo)
	}
}

//登陆认证
func registerLoginRouter(auth *gin.RouterGroup) {
	a := auth.Group("/v1")
	{
		//用户登录
		a.POST("/login", app.Login)
		//查询用户登陆状态
		a.GET("/checkuserLoginByUname", app.CheckUseLoginByUname)
		a.GET("/checkuserLoginByToken", app.CheckUseLoginByToken)
		a.Use(middleware.Auth())
		//登出
		a.POST("/logout", app.Logout)
		//token续期
		//a.GET("/renewal", app.Renewal).Use(middleware.Auth)
	}
}
/*
func registerKeyRouter(key *gin.RouterGroup) {
	k := key.Group("/v1")
	{
		//用户登录
		k.POST("/addkey", app.AddKey)
	}
}
*/
//健康检测
func registerHealthRouter(health *gin.RouterGroup) {
	health.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"errno":"200",
			"errmsg": "OK",
			"data":"",
		})
	})
}
