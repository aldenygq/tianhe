package service

import (
	"encoding/base64"
	"fmt"
	"tianhe/middleware"
	"tianhe/models"
	//"tianhe/pkg"
	"time"
	"errors"
	//"strings"
	"github.com/gin-gonic/gin"
)
func AddCloudSecret(c *gin.Context,param models.ParamAddCloudSecret) (string,error) {
	var (
		secretinfo *models.CloudSecretInfo = &models.CloudSecretInfo{}
		err error 
	)
	secretinfo.AccessKey = param.AccessKey
	secretinfo.AccountOwner = param.AccountOwner
	secretinfo.Cloud = param.Cloud
	secretinfo.CloudAccount = param.CloudAccount
	secretinfo.CloudProduct = param.CloudProduct
	secretinfo.Creator = param.Creator
	secretinfo.Ctime = time.Now().Unix()
	secretinfo.Env = param.Env
	secretinfo.Status = 1
	secretinfo.SecreyKey = base64.StdEncoding.EncodeToString([]byte(param.SecretKey))
	if secretinfo.Exist() {
		middleware.Log(c).Errorf("accesskey:%v exist",param.AccessKey)
		return fmt.Sprintf("accesskey:%v exist",param.AccessKey),errors.New(fmt.Sprintf("accesskey:%v exist",param.AccessKey))
	}
	err = secretinfo.Create()
	if err != nil {
		middleware.Log(c).Errorf("cloud:%v,accesskey:%v,product:%v,account:%v,env:%v,create failed:%v\n",param.Cloud,param.AccessKey,param.ParamCloudProduct,param.CloudAccount,param.Env,err)
		return fmt.Sprintf("create failed"),err 
	}
	

	return fmt.Sprintf("create success"),err 
} 