package pkg

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	
	"github.com/aldenygq/toolkits"
	"oncall/config"
	"oncall/middleware"
)

const (
	SMSTPL = "【%v】您的验证码为：%v，若非本人操作,请忽略该短信！"
)

func SmsCheck(key, code string) bool {
	val,err := middleware.RedisClient.Get(key).Result()
	// 用完后将连接放回连接池
	defer middleware.RedisClient.Close()
	if err != nil || val != code {
		return false
	}

	return true
}
func SmsSet(key, val string,t int64) error {
	//key = cache.RedisSuf + key
	// 从池里获取连接
	err := middleware.RedisClient.Set(key,val,60 * time.Second).Err()
	if err != nil {
		middleware.Logger.Errorf("set key to redis failed:%v\n",err)
		return err
	}
	// 用完后将连接放回连接池
	defer middleware.RedisClient.Close()

	return nil
}

func HttpPostForm(url string, data url.Values) (string, error) {
	
	resp, err := http.PostForm(url, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//发送短信
func SendSms(mobile, msg string) error {
	var err error
	if AClient != nil {
		err = AClient.SendSms(mobile,ASignName,ATplnum,msg)
		if err != nil {
			middleware.Logger.Errorf("send sms by aliyun sms failed:%v\n",err)
			return err
		}
	}else if TClient != nil {
		err = TClient.TencentSendSmsCode(Appid,TSignName,TTplnum,mobile,msg)
		if err != nil {
			middleware.Logger.Errorf("send sms by tencent sms failed:%v\n",err)
			return err
		}
	}

	return err
}

var (
	AClient *toolkits.AliyunSmsClient
	ASignName,TSignName string
	ATplnum,TTplnum string
	Appid string
	TClient *toolkits.TencentSmsClient
)
func InitSms() {
	var err error
	if len(config.Conf.Sms) <= 0 {
		log.Printf("sms info invalid")
		os.Exit(-1)
	}
	
	for _,v := range config.Conf.Sms {
		switch v.Type {
		case "aliyun":
			AClient,err = toolkits.NewAliyunSmsClient(v.Ak,v.Sk,v.Endpoint)
			ASignName = v.SignName
			ATplnum = v.TplNum
		case "tencent":
			TClient,err = toolkits.NewTencentSmsClient(v.Ak,v.Sk,v.SignName,v.Region)
			TSignName = v.SignName
			TTplnum = v.TplNum
			Appid = v.AppId
		}
		if err != nil {
			log.Printf("new %v sms client failed:%v",err)
			os.Exit(-1)
		}
	}
	return
}