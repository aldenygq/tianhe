package service

import (
	"errors"
	"fmt"

	"github.com/aldenygq/toolkits"
	"github.com/gin-gonic/gin"
	"tianhe/middleware"
	"tianhe/models"
	"tianhe/pkg"
	"tianhe/config"
)

var rClient = middleware.RedisClient

// 登录
func Login(c *gin.Context, param models.ParamLogin) (string,string, error) {
	var (
		user *models.Users = &models.Users{}
		//expire *models.UserTokenExpire = &models.UserTokenExpire{}
		expireTime int64
	)
	switch param.Type {
	//手机验证码
	case "verify":
		user.Mobile = param.Mobile
		err := user.GetByMobile()
		if err != nil || user.Status != 1 {
			middleware.Log(c).Errorf(fmt.Sprintf("user %v status unusual", param.Mobile))
			return "",fmt.Sprintf("用户状态异常,请联系管理员查询！"),errors.New("用户状态异常,请联系管理员查询！")
		}
		//验证手机号是否合法
		if toolkits.CheckMobile(param.Mobile) {
			middleware.Log(c).Errorf(fmt.Sprintf("mobile %v invalid", param.Mobile))
			return "",fmt.Sprintf("手机号不合法"),errors.New(" 手机号不合法")
		}
		//验证码校验
		result,err := pkg.SmsCheck(param.Mobile, param.VerifyCode) 
		if err != nil {
			middleware.Log(c).Errorf(fmt.Sprintf("mobild %v verify code check failed:%v",param.Mobile,err))
			return "",fmt.Sprintf("验证码校验失败，请联系管理员排查"),err
		}
		if !result {
			middleware.Log(c).Errorf(fmt.Sprintf("mobile %v verify code %v invalid", param.Mobile, param.VerifyCode))
			return "",fmt.Sprintf("验证码无效"),errors.New("验证码无效")
		}
	//账号密码
	case "account":
		//密码认证
		user.EnName = param.EnName
		err := user.GetByUname()
		if err != nil || user.Status != 1 {
			middleware.Log(c).Errorf(fmt.Sprintf("user %v status unusual", user.EnName))
			return "",fmt.Sprintf("用户状态异常,请联系管理员查询！"),errors.New("用户状态异常,请联系管理员查询！")
		}
		decryptpwd, err := pkg.Decrypt(user.Password, config.Conf.Util.InitKey)
		if err != nil {
			middleware.Log(c).Errorf(fmt.Sprintf("user authentication failed:%v", err))
			return "",fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err 
		}
		if param.PassWord != string(decryptpwd) {
			middleware.Log(c).Errorf(fmt.Sprintf("user passwd :%v invalid,right pwd:%v", param.PassWord, string(decryptpwd)))
			return "",fmt.Sprintf("密码不正确"),errors.New ("密码不正确")
		}
	default:
		middleware.Log(c).Errorf(fmt.Sprintf("login type %v invalid", param.Type))
		return "",fmt.Sprintf("登陆类型不合法"),errors.New("登陆 类型不合法")
	}

	err := user.GetTokenExpireByUser()
	if err != nil {
		expireTime = config.Conf.Route.AuthTokenExpire
	} else {
		expireTime = user.ExpireTime
	}
	//登陆
	token,err := middleware.DoLogin(c, user.EnName,expireTime)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("user :%v login failed:%v\n", user.EnName))
		return "",fmt.Sprintf("登录失败"),err 
	}
	middleware.Log(c).Infof(fmt.Sprintf("user :%v login success", user.EnName))
	return token,fmt.Sprintf("登陆成功"),nil 
}

func CheckUseLoginByUname(c *gin.Context,param models.ParamUserEnName) (string,string,error) {
	val,err := middleware.RedisClient.Get(param.EnName).Result()
	if err != nil || val == "" {
		middleware.Log(c).Errorf(fmt.Sprintf("user %v not login",param.EnName))
		return "",fmt.Sprintf("用户未登录"),err 
	}
	_,err = middleware.ParseToken(val)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("user %v not login",param.EnName))
		return "",fmt.Sprintf("用户未登录"),err 
	}

	return val,fmt.Sprintf("用户已登录"),nil 
}

func SendSms(c *gin.Context,param models.ParamMobile) (string, error) {
	//生成随机数
	code := toolkits.GetRandomNum(6)
	msg := fmt.Sprintf(pkg.SMSTPL, param.Mobile, code)
	err := pkg.SendSms(param.Mobile, msg)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("send sms to mobile %v failed:%v\n", param.Mobile, err))
		return fmt.Sprintf("短信验证码发送失败"),err 
	}

	err = pkg.SmsSet(param.Mobile, code, 60)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("mobile %v set sms code %v to redis failed:%v\n", param.Mobile, code, err))
		return fmt.Sprintf("短信验证码发送失败"),err 
	}
	return fmt.Sprintf("短信验证码发送成功"),err 
}
