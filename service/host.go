package service

import (
	"errors"
	"fmt"
	"time"
	"tianhe/middleware"
	"tianhe/models"
	"tianhe/pkg"
	"github.com/gin-gonic/gin"
)

func AddHost(c *gin.Context, param models.ParamAddHost) (string,error) {
	var (
		host *models.Host = &models.Host{}
		sshinfo *models.HostSshAuth = &models.HostSshAuth{}
	)
	host.Creator = param.Creator
	host.Ctime = time.Now().Unix()
	host.HostIp = param.HostIp
	host.HostName = param.HostName
	host.HostType = param.HostType
	host.Os = param.Os
	host.OsVersion = param.OsVersion
	host.SshPort = param.Port
	hid,_ := pkg.GenerateUniqueID()
	host.HostId = hid
	err := host.Create()
	if err != nil {
		middleware.Log(c).Errorf("add host %v failed:%v\n",host.HostIp,err)
		return fmt.Sprintf("add host %v failed:%v\n",host.HostIp,err),err 
	}

	sshinfo.AuthType = param.AuthType
	sshinfo.Creator = param.Creator
	sshinfo.Ctime = time.Now().Unix()
	sshinfo.HostId = host.HostId
	sshinfo.User = "root"
	switch param.AuthType {
	case "passwd":
		sshinfo.Password = param.Password
	case "ssh_key":
		sshinfo.PrivateKey = param.PrivateKey
	default :
		middleware.Log(c).Errorf("author type is %v,please provide the correct authentication method")
		return fmt.Sprintf("author type is %v,please provide the correct authentication method"),errors.New(fmt.Sprintf("author type is %v,please provide the correct authentication method"))
	}
	err = sshinfo.Create()
	if err != nil {
		middleware.Log(c).Errorf("add host user %v failed:%v\n",sshinfo.User,err)
		return fmt.Sprintf("add host user %v failed:%v\n",sshinfo.User,err),err 
	}

	middleware.Log(c).Infof("add host %v success",host.HostIp)
	return fmt.Sprintf("add host %v success",host.HostIp),nil 
}

func DelHost(c *gin.Context, param models.ParamDelHost) (string,error) {
	var host *models.Host = &models.Host{}
	host.HostId = param.HostId

	err := host.Delete()
	if err != nil {
		middleware.Log(c).Errorf("delete host %v failed:%v\n",host.HostId,err)
		return fmt.Sprintf("delete host %v failed:%v\n",host.HostId,err),err 
	}

	middleware.Log(c).Infof("delete host %v success",host.HostId)
	return fmt.Sprintf("delete host %v success",host.HostId),nil 
}

func HostInfo(c *gin.Context, param models.ParamHostInfo) (*models.Host,string,error) {
	var host *models.Host = &models.Host{}
	host.HostId = param.HostId
	err := host.GetHostById()
	if err != nil {
		middleware.Log(c).Errorf("get host %v info failed:%v\n",host.HostId,err)
		return nil,fmt.Sprintf("delete host %v failed:%v\n",host.HostId,err),err 
	}
	middleware.Log(c).Infof("get host %v success",host.HostId)
	return host,fmt.Sprintf("delete host %v success",host.HostId),nil 
}