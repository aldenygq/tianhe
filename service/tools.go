package service

import (
	"tianhe/pkg"
	"tianhe/models"
	"tianhe/middleware"
	"time"
	"unicode"
	"sort"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/tools/clientcmd"
    //"k8s.io/client-go/tools/clientcmd/api"
)
//校验二维slce存在重复元素
func CheckDuplicates(slice [][]string) bool {
	allElements := make([]string, 0)
	for _, innerSlice := range slice {
		allElements = append(allElements, innerSlice...)
	}
	sort.Strings(allElements)

	for i := 1; i < len(allElements); i++ {
		if allElements[i-1] == allElements[i] {
			return true
		}
	}
	return false
}
 
func CompareTwoDay(day1,day2 string) int64{
	// 创建两个日期
	d1, _ := time.Parse("2006/01/02", day1)
	d2, _ := time.Parse("2006/01/02", day2)
	
	// 比较日期
	if d1.Before(d2) {
		return 0
	} else if d1.After(d2) {
		return 1
	} else {
		return 2
	}
	return 3
}

func getDay() string {
	now := time.Now().Format("2006-01-02")
	
	// fmt.Println("Current date:", now.Format("2006-01-02"))
	return now
}
func IsContainChinese(str string) bool {
	for _, r := range str {
		if unicode.In(r, unicode.Scripts["Han"]) {
			return true
		}
	}
	return false
}


func GetK8sClientByClusterId(c *gin.Context,clusterid string) (*pkg.K8sClient,error) {
	var (
		cluster *models.K8sCluster = &models.K8sCluster{}
	)
	cluster.ClusterId = clusterid
	err := cluster.GetClusterById()
	if err != nil {
		middleware.Log(c).Errorf("get cluster info by id %v failed:%v\n",clusterid,err)
		return nil,err 
	}
	client,err := pkg.NewK8sClient(cluster.Kubeconfig)
	if err != nil {
		middleware.Log(c).Errorf("new k8s cluster %v client failed:%v\n",clusterid,err)
		return nil,err 
	}
	return client,nil 
}

func GetUserByKubeconfig(c *gin.Context,kubeconfig string) (string,error) {
	// 使用clientcmd.RESTConfigFromKubeConfig将kubeconfig字符串解析为api.Config结构
	config, err := clientcmd.Load([]byte(kubeconfig))
	if err != nil {
		middleware.Log(c).Errorf("load kubeconfig failed :%v\n",err)
		return "",err 
	}

	// 获取当前上下文的用户信息
	currentContext := config.CurrentContext
	ctx, ok := config.Contexts[currentContext]
	if !ok {
		middleware.Log(c).Errorf("get kubeconfig context info failed :%v\n",err)
		return "",err 
	}
	user, ok := config.AuthInfos[ctx.AuthInfo]
    if !ok {
		middleware.Log(c).Errorf("get kubeconfig user  info failed :%v\n",err)
		return "",err 
    }

	return user.Username,nil 
}