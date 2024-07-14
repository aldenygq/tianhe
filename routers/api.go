package router

import (
	"github.com/gin-gonic/gin"
	"tianhe/app"
	"tianhe/middleware"
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
	u := user.Group("/user")
	{
		//用户注册
		u.POST("/register", app.UserRegister)
		//忘记密码
		u.POST("/forgotPassword", app.ForgotPassword)

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
	a := auth.Group("/auth")
	{
		//用户登录
		a.POST("/login", app.Login)
		//查询用户登陆状态
		a.GET("/checkuserLoginByUname", app.CheckUseLoginByUname)
		//发送验证码
		a.POST("/sendsms", app.SendSms)
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
func registerHostRouter(host *gin.RouterGroup) {
	h := host.Group("/host").Use(middleware.Auth())
	{
		//添加主机
		h.POST("/addHost", app.AddHost)
		//删除主机
		h.POST("/delHost", app.DelHost)
		//获取主机信息
		h.GET("/hostInfo",app.HostInfo)
	}
}
func registerK8sRouter(k8s *gin.RouterGroup) {
	k := k8s.Group("/k8s").Use(middleware.Auth())
	{
		//添加集群
		k.POST("/register", app.RegisterCluster)
		//集群列表
		k.GET("/list", app.ClusterList)
		//创建 ns
		k.POST("createNs",app.CreateNs)
		k.POST("/register", app.RegisterCluster)
		//删除主机
		//k.POST("/deleter", app.DelHost)
		//获取主机信息
		//k.GET("/info/:id",app.HostInfo)
	}
}
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
