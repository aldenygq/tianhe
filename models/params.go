package models


//secret相关参数
type ParamAddCloudSecret struct {
	ParamCloudAccount
	ParamCloud
	ParamCloudProduct
	ParamEnv
	Creator string 
	AccessKey string `form:"access_key"  json:"access_key" binding:"required,min=0" label:"access_key"`
	SecretKey string `form:"secret_key"  json:"secret_key" binding:"required,min=0" label:"secret_key"`
	AccountOwner string `form:"account_owner" json:"account_owner" binding:"required,min=0" label:"account_owner"`
}
//K8S 相关参数
type ParamGetNodeGroup struct {
	ParamClusterId
}
type ParamNodeListByNodeGroup struct {
	ParamClusterId
	ParamNodeGroup
}
type ParamHostInfo struct {
	ParamHostId
}

type ParamDelHost struct {
	ParamHostId
}
type ParamCreateNs struct {
	ParamClusterId
	ParamNameSpace
}
type ParamNodeInfo struct {
	ParamClusterId
	ParamNode
}
type ParamPodInfo struct {
	ParamClusterId
	ParamNameSpace
	ParamPod
}
type ParamRegisterCluster struct {
	Kubeconfig string  `form:"kubeconfig"  json:"kubeconfig" binding:"required,min=0" label:"k8s认证文件"`
	ParamEnv
	ClusterName string `form:"cluster_name"  json:"cluster_name" binding:"required,min=0" label:"集群名称"`
	ClusterId string 	`form:"cluster_id"  json:"cluster_id" binding:"required,min=0" label:"集群id"`
	Creator string 
	ParamCloud
	ParamCloudAccount
}

type ParamCreateConfigmap struct {
	ParamClusterId
	ParamNameSpace
	ConfigMapName string `form:"configmap_name"  json:"configmap_name" binding:"required,min=0" label:"configmap名称"`
	KV  map[string]string `form:"kv"  json:"kv" binding:"required,len=0" label:"k/v"`
}
type ParamCreateResourceYaml struct {
	ParamClusterId
	ResourceType string `form:"resource_type"  json:"resource_type" binding:"required,min=0" label:"资源类型"`
	ResourceYaml string `form:"resource_yaml"  json:"resource_yaml" binding:"required,min=0" label:"资源 yaml 文件"`
}
type ParamCreateSecret struct {
	ParamClusterId
	ParamNameSpace
	SecretName string `form:"secret_name"  json:"secret_name" binding:"required,min=0" label:"secret名称"`
	Type string `form:"type"  json:"type" binding:"required,min=0" label:"secret 类型"`
	KV  map[string]string `form:"kv"  json:"kv" binding:"omitempty,len=0" label:"k/v"`
	IsEncrypt bool `form:"is_encrypt"  json:"is_encrypt" binding:"required,len=0" label:"是否加密"`
	ImageRepositoryUrl string `form:"image_repository_url"  json:"image_repository_url" binding:"omitempty,min=0" label:"镜像 url"`
	RepositoryUser string `form:"repository_user"  json:"repository_user" binding:"omitempty,min=0" label:"镜像仓库用户"`
	RepositoryPassword string `form:"repository_password"  json:"repository_password" binding:"omitempty,min=0" label:"镜像仓库用户密码"`
	Cert string `form:"cert"  json:"cert" binding:"omitempty,min=0" label:"证书"`
	Key string `form:"key"  json:"key" binding:"omitempty,min=0" label:"证书key"`
} 

type ParamReourceYaml struct {
	ParamClusterId
	NameSpace string  `form:"namespace"  json:"namespace" binding:"omitempty,min=0" label:"namespace"`
	ResourceType string `form:"resource_type"  json:"resource_type" binding:"required,min=0" label:"资源类型"`
	ResourceName string `form:"resource_name"  json:"resource_name" binding:"required,min=0" label:"资源名称"`
}
type ParamReourceList struct {
	ParamClusterId
	NameSpace string  `form:"namespace"  json:"namespace" binding:"omitempty,min=0" label:"namespace"`
	ResourceType string `form:"resource_type"  json:"resource_type" binding:"required,min=0" label:"资源类型"`
}
type ParamReourceInfo struct {
	ParamClusterId
	NameSpace string  `form:"namespace"  json:"namespace" binding:"omitempty,min=0" label:"namespace"`
	ResourceType string `form:"resource_type"  json:"resource_type" binding:"required,min=0" label:"资源类型"`
	ResourceName string `form:"resource_name"  json:"resource_name" binding:"required,min=0" label:"资源名称"`
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
type ParamNodeGroup struct {
	NodeGroupName string `form:"node_group_name"  json:"node_group_name" binding:"required,min=0" label:"节点池名称"`
}
type ParamDeployment struct {
	Deployment string `form:"deployment"  json:"deployment" binding:"required,min=0" label:"deployment"`
}
type ParamStatefulSet struct {
	StatefulSet string `form:"statefulset"  json:"statefulset" binding:"required,min=0" label:"statefulset"`
}
type ParamCloud struct {
	Cloud string `form:"cloud"  json:"cloud" binding:"required,min=0" label:"云厂商"`
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
type ParamCloudAccount struct {
	CloudAccount string  `form:"cloud_account"  json:"cloud_account" binding:"required,min=0" label:"云账号"`
}
type ParamCloudProduct struct {
	CloudProduct string  `form:"cloud_product"  json:"cloud_product" binding:"required,min=0" label:"云产品"`
}
type ParamEnv struct {
	Env string  `form:"env"  json:"env" binding:"required,min=0" label:"环境"`
}
/*
type ParamSecretType struct {
	SecretType string `form:"secret_type"  json:"secret_type" binding:"required,min=0" label:"密钥类型"`
} 
*/


type ParamHostId struct {
	HostId string `form:"host_id"  json:"host_id" binding:"required,min=0" label:"主机id"`
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
 }




//用户/登录相关
//请求header参数
type ParamHeader struct {
	Token string `header:"Access_Token" binding:"required,min=1" label:"请求header(token)"`
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

//值班相关
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
	PageNum int `form:"page_num"  json:"page_num" binding:"required,gt=0" label:"起始页"`
	PageSize int `form:"page_size"  json:"page_size" binding:"required,gt=0" label:"每页数量"`
}

type ParamOncallRule struct {
	CnTitle string `form:"cn_title"  json:"cn_title" binding:"required,min=1,max=20" label:"中文标题"`
	EnTitle string `form:"en_title" json:"en_title" binding:"required,min=1,max=20" label:"英文标题"`
	OncallCycleType string `form:"oncall_cycle_type" json:"oncall_cycle_type" binding:"required,min=1" label:"值班周期类型"`
	StartDay string  `form:"start_day" json:"start_day" binding:"required,min=1" label:"开始日期"`
	RotationNum int64 `form:"rotation_num" json:"rotation_num" binding:"required" label:"轮转次数"`
	OncallPeople [][]string`form:"oncall_people"  json:"oncall_people" binding:"required" label:"值班人员信息" description:"day类型:长度最大为7,最大值为30,custom类型必传"` 
	IsSkipWeekend int64 `form:"is_skip_weekend"  json:"is_skip_weekend" binding:"required,gt=0" label:"是否跳过周末值班"`
	SubscribeNotifyInfo []*SubscribeNotify `form:"subscribe_notify_info"  json:"subscribe_notify_info" binding:"required,dive" label:"订阅通知提醒信息"`
	SubscribeGroups  []*SubscribeGroup `form:"subscribe_groups"  json:"subscribe_groups" binding:"required,dive" label:"订阅组信息"`
	IsTemporaryOncall int64 `form:"is_temporary_oncall"  json:"is_temporary_oncall" binding:"required,gt=0" label:"是否开启临时值班"`
	TemporaryOncallInfo *TemporaryOncall `form:"temporary_oncall_info"  json:"temporary_oncall_info" binding:"omitempty,min=1,dive" label:"临时值班信息"`
	Status int64 `form:"status"  json:"status" binding:"required,gt=0" label:"是否启用"`
}
type OncallPeopleInfo struct {
	OncallPeoples []string `form:"oncall_peoples" json:"oncall_peoples" binding:"required" label:"值班人员列表"` 
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
	//SubscribeChannel string `form:"subscribe_channel" json:"subscribe_channel" binding:"required,min=1" label:"订阅渠道" description:"订阅渠道:wechat(企微)、dingtalk(钉钉)、feishu(飞书),默认dingtalk(钉钉)"`
	SubscribeChannel string `form:"subscribe_channel" json:"subscribe_channel" binding:"required,min=1" label:"订阅渠道"`
	//	SubscribeType string `form:"subscribe_type" json:"subscribe_type" binding:"required,min=1" label:"订阅方式" description:"订阅方式:group(群)、app(应用),默认group(群)"`
	SubscribeType string `form:"subscribe_type" json:"subscribe_type" binding:"required,min=1" label:"订阅方式"`
	SubscribeAddr string `form:"subscribe_addr" json:"subscribe_addr" binding:"required,min=1" label:"订阅地址"`
}
type SubscribeNotify struct{
	NotifyContent string `form:"notify_content" json:"notify_content" binding:"omitempty,min=1" label:"通知内容"`
	//IsEnabled int64 `form:"is_enabled" json:"is_enabled" binding:"omitempty,gt=0" label:"是否开启通知" description:"是否开启通知,1表示不开启，2表示开启，默认开启(1)"`
	NotifyTime string `form:"notify_time" json:"notify_time" binding:"omitempty,min=1" label:"通知时间"`
}


type ParamDutyPerson struct {
	RuleName string `form:"rule_name" json:"rule_name" binding:"omitempty,min=1" label:"规则名称" description:"若指定值班规则，则参数为值班规则名称，中英文不限"`
	User []string `form:"user" json:"user" binding:"omitempty,min=1,len=1" label:"值班人员" description:"当前值班人员"`
}


type ParamDefaultInfo struct {
	Operate string `form:"operate" json:"operate" binding:"required,min=1" label:"操作类型" description:"操作类型,可选参数有:add_oncall_rule(创建值班规则)、modify_oncall_rule(修改值班规则)、search_oncall_rule(搜索值班规则)、search_oncall_user(搜索值班人员)，中英文不限"`
	RuleId int64 `form:"rule_id" json:"rule_id" binding:"omitempty,gt=0" label:"规则id" description:"规则id"`
}
type ParamOncallInfo struct {
	RuleId int64 `form:"rule_id" json:"rule_id" binding:"required,gt=0" label:"值班规则id" description:"值班规则id"`
}