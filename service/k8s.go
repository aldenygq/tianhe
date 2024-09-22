package service

import (
	"encoding/base64"
	"fmt"
	"tianhe/middleware"
	"tianhe/models"
	"tianhe/pkg"
	"time"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	batchV1 "k8s.io/api/batch/v1"
	networkV1 "k8s.io/api/networking/v1"
	storageV1 "k8s.io/api/storage/v1"
	rbacV1 "k8s.io/api/rbac/v1"
)

func AddonList(c *gin.Context,param models.ParamClusterId) (interface{},string,error) {
	var (
		addons interface{}
		err error 
		cluster *models.K8sCluster = &models.K8sCluster{}
		secretinfo *models.CloudSecretInfo = &models.CloudSecretInfo{}
	)
	cluster.ClusterId = param.ClusterId
	err = cluster.GetClusterById()
	if err != nil {
		middleware.Log(c).Errorf("get cluster %v failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("et cluster %v failed:%v\n",param.ClusterId,err),err 
	}
	secretinfo.CloudAccount = cluster.CloudAccount
	secretinfo.Cloud = cluster.Cloud
	secretinfo.Env = cluster.Env
	secretinfo.CloudProduct = cluster.CloudProduct
	err = secretinfo.GetSecretInfoByType()
	if err != nil {
		middleware.Log(c).Errorf("cloud %v ,cloud product:%v,cloud account %v,get secret info failed:%v\n",cluster.Cloud,cluster.CloudProduct,cluster.CloudAccount,err)
		return nil,fmt.Sprintf("cloud %v ,cloud product:%v,cloud account %v,get secret info failed:%v\n",cluster.Cloud,cluster.CloudProduct,cluster.CloudAccount,err),err 
	}
	// 解码Base64字符串
	decodedBytes, err := base64.StdEncoding.DecodeString(secretinfo.SecreyKey)
	if err != nil {
		middleware.Log(c).Errorf("access key:%v,decode base64 secret key info failed:%v\n",secretinfo.AccessKey,err)
		return nil,fmt.Sprintf("decode base64 secret key info by failed:%v\n",err),err 
	}
	switch cluster.Cloud {
	case "idc":
		// id此处不处理
		//nodegroups = nil 
	case "aliyun":
		addons ,err = pkg.AddonList(param.ClusterId,secretinfo.AccessKey,string(decodedBytes))
	//case "aws":
	//	nodegroups ,err = pkg.NodeGroupListByAws(param.ClusterId)
	//case "huaweicloud":
	//	nodegroups ,err = pkg.NodeGroupListByHuaweiCloud(param.ClusterId)
	//case "qcloud":
	//	nodegroups ,err = pkg.NodeGroupListByQcloud(param.ClusterId)
	default:
		middleware.Log(c).Errorf("cloud invalid")
		return nil,fmt.Sprintf("cloud invalid"),errors.New(fmt.Sprintf("cloud invalid"))
	}
	if err != nil {
		middleware.Log(c).Errorf("get addon list by cluster:%v failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("get addon list by cluster:%v failed:%v\n",param.ClusterId,err),err
	}
	return addons,fmt.Sprintf("get addon list by cluster:%v success",param.ClusterId),nil 
}
func NodeGroupList(c *gin.Context,param models.ParamGetNodeGroup) (interface{},string,error) {
	var (
		nodegroups interface{}
		err error 
		cluster *models.K8sCluster = &models.K8sCluster{}
		secretinfo *models.CloudSecretInfo = &models.CloudSecretInfo{}
	)
	cluster.ClusterId = param.ClusterId
	err = cluster.GetClusterById()
	if err != nil {
		middleware.Log(c).Errorf("get cluster %v failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("et cluster %v failed:%v\n",param.ClusterId,err),err 
	}
	secretinfo.CloudAccount = cluster.CloudAccount
	secretinfo.Cloud = cluster.Cloud
	secretinfo.Env = cluster.Env
	secretinfo.CloudProduct = cluster.CloudProduct
	err = secretinfo.GetSecretInfoByType()
	if err != nil {
		middleware.Log(c).Errorf("cloud %v ,cloud product:%v,cloud account %v,get secret info failed:%v\n",cluster.Cloud,cluster.CloudProduct,cluster.CloudAccount,err)
		return nil,fmt.Sprintf("cloud %v ,cloud product:%v,cloud account %v,get secret info failed:%v\n",cluster.Cloud,cluster.CloudProduct,cluster.CloudAccount,err),err 
	}
	// 解码Base64字符串
	decodedBytes, err := base64.StdEncoding.DecodeString(secretinfo.SecreyKey)
	if err != nil {
		middleware.Log(c).Errorf("access key:%v,decode base64 secret key info failed:%v\n",secretinfo.AccessKey,err)
		return nil,fmt.Sprintf("decode base64 secret key info by failed:%v\n",err),err 
	}
	switch cluster.Cloud {
	case "idc":
		// idc无节点池概念
		nodegroups = nil 
	case "aliyun":
		nodegroups ,err = pkg.NodeGroupListByAliyun(param.ClusterId,secretinfo.AccessKey,string(decodedBytes))
	//case "aws":
	//	nodegroups ,err = pkg.NodeGroupListByAws(param.ClusterId)
	//case "huaweicloud":
	//	nodegroups ,err = pkg.NodeGroupListByHuaweiCloud(param.ClusterId)
	//case "qcloud":
	//	nodegroups ,err = pkg.NodeGroupListByQcloud(param.ClusterId)
	default:
		middleware.Log(c).Errorf("cloud invalid")
		return nil,fmt.Sprintf("cloud invalid"),errors.New(fmt.Sprintf("cloud invalid"))
	}
	if err != nil {
		middleware.Log(c).Errorf("get node group list by cluster:%v failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("get node group list by cluster:%v failed:%v\n",param.ClusterId,err),err
	}
	return nodegroups,fmt.Sprintf("get node group list by cluster:%v success",param.ClusterId),nil 
}
func NodeListByNodeGroup(c *gin.Context,param models.ParamNodeListByNodeGroup) (interface{},string,error) {
	var (
		nodelist interface{}
		err error 
		secretinfo *models.CloudSecretInfo = &models.CloudSecretInfo{}
		cluster *models.K8sCluster = &models.K8sCluster{}
	)
	cluster.ClusterId = param.ClusterId
	err = cluster.GetClusterById()
	if err != nil {
		middleware.Log(c).Errorf("get cluster %v failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("et cluster %v failed:%v\n",param.ClusterId,err),err 
	}
	secretinfo.CloudAccount = cluster.CloudAccount
	secretinfo.Cloud = cluster.Cloud
	secretinfo.Env = cluster.Env
	secretinfo.CloudProduct = cluster.CloudProduct
	err = secretinfo.GetSecretInfoByType()
	if err != nil {
		middleware.Log(c).Errorf("cloud %v ,cloud product:%v,cloud account %v,get secret info failed:%v\n",cluster.Cloud,cluster.CloudProduct,cluster.CloudAccount,err)
		return nil,fmt.Sprintf("cloud %v ,cloud product:%v,cloud account %v,get secret info failed:%v\n",cluster.Cloud,cluster.CloudProduct,cluster.CloudAccount,err),err 
	}
	// 解码Base64字符串
	decodedBytes, err := base64.StdEncoding.DecodeString(secretinfo.SecreyKey)
	if err != nil {
		middleware.Log(c).Errorf("access key:%v,decode base64 secret key info failed:%v\n",secretinfo.AccessKey,err)
		return nil,fmt.Sprintf("decode base64 secret key info by failed:%v\n",err),err 
	}
	//获取节点列表
	switch cluster.Cloud {
	case "idc":
		//idc无节点池概念,直接获取集群全部节点
		client,err := GetK8sClientByClusterId(c,param.ClusterId)
		if err != nil {
			middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
			return "",fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
		}
		nodelist,err = client.NodeList() 
	case "aliyun":
		nodelist,err = pkg.NodeListByNodeGroup(param.ClusterId,param.NodeGroupName,secretinfo.AccessKey,string(decodedBytes))
	//case "aws":
	//	nodegroups ,err = pkg.NodeGroupListByAws(param.ClusterId)
	//case "huaweicloud":
	//	nodegroups ,err = pkg.NodeGroupListByHuaweiCloud(param.ClusterId)
	//case "qcloud":
	//	nodegroups ,err = pkg.NodeGroupListByQcloud(param.ClusterId)
	default:
		middleware.Log(c).Errorf("cloud invalid")
		return nil,fmt.Sprintf("cloud invalid"),errors.New(fmt.Sprintf("cloud invalid"))
	}
	if err != nil {
		middleware.Log(c).Errorf("get node list by cluster:%v and node group %v failed:%v\n",param.ClusterId,param.NodeGroupName,err)
		return nil,fmt.Sprintf("get node list by cluster:%v and node group %v failed:%v\n",param.ClusterId,param.NodeGroupName,err),err
	}
	return nodelist,fmt.Sprintf("get node list by cluster:%v and node group %v success",param.ClusterId,param.NodeGroupName),nil 
}

func GetKubeconfigExpire(c *gin.Context, param models.ParamClusterId) (int64,string,error) {
	var cluster *models.K8sCluster = &models.K8sCluster{}
	var timestamp int64 
	cluster.ClusterId = param.ClusterId
	err := cluster.GetClusterById()
	if err != nil {
		middleware.Log(c).Errorf("get cluster info by id %v failed:%v\n",param.ClusterId,err)
		return timestamp,fmt.Sprintf("get cluster info by id failed:%v\n",err),err 
	}
	timestamp,err = GetKubeconfigExpireTime(c,cluster.Kubeconfig)
	if err != nil {
		middleware.Log(c).Errorf("%v",err)
		return timestamp,fmt.Sprintf("get kubecinfig expire time  failed:%v",err),err 
	}
	return timestamp,fmt.Sprintf("get kubecinfig expire time  success"),nil 
}
func GetKubeconfigExpireTime(c *gin.Context,kubeconfig string) (int64,error) {
	var timestamp int64 
	// 解码Base64字符串
	decodedBytes, err := base64.StdEncoding.DecodeString(kubeconfig)
	if err != nil {
		middleware.Log(c).Errorf("decode base64 kubeconfig info by failed:%v\n",err)
		return timestamp,err 
	}
	cmd := fmt.Sprintf("echo %v | grep client-certificate-data | awk -F ' ' '{print $2}' |base64 -d| openssl x509 -text -noout -dates | grep After |awk -F '=' '{print $2}' | grep -v '^$'",string(decodedBytes))
	content,err := pkg.RunCmd(cmd)
	if err != nil {
		middleware.Log(c).Errorf("run command failed:%v\n",err)
		return timestamp,err 
	}
	t, err := time.Parse("Jan 02 15:04:05 2006 GMT", strings.TrimSuffix(content, "\x0a"))
	if err != nil {
		middleware.Log(c).Errorf("parse time failed:%v\n",err)
		return timestamp,err 
	}
	timestamp = t.Unix()
	//tm := time.Unix(timestamp,0)
	//fmt.Printf("time:%v\n",tm)
	//fmt.Println(tm.Format("2006-01-02 15:04:05"))
	return timestamp,nil 
}
func ClusterUserList(c *gin.Context, param models.ParamClusterId) (interface{},string,error) {
	var cluster *models.K8sCluster = &models.K8sCluster{}
	cluster.ClusterId = param.ClusterId
	users,err := cluster.ClusterUsers()
	if err != nil {
		middleware.Log(c).Errorf("get user list by cluster %v failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("get user list by cluster %v failed:%v\n",param.ClusterId,err),err 		
	}

	return users,fmt.Sprintf("get user list by cluster %v success",param.ClusterId),nil 
}
func RegisterCluster(c *gin.Context, param models.ParamRegisterCluster) (string,error) {
	var cluster *models.K8sCluster = &models.K8sCluster{}
	username,err := GetUserByKubeconfig(c,param.Kubeconfig)
	if err != nil {
		middleware.Log(c).Errorf("get cluster %v user info by kubeconfig failed:%v\n",param.ClusterName,err)
		return fmt.Sprintf("get cluster %v user info by kubeconfig failed:%v\n",param.ClusterName,err),err 
	}
	timestamp,err := GetKubeconfigExpireTime(c,param.Kubeconfig)
	if err != nil {
		middleware.Log(c).Errorf("get cluster %v kubeconfig failed:%v\n",err)
		return fmt.Sprintf("get cluster %v kubeconfig failed:%v\n",param.ClusterName,err),err 
	}
	if timestamp <= time.Now().Unix() {
		middleware.Log(c).Errorf("cluster %v kubeconfig expired:%v",param.ClusterName)
		return fmt.Sprintf("cluster %v kubeconfig expired:%v",param.ClusterName),errors.New(fmt.Sprintf("cluster %v kubeconfig expired:%v",param.ClusterName)) 
	}
	cluster.ClusterId = param.ClusterId
	cluster.ClusterName = param.ClusterName
	cluster.Creator = param.Creator
	cluster.Ctime = time.Now().Unix()
	cluster.Env = param.Env
	cluster.Kubeconfig = base64.StdEncoding.EncodeToString([]byte(param.Kubeconfig))
	cluster.ClusterUser = username
	cluster.Status= 1
	cluster.CloudAccount = param.CloudAccount
	cluster.Cloud = param.Cloud
	switch cluster.Cloud {
	case "aliyun":
		cluster.CloudProduct = "ack"
	case "aws":
		cluster.CloudProduct = "eks"
	case "huaweicloud":
		cluster.CloudProduct = "cce"
	case "qcloud":
		cluster.CloudProduct = "tke"
	case "gcp":
		cluster.CloudProduct = "gke"
	case "idc":
		cluster.CloudProduct = "idc"
	default:
		middleware.Log(c).Errorf("param cloud %v invalid",param.Cloud)
		return fmt.Sprintf("param cloud %v invalid",param.Cloud),errors.New(fmt.Sprintf("param cloud %v invalid",param.Cloud)) 
	}
	err = cluster.Create()
	if err != nil {
		middleware.Log(c).Errorf("register k8s cluster %v failed:%v\n",param.ClusterName,err)
		return fmt.Sprintf("register k8s cluster %v failed:%v\n",param.ClusterName,err),err 
	}

	middleware.Log(c).Infof("register k8s cluster %v success",param.ClusterName)
	return fmt.Sprintf("register k8s cluster %v success",param.ClusterName),nil 
}

func CreateNs(c *gin.Context, param models.ParamCreateNs) (string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	err = client.CreateNs(param.NameSpace)
	if err != nil {
		middleware.Log(c).Errorf("k8s cluster %v create namespace failed:%v\n",param.ClusterId,param.NameSpace,err)
		return fmt.Sprintf("k8s cluster %v create namespace failed:%v\n",param.ClusterId,param.NameSpace,err),err 
	}
	return fmt.Sprintf("create ns %v sucsess",param.NameSpace),nil 
}

func ClusterList(c *gin.Context) ([]*models.K8sCluster,string,error) {
	var cluster *models.K8sCluster = &models.K8sCluster{}
	list,err := cluster.List()
	if err != nil {
		middleware.Log(c).Errorf("get k8s cluster list failed:%v\n",err)
		return nil,fmt.Sprintf("get k8s cluster list failed:%v\n",err),err 
	}
	return list,fmt.Sprintf("get k8s cluster list success"),nil
} 

func ResourceEvent(c *gin.Context,param models.ParamReourceInfo) (interface{},string,error) {
	var (
		filter string 
		event interface{}
		err error 
	)
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	switch param.ResourceType{
	case "pod":
		filter = fmt.Sprintf("involvedObject.name=%s", param.ResourceName)
	case "node":
		filter = fmt.Sprintf("involvedObject.kind=Node,involvedObject.name=%s", param.ResourceName)
	default:
		middleware.Log(c).Errorf("resource type %v invalid",param.ResourceType)
		return nil,fmt.Sprintf("resource type %v invalid",param.ResourceType),errors.New(fmt.Sprintf("resource type %v invalid",param.ResourceType))
	}
	event,err = client.Event(param.NameSpace,filter)
    if err != nil {
		middleware.Log(c).Errorf("get resource type %v resource %v event by cluster %v and ns %v failed:%v\n",param.ResourceType,param.ResourceName,param.ParamClusterId,param.NameSpace,err)
		return nil,fmt.Sprintf("get resource type %v resource %v event by cluster %v and ns %v failed:%v\n",param.ResourceType,param.ResourceName,param.ParamClusterId,param.NameSpace,err),err 
    }

	return event,fmt.Sprintf("get resource type %v resource %v event  by cluster %v and ns %v success",param.ResourceType,param.ResourceName,param.ParamClusterId,param.NameSpace),nil 
}

func PodLog(c *gin.Context,param models.ParamPodInfo) (interface{},string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	log,err := client.Log(param.NameSpace,param.PodName)
	if err != nil {
		middleware.Log(c).Errorf("get pod %v log by cluster %v and ns %v failed:%v\n",param.ParamPod,param.ClusterId,param.NameSpace,err)
		return nil,fmt.Sprintf("get pod %v log by cluster %v and ns %v failed:%v\n",param.ParamPod,param.ClusterId,param.NameSpace,err),err 
    }
	return log,fmt.Sprintf("get pod %v log by cluster %v and ns %v success",param.ParamPod,param.ClusterId,param.NameSpace),nil 
}

func NodeLable(c *gin.Context,param models.ParamNodeInfo) (interface{},string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	label,err := client.NodeLable(param.NodeName)
	if err != nil {
		middleware.Log(c).Errorf("get node %v label by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err)
		return nil,fmt.Sprintf("get node %v label by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err),err 		
	}

	return label,fmt.Sprintf("get node %v label by cluster %v success",param.NodeName,param.ClusterId),nil 
}

func NodeTaint(c *gin.Context,param models.ParamNodeInfo) (interface{},string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	taint,err := client.NodeTaint(param.NodeName)
	if err != nil {
		middleware.Log(c).Errorf("get node %v taint by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err)
		return nil,fmt.Sprintf("get node %v taint by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err),err 		
	}

	return taint,fmt.Sprintf("get node %v taint by cluster %v success",param.NodeName,param.ClusterId),nil 

}

func PatchNodeLable(c *gin.Context,param models.ParamPatchNodeLabel) (string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	err = client.PatchNodeLable(param.NodeName,param.Labels)
	if err != nil {
		middleware.Log(c).Errorf("patch node %v label by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err)
		return fmt.Sprintf("patch node %v label by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err),err 		
	}
	return fmt.Sprintf("patch node %v label by cluster %v success",param.NodeName,param.ClusterId),nil 
}

func PatchNodeTaint(c *gin.Context,param models.ParamPatchNodeTaint) (string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	err = client.PatchNodeTaint(param.NodeName,param.Taints)
	if err != nil {
		middleware.Log(c).Errorf("patch node %v taint by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err)
		return fmt.Sprintf("patch node %v taint by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err),err 		
	}
	return fmt.Sprintf("patch node %v taint by cluster %v success",param.NodeName,param.ClusterId),nil
}

func PatchNodeSchedule(c *gin.Context,param models.ParamPatchNodeSchedule) (string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	err = client.PatchNodeSchedule(param.NodeName,param.ScheduleRule)
	if err != nil {
		middleware.Log(c).Errorf("patch node %v schedule rule %v by cluster %v failed:%v\n",param.NodeName,param.ScheduleRule,param.ClusterId,err)
		return fmt.Sprintf("patch node %v schedule rule %v  by cluster %v failed:%v\n",param.NodeName,param.ScheduleRule,param.ClusterId,err),err 		
	}
	return fmt.Sprintf("patch node %v schedule rule %v  by cluster %v success",param.NodeName,param.ScheduleRule,param.ClusterId),nil
}

func PatchNodeDrain(c *gin.Context,param models.ParamNodeInfo) (string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	err = client.PatchNodeDrain(param.NodeName)
	if err != nil {
		middleware.Log(c).Errorf("patch node %v drain  by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err)
		return fmt.Sprintf("patch node %v drain by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err),err 		
	}
	return fmt.Sprintf("patch node %v drain  by cluster %v success",param.NodeName,param.ClusterId),nil
}

func PodsInNode(c *gin.Context,param models.ParamNodeInfo) (interface{},string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return nil,fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	pods,err := client.PodsInNode(param.NodeName)
	if err != nil {
		middleware.Log(c).Errorf("get pod list by node %v and cluster %v failed:%v\n",param.NodeName,param.ClusterId,err)
		return nil,fmt.Sprintf("get pod list by node %v and cluster %v failed:%v\n",param.NodeName,param.ClusterId,err),err 		
	}
	return pods,fmt.Sprintf("get pod list by node %v and cluster %v success",param.NodeName,param.ClusterId),nil
}

func ReourceYaml(c *gin.Context,param models.ParamReourceYaml) (string,string,error) {
	var (
		err error
		resource interface{}
	) 
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return "",fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	switch param.ResourceType {
	case "node":
		resource,err = client.NodeInfo(param.ResourceName)
	case "pod":
		resource,err = client.PodInfo(param.NameSpace,param.ResourceName)
	case "deployment":
		resource,err = client.DeploymentInfo(param.NameSpace,param.ResourceName) 
	case "svc":
		resource,err = client.SvcInfo(param.NameSpace,param.ResourceName) 
	case "statefulset":
		resource,err = client.StatefulSetInfo(param.NameSpace,param.ResourceName) 
	case "daemonset":
		resource,err = client.DaemonSetInfo(param.NameSpace,param.ResourceName) 
	case "job":
		resource,err = client.JobInfo(param.NameSpace,param.ResourceName) 
	case "crobjob":
		resource,err = client.CronJobInfo(param.NameSpace,param.ResourceName) 
	case "namespace":
		resource,err = client.NsInfo(param.ResourceName) 
	case "ingress":
		resource,err = client.IngressInfo(param.NameSpace,param.ResourceName)
	case "configmap":
		resource,err = client.ConfigMapInfo(param.NameSpace,param.ResourceName)
	case "secret":
		resource,err = client.SecretInfo(param.NameSpace,param.ResourceName)
	case "pvc":
		resource,err = client.PvcInfo(param.NameSpace,param.ResourceName)
	case "pv":
		resource,err = client.PvInfo(param.ResourceName)
	case "storageclass":
		resource,err = client.StorageClassInfo(param.NameSpace,param.ResourceName)
	case "serviceaccount":
		resource,err = client.ServiceAccountInfo(param.NameSpace,param.ResourceName)
	case "role":
		resource,err = client.RoleInfo(param.NameSpace,param.ResourceName)
	case "rolebinding":
		resource,err = client.RoleBindingInfo(param.NameSpace,param.ResourceName)
	default:
		middleware.Log(c).Errorf("resource type:%v invalid",param.ResourceType)
		return "",fmt.Sprintf("resource type:%v invalid",param.ResourceType),errors.New(fmt.Sprintf("resource type:%v invalid",param.ResourceType))
	}
	if err != nil {
		middleware.Log(c).Errorf("get type %v resource %v info by ns %v and cluster %v failed:%v\n",param.ResourceType,param.ResourceName,param.NameSpace,param.ParamClusterId,err)
		return "",fmt.Sprintf("get type %v resource %v info by ns %v and cluster %v failed:%v\n",param.ResourceType,param.ResourceName,param.NameSpace,param.ParamClusterId,err),err 
	}
	/*
	node,err := client.NodeInfo(param.NodeName)
	if err != nil {
		middleware.Log(c).Errorf("get  %v %v by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err)
		return "",fmt.Sprintf("get  node %v by cluster %v failed:%v\n",param.NodeName,param.ClusterId,err),err 		
	}
	*/
	out,err := pkg.ToYAML(resource)
	if err != nil {
		middleware.Log(c).Errorf("resource type %v,resource name:%v,to yaml failed:%v\n",param.ResourceType,param.ResourceName,err)
		return "",fmt.Sprintf("resource type %v,resource name:%v,to yaml failed:%v\n",param.ResourceType,param.ResourceName,err),err 
	}
	return out,fmt.Sprintf("resource type %v,resource name:%v,to yaml to yaml success",param.ResourceType,param.ResourceName),nil
}

func ReourceList(c *gin.Context,param models.ParamReourceList) (interface{},string,error) {
	var (
		err error
		resources interface{}
	) 
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return "",fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	switch param.ResourceType {
	case "namespace":
		resources,err = client.NsList() 
	case "deployment":
		resources,err = client.DeploymentList(param.NameSpace) 
	case "node":
		resources,err = client.NodeList() 
	case "pod":
		resources,err = client.PodList(param.NameSpace) 
	case "svc":
		resources,err = client.SvcList(param.NameSpace) 
	case "statefulset":
		resources,err = client.StatefulSetList(param.NameSpace) 
	case "daemonset":
		resources,err = client.DaemonSetList(param.NameSpace) 
	case "job":
		resources,err = client.JobList(param.NameSpace) 
	case "cronjob":
		resources,err = client.CronJobList(param.NameSpace) 
	case "ingress":
		resources,err = client.IngressList(param.NameSpace) 
	case "configmap":
		resources,err = client.ConfigMapList(param.NameSpace) 
	case "secret":
		resources,err = client.SecretList(param.NameSpace) 
	case "pvc":
		resources,err = client.PvcList(param.NameSpace) 
	case "pv":
		resources,err = client.PvList() 
	case "storageclass":
		resources,err = client.StorageClassList() 
	case "serviceaccount":
		resources,err = client.ServiceAccountList() 
	case "role":
		resources,err = client.RoleList(param.NameSpace) 
	case "rolebinding":
		resources,err = client.RoleBindingList(param.NameSpace) 
	default:
		middleware.Log(c).Errorf("search resource type %v invalid",param.ResourceType)
		return nil,fmt.Sprintf("search resource type %v invalid",param.ResourceType),errors.New(fmt.Sprintf("search resource type %v invalid",param.ResourceType))
	}

	if err != nil {
		middleware.Log(c).Errorf("get resource type %v list by cluster %v failed:%v\n",param.ResourceType,param.ParamClusterId,err)
		return nil,fmt.Sprintf("get resource type %v list by cluster %v failed:%v\n",param.ResourceType,param.ParamClusterId),err 
	}
	
	return resources,fmt.Sprintf("get resource type %v list by cluster %v success",param.ResourceType,param.ParamClusterId),nil 
}

func ResourceInfo(c *gin.Context,param models.ParamReourceYaml) (interface{},string,error) {
	var (
		err error
		resources interface{}
	) 
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return "",fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	switch param.ResourceType {
	case "namespace":
		resources,err = client.NsInfo(param.ResourceName) 
	case "deployment":
		resources,err = client.DeploymentInfo(param.NameSpace,param.ResourceName) 
	case "node":
		resources,err = client.NodeInfo(param.ResourceName) 
	case "pod":
		resources,err = client.PodInfo(param.NameSpace,param.ResourceName) 
	case "svc":
		resources,err = client.SvcInfo(param.NameSpace,param.ResourceName) 
	case "statefulset":
		resources,err = client.StatefulSetInfo(param.NameSpace,param.ResourceName) 
	case "daemonset":
		resources,err = client.DaemonSetInfo(param.NameSpace,param.ResourceName) 
	case "job":
		resources,err = client.JobInfo(param.NameSpace,param.ResourceName) 
	case "cronjob":
		resources,err = client.CronJobInfo(param.NameSpace,param.ResourceName) 
	case "ingress":
		resources,err = client.IngressInfo(param.NameSpace,param.ResourceName) 
	case "configmap":
		resources,err = client.ConfigMapInfo(param.NameSpace,param.ResourceName) 
	case "secret":
		resources,err = client.SecretInfo(param.NameSpace,param.ResourceName) 
	case "pvc":
		resources,err = client.PvcInfo(param.NameSpace,param.ResourceName) 
	case "pv":
		resources,err = client.PvInfo(param.ResourceName) 
	case "storageclass":
		resources,err = client.StorageClassInfo(param.NameSpace,param.ResourceName) 
	case "serviceaccount":
		resources,err = client.ServiceAccountInfo(param.NameSpace,param.ResourceName)
	case "role":
		resources,err = client.RoleInfo(param.NameSpace,param.ResourceName)
	case "rolebinding":
		resources,err = client.RoleBindingInfo(param.NameSpace,param.ResourceName)
	default:
		middleware.Log(c).Errorf("search resource type %v invalid",param.ResourceType)
		return nil,fmt.Sprintf("search resource type %v invalid",param.ResourceType),errors.New(fmt.Sprintf("search resource type %v invalid",param.ResourceType))
	}

	if err != nil {
		middleware.Log(c).Errorf("get resource type %v list by cluster %v failed:%v\n",param.ResourceType,param.ParamClusterId,err)
		return nil,fmt.Sprintf("get resource type %v list by cluster %v failed:%v\n",param.ResourceType,param.ParamClusterId),err 
	}
	
	return resources,fmt.Sprintf("get resource type %v list by cluster %v success",param.ResourceType,param.ParamClusterId),nil 

}
func ClusterEvent(c *gin.Context,param models.ParamClusterId) (interface{},string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return "",fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}

	events,err := client.ClusterEvent()
	if err != nil {
		middleware.Log(c).Errorf("get cluster %v event failed:%v\n",param.ClusterId,err)
		return "",fmt.Sprintf("get cluster %v event failed:%v\n",param.ClusterId,err),err 
	}

	return events,fmt.Sprintf("get cluster %v event success",param.ClusterId),nil 
}

func DeleteResource(c *gin.Context,param models.ParamReourceYaml) (string,error) {
	var err error 
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}

	switch param.ResourceType {
	case "node":
		err = client.DeleteNode(param.ResourceName)
	case "pod":
		err = client.DeletePod(param.NameSpace,param.ResourceName)
	case "deployment":
		err = client.DeleteDeployment(param.NameSpace,param.ResourceName)
	case "statefulset":
		err = client.DeleteStatefulSet(param.NameSpace,param.ResourceName)
	case "daemonset":
		err = client.DeleteDaemonSet(param.NameSpace,param.ResourceName)
	case "job":
		err = client.DeleteJob(param.NameSpace,param.ResourceName)
	case "cronjob":
		err = client.DeleteCronJob(param.NameSpace,param.ResourceName)
	case "svc":
		err = client.DeleteService(param.NameSpace,param.ResourceName)
	case "ingress":
		err = client.DeleteIngress(param.NameSpace,param.ResourceName)
	case "secret":
		err = client.DeleteSecret(param.NameSpace,param.ResourceName)
	case "pvc":
		err = client.DeletePvc(param.NameSpace,param.ResourceName)
	case "pv":
		err = client.DeletePv(param.ResourceName)
	case "storageclass":
		err = client.DeleteStorageClass(param.ResourceName)
	case "serviceaccount":
		err = client.DeleteServiceAccount(param.NameSpace,param.ResourceName)
	case "role":
		err = client.DeleteRole(param.NameSpace,param.ResourceName)
	case "rolebinding":
		err = client.DeleteRoleBinding(param.NameSpace,param.ResourceName)
	default:
		middleware.Log(c).Errorf("resource type %v invalid",param.ResourceType)
		return fmt.Sprintf("resource type %v invalid",param.ResourceType),errors.New(fmt.Sprintf("resource type %v invalid",param.ResourceType))
	}

	return fmt.Sprintf("delete resource %v by cluster %v success",param.ResourceName,param.ClusterId),nil 
}

func WorkloadRollUpdate(c *gin.Context,param models.ParamReourceInfo) (string,error) {
	var err error 
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	switch param.ResourceType{
	case "deployment":
		err = client.DeployRollUpdate(param.ResourceName,param.NameSpace)
	case "statefulset":
		err = client.StatefulSetRollUpdate(param.ResourceName,param.NameSpace)
	case "daemonset":
		err = client.DaemonSetRollUpdate(param.ResourceName,param.NameSpace)
	default:
		middleware.Log(c).Errorf("resource type %v invalid",param.ResourceType)
		return fmt.Sprintf(fmt.Sprintf("resource type %v invalid",param.ResourceType)),errors.New(fmt.Sprintf("resource type %v invalid",param.ResourceType))
	}
	if err != nil {
		middleware.Log(c).Errorf("resoure name %v roll restart by cluster %v failed:%v\n",param.ResourceName,param.ClusterId,err)
		return fmt.Sprintf("resoure name %v roll restart by cluster %v failed:%v\n",param.ResourceName,param.ClusterId,err),err 
	}
	return fmt.Sprintf("deployment %v roll restart by cluster %v success",param.ResourceName,param.ClusterId),nil 
}
func CreateConfigMap(c *gin.Context,param models.ParamCreateConfigmap) (string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	err = client.CreateConfigMap(param.NameSpace,param.ConfigMapName,param.KV)
	if err != nil {
		middleware.Log(c).Errorf("cluster:%v,namespace:%v,create configmap %v failed:%v\n",param.ClusterId,param.NameSpace,param.ConfigMapName,err)
		return fmt.Sprintf("cluster:%v,namespace:%v,create configmap %v failed:%v\n",param.ClusterId,param.NameSpace,param.ConfigMapName,err),err 
	}
	return fmt.Sprintf("cluster:%v,namespace:%v,create configmap %v success",param.ClusterId,param.NameSpace,param.ConfigMapName),nil 
}
func UpdateConfigMap(c *gin.Context,param models.ParamCreateConfigmap) (string,error) {
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	err = client.UpdateConfigMap(param.NameSpace,param.ConfigMapName,param.KV)
	if err != nil {
		middleware.Log(c).Errorf("cluster:%v,namespace:%v,update configmap %v failed:%v\n",param.ClusterId,param.NameSpace,param.ConfigMapName,err)
		return fmt.Sprintf("cluster:%v,namespace:%v,update configmap %v failed:%v\n",param.ClusterId,param.NameSpace,param.ConfigMapName,err),err 
	}
	return fmt.Sprintf("cluster:%v,namespace:%v,udpate configmap %v success",param.ClusterId,param.NameSpace,param.ConfigMapName),nil 
}
func CreateSecret(c *gin.Context,param models.ParamCreateSecret) (string,error) {
	var (
		err error 
		secretmap map[string][]byte = make(map[string][]byte,0)
	)
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	if param.IsEncrypt {
		for k,v := range param.KV {
			secretmap[k] = []byte(base64.StdEncoding.EncodeToString([]byte(v)))
		}
	}else{
		for k,v := range param.KV {
			secretmap[k] = []byte(v)
		}
	}
	switch param.Type {
	case "tlscert":
		if param.Cert == "" || param.Key == "" {
			middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v,cert %v or key %v invalid",param.Cert,param.Key)
			return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,cert %v or key %v invalid",param.Cert,param.Key),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,cert %v or key %v invalid",param.Cert,param.Key))
		}
		err = client.CreateSecretByTlsCert(param.NameSpace,param.SecretName,param.Cert,param.Key)
	case "imagecert":
		if param.ImageRepositoryUrl == "" || param.RepositoryUser == "" {
			middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v,url %v or user %v invalid",param.ImageRepositoryUrl,param.RepositoryUser)
			return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,url %v or user %v invalid",param.ImageRepositoryUrl,param.RepositoryUser),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,url %v or user %v invalid",param.ImageRepositoryUrl,param.RepositoryUser))
		} 
		err = client.CreateSecretByImageCert(param.NameSpace,param.SecretName,param.ImageRepositoryUrl,param.RepositoryUser,param.RepositoryPassword)
	case "opaque":
		if secretmap == nil {
			middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v, secretmap %v invalid",secretmap)
			return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,secretmap %v invalid",secretmap),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,secretmap %v invalid",secretmap))			
		}
		err = client.CreateSecretByOpaque(param.NameSpace,param.SecretName,secretmap)
	default:
		middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v invalid",param.Type)
		return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v invalid",param.Type),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v invalid",param.Type))
	}
	if err != nil {
		middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v, create secret %v failed:%v\n",param.ClusterId,param.NameSpace,param.Type,param.SecretName,err)
		return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v, create secret %v failed:%v\n",param.ClusterId,param.NameSpace,param.Type,param.SecretName,err),err 
	}
	return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v, create secret %v success",param.ClusterId,param.NameSpace,param.Type,param.SecretName),nil 
}
func UpdateSecret(c *gin.Context,param models.ParamCreateSecret) (string,error) {
	var (
		err error 
		secretmap map[string][]byte = make(map[string][]byte,0)
	)
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	if param.IsEncrypt {
		for k,v := range param.KV {
			secretmap[k] = []byte(base64.StdEncoding.EncodeToString([]byte(v)))
		}
	}else{
		for k,v := range param.KV {
			secretmap[k] = []byte(v)
		}
	}
	switch param.Type {
	case "tlscert":
		if param.Cert == "" || param.Key == "" {
			middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v,cert %v or key %v invalid",param.Cert,param.Key)
			return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,cert %v or key %v invalid",param.Cert,param.Key),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,cert %v or key %v invalid",param.Cert,param.Key))
		}
		err = client.UpdateSecretByTlsCert(param.NameSpace,param.SecretName,param.Cert,param.Key)
	case "imagecert":
		if param.ImageRepositoryUrl == "" || param.RepositoryUser == "" {
			middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v,url %v or user %v invalid",param.ImageRepositoryUrl,param.RepositoryUser)
			return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,url %v or user %v invalid",param.ImageRepositoryUrl,param.RepositoryUser),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,url %v or user %v invalid",param.ImageRepositoryUrl,param.RepositoryUser))
		} 
		err = client.UpdateSecretByImageCert(param.NameSpace,param.SecretName,param.ImageRepositoryUrl,param.RepositoryUser,param.RepositoryPassword)
	case "opaque":
		if secretmap == nil {
			middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v, secretmap %v invalid",secretmap)
			return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,secretmap %v invalid",secretmap),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v,secretmap %v invalid",secretmap))			
		}
		err = client.UpdateSecretByOpaque(param.NameSpace,param.SecretName,secretmap)
	default:
		middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v invalid",param.Type)
		return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v invalid",param.Type),errors.New(fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v invalid",param.Type))
	}
	if err != nil {
		middleware.Log(c).Errorf("cluster:%v,namespace:%v,secret type:%v, update secret %v failed:%v\n",param.ClusterId,param.NameSpace,param.Type,param.SecretName,err)
		return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v, update secret %v failed:%v\n",param.ClusterId,param.NameSpace,param.Type,param.SecretName,err),err 
	}
	return fmt.Sprintf("cluster:%v,namespace:%v,secret type:%v, update secret %v success",param.ClusterId,param.NameSpace,param.Type,param.SecretName),nil 
}

func CreateResourceByYaml(c *gin.Context,param models.ParamCreateResourceYaml) (string,error ) {
	var (
		resource interface{}
		err error 
	)
	client,err := GetK8sClientByClusterId(c,param.ClusterId)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}
	switch param.ResourceType{
	case "namespace":
		var namespace coreV1.Namespace
		err = pkg.CheckYamlFormat(param.ResourceYaml,namespace)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = namespace
	case "deployment":
		var deployment appsV1.Deployment
		err = pkg.CheckYamlFormat(param.ResourceYaml,deployment)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = deployment
	case "statefulset":
		var statefulset appsV1.StatefulSet
		err = pkg.CheckYamlFormat(param.ResourceYaml,statefulset)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = statefulset
	case "daemonset":
		var daemonset appsV1.DaemonSet
		err = pkg.CheckYamlFormat(param.ResourceYaml,daemonset)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = daemonset
	case "job":
		var job batchV1.Job
		err = pkg.CheckYamlFormat(param.ResourceYaml,job)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = job
	case "cronjob":
		var cronjob batchV1.CronJob
		err = pkg.CheckYamlFormat(param.ResourceYaml,cronjob)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = cronjob
	case "pod":
		var pod coreV1.Pod
		err = pkg.CheckYamlFormat(param.ResourceYaml,pod)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = pod
	case "svc":
		var svc coreV1.Service
		err = pkg.CheckYamlFormat(param.ResourceYaml,svc)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = svc
	case "ingress":
		var ingress networkV1.Ingress
		err = pkg.CheckYamlFormat(param.ResourceYaml,ingress)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = ingress
	case "configmap":
		var configmap coreV1.ConfigMap
		err = pkg.CheckYamlFormat(param.ResourceYaml,configmap)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = configmap
	case "secret":
		var secret coreV1.Secret
		err = pkg.CheckYamlFormat(param.ResourceYaml,secret)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = secret
	case "pvc":
		var pvc coreV1.PersistentVolumeClaim
		err = pkg.CheckYamlFormat(param.ResourceYaml,pvc)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = pvc
	case "pv":
		var pv coreV1.PersistentVolume
		err = pkg.CheckYamlFormat(param.ResourceYaml,pv)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = pv
	case "storageclass":
		var storageclass storageV1.StorageClass
		err = pkg.CheckYamlFormat(param.ResourceYaml,storageclass)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = storageclass
	case "serviceaccount":
		var serviceaccount coreV1.ServiceAccount
		err = pkg.CheckYamlFormat(param.ResourceYaml,serviceaccount)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = serviceaccount
	case "role":
		var role rbacV1.Role
		err := pkg.CheckYamlFormat(param.ResourceYaml,role)
		if err != nil {
			middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
			return fmt.Sprintf("resource yaml format invalid:%v\n",err),err 
		}
		resource = role
	default:
		middleware.Log(c).Errorf("resource type %v invalid",param.ResourceType)
		return fmt.Sprintf("resource type %v invalid",param.ResourceType),errors.New(fmt.Sprintf("resource type %v invalid",param.ResourceType))
	}
	err = client.CreateResourceByYaml(resource)
	if err != nil {
		middleware.Log(c).Errorf("resource yaml format invalid:%v\n",err)
		return fmt.Sprintf("resource yaml format invalid:%v\n",err),err
	}
	return "",nil 
}
