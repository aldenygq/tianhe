package router

import (
	"github.com/gin-gonic/gin"
	"oncall/app"
	"oncall/middleware"
)

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

//用户中心
func registerUserRouter(user *gin.RouterGroup) {
	u := user.Group("/v1")
	{
		//用户注册
		u.POST("/register", app.UserRegister)
		//用户信息
		u.GET("/info", app.UserInfo).Use(middleware.Auth)
		//修改密码
		u.POST("/modifypassword", app.ModifyPassword).Use(middleware.Auth)
		//忘记密码
		u.POST("/forgotpassword", app.ForgotPassword)
		//启用/禁用用户
		u.POST("/modifystatus", app.ModifyUserStatus).Use(middleware.Auth)
		//用户列表
		u.GET("/list", app.UserList).Use(middleware.Auth)
		//删除用户
		u.DELETE("/delete", app.DeleteUser).Use(middleware.Auth)
		//修改用户信息
		u.POST("/modifyuserinfo",app.ModifyUserInfo).Use(middleware.Auth)
	}
}

//登陆认证
func registerLoginRouter(auth *gin.RouterGroup) {
	a := auth.Group("/v1")
	{
		//用户登录
		a.POST("/login", app.Login)
		//发送密码
		a.POST("/sendsms", app.SendSms)
		//登出
		a.POST("/logout", app.Logout).Use(middleware.Auth)
		//查询用户登陆状态
		a.GET("/checkuserLogin", app.CheckUseLogin)
		//token续期
		//a.GET("/renewal", app.Renewal).Use(middleware.Auth)
	}
}

func registerKeyRouter(key *gin.RouterGroup) {
	k := key.Group("/v1")
	{
		//用户登录
		k.POST("/addkey", app.AddKey)
	}
}
//健康检测
func registerHealthRouter(health *gin.RouterGroup) {
	health.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
}
