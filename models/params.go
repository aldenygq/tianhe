package models


type ParamHeader struct {
	Token string `header:"Access_Token" binding:"required,min=1" label:"请求header(token)"`
}
type ParamRegisterCluster struct {
	Kubeconfig string  `form:"kubeconfig"  json:"kubeconfig" binding:"required,min=0" label:"k8s认证文件"`
	Env string `form:"env"  json:"env" binding:"required,min=0" label:"环境"`
	ClusterName string `form:"cluster_name"  json:"cluster_name" binding:"required,min=0" label:"集群名称"`
	ClusterId string 	`form:"cluster_id"  json:"cluster_id" binding:"required,min=0" label:"集群id"`
	Creator string 
}
type ParamCreateNs struct {
	ParamClusterId
	ParamNameSpace
}
type ParamReourceYaml struct {
	ParamClusterId
	ParamNode
	NameSpace string  `form:"namespace"  json:"namespace" binding:"omitempty,min=0" label:"namespace"`
	ResourceType string `form:"resource_type"  json:"resource_type" binding:"required,min=0" label:"资源类型"`
	ResourceName string `form:"resource_name"  json:"resource_name" binding:"required,min=0" label:"资源名称"`
}
type ParamNodeInfo struct {
	ParamClusterId
	ParamNode
}
type ParamPatchNodeLabel struct {
	ParamClusterId
	ParamNode
	Labels map[string]string `form:"labels"  json:"labels" binding:"required,gt=0" label:"标签信息"`
}
type ParamPatchNodeTaint struct {
	ParamClusterId
	ParamNode
	Taints map[string]string `form:"taint"  json:"labels" binding:"required,gt=0" label:"污点信息"`
}
type ParamPatchNodeSchedule struct {
	ParamClusterId
	ParamNode
	ScheduleRule string `form:"schedule_rule"  json:"schedule_rule" binding:"required,min=0" label:"调度策略"`
}
type ParamClusterId struct {
	ClusterId string 	`form:"cluster_id"  json:"cluster_id" binding:"required,min=0" label:"集群id"`
}
type ParamNameSpace struct {
	NameSpace string  `form:"namespace"  json:"namespace" binding:"required,min=0" label:"namespace"`
}
type ParamPod struct {
	PodName string  `form:"podname"  json:"podname" binding:"required,min=0" label:"podname"`
}
type ParamNode struct {
	NodeName string  `form:"nodename"  json:"nodename" binding:"required,min=0" label:"nodename"`
}

type ParamPodInfo struct {
	ParamClusterId
	ParamNameSpace
	ParamPod
}

type ParamDelHost struct {
	ParamHostId
}
type ParamHostId struct {
	HostId string `form:"host_id"  json:"host_id" binding:"required,min=0" label:"主机id"`
}
type ParamHostInfo struct {
	ParamHostId
}

 // host参数
type ParamAddHost struct {
	HostName string `form:"host_name"  json:"host_name" binding:"required,min=0" label:"主机名"`
	HostIp string `form:"host_ip"  json:"host_ip" binding:"required,min=0" label:"主机ip"`
	// kvm/cloud/physical
	HostType   string `form:"host_type" json:"host_type" binding:"required,min=0" label:"主机类型"`
	//ssh连接端口，默认为 22
	Port       int64 `form:"port" json:"port" binding:"required,gt=0" label:"ssh连接端口"`
	AuthType   string `form:"auth_type" json:"auth_type" binding:"required,min=0" label:"登录认证类型"`
	User       string `form:"user" json:"user" binding:"required,min=0" label:"登录用户"`
	Password   string `form:"password" json:"password" binding:"required,min=0" label:"认证密码"`
	Os  string `form:"os" json:"os" binding:"required,min=0" label:"系统版本"`
	OsVersion  string `form:"os_version" json:"os_version" binding:"required,min=0" label:"系统版本"`
	PrivateKey string `form:"private_key" json:"private_key" binding:"required,min=0" label:"认证密钥"`
	Creator string 
	//Status string `form:"status" json:"status" binding:"required,min=0" label:"认证密钥"`
 }





 type ParamSetUserTokenExpire struct {
	ExpireTime int64 `form:"expire_time"  json:"expire_time" binding:"omitempty,gt=0" label:"token 有效期"`
}
type ParamLogin struct {
	Type string `form:"type"  json:"type" binding:"required,min=1" label:"登陆类型" description:"登陆类型:verify(验证码方式)/account(账密)"`
	Mobile string `form:"mobile"  json:"mobile" binding:"omitempty,min=1,max=11" label:"手机号码"`
	VerifyCode string `form:"verify_code"  json:"verify_code" binding:"omitempty,min=1,max=6" label:"验证码"`
	EnName string `form:"en_name"  json:"en_name" binding:"omitempty,min=1,max=10" label:"英文名称"`
	Email string `form:"email"  json:"email" binding:"omitempty,min=1" label:"邮箱"`
	PassWord string `form:"password"  json:"password" binding:"omitempty,min=1" label:"密码"`
}
type ParamUserRegister struct {
	Email string `form:"email"  json:"email" binding:"required,min=1" label:"邮箱"`
	EnName string `form:"en_name"  json:"en_name" binding:"required,min=1,max=10" label:"英文名称"`
	PassWord string `form:"password"  json:"password" binding:"required,min=1" label:"密码"`
	Mobile string `form:"mobile"  json:"mobile" binding:"omitempty,min=1,max=11" label:"手机号码"`
}
type ParamUserEmail struct {
	Email string `form:"email"  json:"email" binding:"required,min=1" label:"邮箱"`
}
type ParamUserStatus struct {
	Status int64  `form:"status"  json:"status" binding:"required,gt=0" label:"用户状态"`
}
type ParamModifyUserStatus struct {
	ParamUserEnName
	ParamUserStatus
}
type ParamModifyUserInfo struct {
	ParamUserEnName
	ParamUserEmail
	ParamMobile
}
type ParamForgotPassword struct {
	ParamMobile
	ParamPassword
}

type ParamUserList struct {
	EnName string `form:"en_name"  json:"en_name" binding:"omitempty,min=1,max=10" label:"英文名称"`
	Mobile string `form:"mobile"  json:"mobile" binding:"omitempty,min=1,max=11" label:"手机号码"`
	Status int64 `form:"status"  json:"status" binding:"omitempty,gt=0" label:"状态"`
}

type ParamPassword struct {
	PassWord string `form:"password"  json:"password" binding:"required,min=1" label:"密码"`
}
type ParamUserEnName struct {
	EnName string `form:"en_name"  json:"en_name" binding:"required,min=1,max=10" label:"英文名称"`
}
type ParamMobile struct {
	Mobile string `form:"mobile"  json:"mobile" binding:"required,min=1,max=11" label:"手机号码"`
}
type ParamModifyUserPassword struct {
	PassWord string `form:"password"  json:"password" binding:"required,min=1" label:"密码"`
}
type ParamAddOncallRule struct {
	ParamOncallRule
	CreatorInfo
}
type ParamModifyOncallRule struct {
	Id int64 `form:"id"  json:"id" binding:"required,gt=0" label:"规则id"`
	ParamOncallRule
	UpdatorInfo
}

type ParamSearch struct {
	//Id int64 `form:"id"  json:"id" binding:"required,gt=0" label:"规则id"`
	ParamOncallRule
	CreatorInfo
	UpdatorInfo
}

type ParamOncallRule struct {
	CnTitle string `form:"cn_title"  json:"cn_title" binding:"required,min=1,max=10" label:"中文标题"`
	EnTitle string `form:"en_title" json:"en_title" binding:"required,min=1,max=10" label:"英文标题"`
	OncallCycleType string `form:"oncall_cycle_type"  json:"oncall_cycle_type" binding:"required,min=1" label:"值班周期类型,支持:day(天)、custom(自定义)、month(月)，默认周类型，即每轮7天"`
	StartDay string  `form:"start_day" json:"start_day" binding:"required,min=1" label:"开始日期,日期不得小于当前日期" `
	RotationNum int64 `form:"rotation_num"  json:"rotation_num" binding:"required,gt=1" label:"轮转次数,如为0，则表示持续轮转"`
	PerRotationDays int64 `form:"per_rotation_days"  json:"per_rotation_days" binding:"omitempty,gt=0" label:"每轮的轮转天数，最小值为1,最大值为30,custom类型必传"`
	OncallPeople []string`form:"oncall_people"  json:"oncall_people" binding:"required,len=1" label:"值班人员信息"`
	IsSkipWeekend int64 `form:"is_skip_weekend"  json:"is_skip_weekend" binding:"required,gt=0" label:"是否跳过周末值班" description:"是否跳过周末值班，1表示不跳过，2表示跳过，默认为不跳过(1),即周末正常值班"`
	SubscribeNotifyInfo []*SubscribeNotify `form:"subscribe_notify_info"  json:"subscribe_notify_info" binding:"required,dive,len=1" label:"订阅通知提醒信息"`
	SubscribeGroups  []*SubscribeGroup `form:"subscribe_groups"  json:"subscribe_groups" binding:"required,dive,len=1" label:"订阅组信息"`
	IsTemporaryOncall int64 `form:"is_temporary_oncall"  json:"is_temporary_oncall" binding:"required,gt=0" label:"是否开启临时值班" description:"是否开启临时值班：1(不开启),2(开启)，默认是1不开启，当临时值班开启后，默认覆盖现有值班规则"`
	TemporaryOncallInfo *TemporaryOncall `form:"temporary_oncall_info"  json:"temporary_oncall_info" binding:"omitempty,min=1,dive" label:"临时值班信息"`
	Status int64 `form:"status"  json:"status" binding:"required,min=1" label:"是否启用" description:"是否启用,1表示启用，2表示不启用,3表示删除，默认启用(1)"`
}
type CreatorInfo struct{
	Creator string `form:"creator"  json:"creator" binding:"required,min=1" label:"创建人"`
}
type UpdatorInfo struct{
	Updator string `form:"updator"  json:"updator" binding:"required,min=1" label:"最后一次修改人"`
}

type TemporaryOncall struct {
	TemporaryOncallStartTime string `form:"temporary_oncall_start_time" json:"temporary_oncall_start_time" binding:"omitempty,min=1" label:"临时值班开始时间"`
	TemporaryOncallDays int64 `form:"temporary_oncall_days" json:"temporary_oncall_days" binding:"omitempty,gt=0" label:"临时值班天数"`
	IsCoverRoutineRule int64 `form:"is_cover_rotine_rule" json:"is_cover_rotine_rule" binding:"omitempty,gt=0" label:"是否覆盖普通值班规则" description:"是否覆盖普通规则,1表示覆盖,2表示不覆盖,默认覆盖(1),覆盖意为值班人员顺延"`
	TemporaryOncallPeopleInfos []string `form:"temorary_oncall_people" json:"temorary_oncall_people" binding:"required,len=1" label:"临时值班人员信息"`
}
type SubscribeGroup struct {
	SubscribeChannel string `form:"subscribe_channel" json:"subscribe_channel" binding:"required,min=1" label:"订阅渠道" description:"订阅渠道:wechat(企微)、dingtalk(钉钉)、feishu(飞书),默认dingtalk(钉钉)"`
	SubscribeType string `form:"subscribe_type" json:"subscribe_type" binding:"required,min=1" label:"订阅方式" description:"订阅方式:group(群)、app(应用),默认group(群)"`
	SubscribeAddr string `form:"subscribe_addr" json:"subscribe_addr" binding:"required,min=1" label:"订阅地址"`
}
type 	SubscribeNotify struct{
	NotifyContent string `form:"notify_content" json:"notify_content" binding:"omitempty,min=1" label:"通知内容，包含：今日值班列表、下一周期值班列表换班提醒"`
	//IsEnabled int64 `form:"is_enabled" json:"is_enabled" binding:"omitempty,gt=0" label:"是否开启通知" description:"是否开启通知,1表示不开启，2表示开启，默认开启(1)"`
	NotifyTime string `form:"notify_time" json:"notify_time" binding:"omitempty,min=1" label:"通知时间" description:"日类"`
}

/*
type OncallPeople struct {
	OncallUsers []User `form:"oncall_users" json:"oncall_users" binding:"required,len=1,dive" label:"值班人员列表"`
}
type User struct {
	 Type string `form:"type" json:"type" binding:"omitempty,min=1" label:"值班人员类型，main(主)、back(备),非必填，如不填，则表示均为主"`
	 User string `form:"user" json:"user" binding:"request,min=1" label:"值班人员信息，显示用户中文名"`
}
 */

type ParamDutyPerson struct {
	RuleName string `form:"rule_name" json:"rule_name" binding:"omitempty,min=1" label:"规则名称" description:"若指定值班规则，则参数为值班规则名称，中英文不限"`
	User []string `form:"user" json:"user" binding:"omitempty,min=1,len=1" label:"值班人员" description:"当前值班人员"`
}


type ParamDefaultInfo struct {
	Operate string `form:"operate" json:"operate" binding:"required,min=1" label:"操作类型" description:"操作类型,可选参数有:add_oncall_rule(创建值班规则)、modify_oncall_rule(修改值班规则)、search_oncall_rule(搜索值班规则)、search_oncall_user(搜索值班人员)，中英文不限"`
}