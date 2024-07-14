package service

import (
	"encoding/base64"
	"fmt"
	"tianhe/middleware"
	"tianhe/models"
	"tianhe/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterCluster(c *gin.Context, param models.ParamRegisterCluster) (string,error) {
	var cluster *models.K8sCluster = &models.K8sCluster{}
	cluster.ClusterId = param.ClusterId
	cluster.ClusterName = param.ClusterName
	cluster.Creator = param.Creator
	cluster.Ctime = time.Now().Unix()
	cluster.Env = param.Env
	cluster.Kubeconfig = base64.StdEncoding.EncodeToString([]byte(param.Kubeconfig))


	err := cluster.Create()
	if err != nil {
		middleware.LogErr(c).Errorf("register k8s cluster %v failed:%v\n",param.ClusterName,err)
		return fmt.Sprintf("register k8s cluster %v failed:%v\n",param.ClusterName,err),err 
	}

	middleware.LogInfo(c).Infof("register k8s cluster %v success",param.ClusterName)
	return fmt.Sprintf("register k8s cluster %v success",param.ClusterName),nil 
}

func CreateNs(c *gin.Context, param models.ParamCreateNs) (string,error) {
	var (
		cluster *models.K8sCluster = &models.K8sCluster{}
	)
	cluster.ClusterId = param.ClusterId
	err := cluster.GetClusterById()
	if err != nil {
		middleware.LogErr(c).Errorf("get cluster info by id %v failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("get cluster info by id %v failed:%v\n",param.ClusterId,err),err 
	}
	client,err := pkg.NewK8sClient(cluster.Kubeconfig)
	if err != nil {
		middleware.LogErr(c).Errorf("new k8s cluster %v client failed:%v\n",param.ClusterId,err)
		return fmt.Sprintf("new k8s cluster %v client failed:%v\n",param.ClusterId,err),err 
	}

	err = client.CreateNs(param.NameSpace)
	if err != nil {
		middleware.LogErr(c).Errorf("k8s cluster %v create namespace failed:%v\n",param.ClusterId,param.NameSpace,err)
		return fmt.Sprintf("k8s cluster %v create namespace failed:%v\n",param.ClusterId,param.NameSpace,err),err 
	}
	return fmt.Sprintf("create ns %v sucsess",param.NameSpace),nil 
}

func ClusterList(c *gin.Context) ([]*models.K8sCluster,string,error) {
	var cluster *models.K8sCluster = &models.K8sCluster{}
	list,err :=cluster.List()
	if err != nil {
		middleware.LogErr(c).Errorf("get k8s cluster list failed:%v\n",err)
		return nil,fmt.Sprintf("get k8s cluster list failed:%v\n",err),err 
	}
	return list,fmt.Sprintf("get k8s cluster list success"),nil
} 
