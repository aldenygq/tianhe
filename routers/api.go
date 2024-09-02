package router

import (
	"github.com/gin-gonic/gin"
	"tianhe/app"
	"tianhe/middleware"
)

func registerOncallRouter(oncall *gin.RouterGroup) {
	t := oncall.Group("/oncall").Use(middleware.Auth())
	{
		//增加值班策略默认信息
		t.GET("/DefaultRuleInfo", app.DefaultInfo)
		//增加值班策略
		t.POST("/AddOncallRule", app.AddOncall)
		//值班详情
		t.GET("/OncallInfo",app.OncallInfo)
		//获取值班规则列表
		t.GET("/OncallRules",app.OncallRules)
		//删除值班规则
		t.DELETE("/DeleteOncallRule",app.DeleteOncall)
		//修改值班规则
		t.POST("/ModifyOncallRule",app.ModifyOncallRule)
	}
}

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
	}
}

func registerSecretRouter(key *gin.RouterGroup) {
	k := key.Group("/secret").Use(middleware.Auth())
	{
		k.POST("/addCloudSecret", app.AddCloudSecret)
	}
}

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
		k.POST("/createNs",app.CreateNs) 
		//查看 pod 事件
		//k.GET("/resourceEvent",app.PodEvent)
		k.GET("/resourceEvent",app.ResourceEvent)
		//查看 podlog
		k.GET("/podLog",app.PodLog)
		//node标签
		k.GET("/nodeLable",app.NodeLable)
		//node污点
		k.GET("/nodeTaint",app.NodeTaint)
		//给 node 打标签
		k.POST("/patchNodeLable",app.PatchNodeLable)
		//给 node 设置污点
		k.POST("/patchNodeTaint",app.PatchNodeTaint)
		//设置 node 调度策略,enable(可调度)/disable(不可调度)
		k.POST("/patchNodeSchedule",app.PatchNodeSchedule)
		//设置 node 排水
		k.POST("/patchNodeDrain",app.PatchNodeDrain)
		//node下 pod 列表
		k.GET("/podsInNode",app.PodsInNode)
		//工作负载列表
		k.GET("/resourceList",app.ResourceList)
		//resource yaml
		k.GET("/resourceYaml",app.ResourceYaml)
		//resource info
		k.GET("/resourceInfo",app.ResourceInfo)
		//cluster event
		k.GET("/clusterEvent",app.ClusterEvent)
		//删除资源
		k.POST("/deleteResource",app.DeleteResource)
		//工作负载滚动重启，涉及：deployment/statefulset/daemonset 
		k.POST("/workloadRollupdate",app.WorkloadRollUpdate)
		//创建 configmap
		k.POST("/createConfigMap",app.CreateConfigMap)
		//更新 configmap
		k.POST("/updateConfigMap",app.UpdateConfigMap)
		//创建 secret
		k.POST("/createSecret",app.CreateSecret)
		//更新 secret
		k.POST("/updateSecret",app.UpdateSecret)
		//创建 resource yaml
		k.POST("/createResourceByYaml",app.CreateResourceByYaml)
		//获取k8s集群user列表
		k.GET("/userList",app.ClusterUserList)
		//获取节点池列表
		k.GET("/nodeGroup",app.NodeGroupList)
		//获取节点池下节点列表
		k.GET("/nodeListByNodeGroup",app.NodeListByNodeGroup)
		//获取集群插件列表
		k.GET("/addonList",app.AddonList)
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
