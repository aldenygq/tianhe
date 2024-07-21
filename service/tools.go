package service

import (
	"tianhe/pkg"
	"tianhe/models"
	"tianhe/middleware"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

func CompareTwoDay(day1,day2 string) int64{
	// 创建两个日期
	d1, _ := time.Parse("2006-01-02", day1)
	d2, _ := time.Parse("2006-01-02", day2)
	
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
		middleware.LogErr(c).Errorf("get cluster info by id %v failed:%v\n",clusterid,err)
		return nil,err 
	}
	client,err := pkg.NewK8sClient(cluster.Kubeconfig)
	if err != nil {
		middleware.LogErr(c).Errorf("new k8s cluster %v client failed:%v\n",clusterid,err)
		return nil,err 
	}
	return client,nil 
}
