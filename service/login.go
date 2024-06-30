package service

import (
	"errors"
	"fmt"
	
	"github.com/aldenygq/toolkits"
	"github.com/gin-gonic/gin"
	"oncall/middleware"
	"oncall/models"
	"oncall/pkg"
	//"oncall/databases"
	"oncall/config"
)
var rClient = middleware.RedisClient

//校验用户登录状态
func CheckUseLogin(token string) (bool,string,error) {
	var (
		user *models.Users = &models.Users{}
	)
	ret,err := middleware.ParseToken(token)
	if err != nil {
		middleware.Logger.Errorf("parse token %v failed:%v\n",token,err)
		return false,fmt.Sprintf("校验用户登录状态失败,请联系系统管理员查看。"),err
	}
	user.EnName = ret.UEnName
	//用户是否存在
	has,err := user.Exist()
	if !has {
		middleware.Logger.Errorf("user %v not exist",ret.UEnName)
		return false,fmt.Sprintf("用户不存在"),errors.New("用户不存在")
	}
	if err != nil {
		middleware.Logger.Errorf("user authentication failed:%v\n",err)
		return false,fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err
	}
	
	//用户是否禁用
	result,err := user.IsDisable()
	if result {
		middleware.Logger.Errorf("user %v is disabled",ret.UEnName)
		return false,fmt.Sprintf("用户已禁用"),errors.New("用户已禁用")
	}
	if err != nil {
		middleware.Logger.Errorf("user authentication failed:%v\n",err)
		return false,fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err
	}
	
	if !middleware.CheckToken(token) {
		middleware.Logger.Errorf("user %v no login")
		return false,fmt.Sprintf("用户未登录"),errors.New(fmt.Sprintf("user %v no login",user.EnName))
	}
	
	middleware.Logger.Infof("用户登陆正常")
	return true,fmt.Sprintf("用户登陆正常"),nil
}

//登录
func Login(c *gin.Context,param models.ParamLogin) (string,error) {
	var (
		user *models.Users = &models.Users{}
	)
	switch param.Type {
	//手机验证码
	case "verify":
		//验证手机号是否合法
		if toolkits.CheckMobile(param.Mobile) {
			middleware.Logger.Errorf("mobile %v invalid",param.Mobile)
			return fmt.Sprintf("手机号不合法"),errors.New("手机号不合法")
		}
		//验证码校验
		if !pkg.SmsCheck(param.Mobile, param.VerifyCode) {
			middleware.Logger.Errorf("mobile %v verify code %v invalid",param.Mobile,param.VerifyCode)
			return fmt.Sprintf("验证码无效"),errors.New("验证码无效")
		}
		user.Mobile = param.Mobile
	//账号密码
	case "account","email":
		//密码认证
		decryptpwd,err := pkg.Decrypt(param.PassWord,config.Conf.Util.InitKey)
		if err != nil {
			middleware.Logger.Errorf("user authentication failed:%v\n",err)
			return fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err
		}
		if param.PassWord != string(decryptpwd) {
			middleware.Logger.Errorf("user passwd :%v invalid,right pwd:%v\n",param.PassWord,string(decryptpwd))
			return fmt.Sprintf("密码不正确"),errors.New("密码不正确")
		}
		user.EnName = param.EnName
		user.Email = param.Email
	default:
		middleware.Logger.Errorf("login type %v invalid",param.Type)
		return fmt.Sprintf("登陆类型不合法"),errors.New("登陆类型不合法")
	}
	
	//用户是否禁用
	result,err := user.IsDisable()
	if result {
		middleware.Logger.Errorf("user %v is disabled",user.EnName)
		return fmt.Sprintf("用户已禁用"),errors.New("用户已禁用")
	}
	if err != nil {
		middleware.Logger.Errorf("user authentication failed:%v\n",err)
		return fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err
	}
	
	//用户是否存在
	has,err := user.Exist()
	if !has {
		middleware.Logger.Errorf("user %v not exist",user.EnName)
		return fmt.Sprintf("用户不存在"),errors.New("用户不存在")
	}
	if err != nil {
		middleware.Logger.Errorf("user authentication failed:%v\n",err)
		return fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err
	}
	//登陆
	err = middleware.DoLogin(c, user.EnName)
	if err != nil {
		middleware.Logger.Errorf("user :%v login failed:%v\n",user.EnName)
		return fmt.Sprintf("登录失败"),err
	}
	middleware.Logger.Errorf("user :%v login success",user.EnName)
	return fmt.Sprintf("登陆成功"),nil
}

func SendSms(param models.ParamMobile) (string,error) {
	//生成随机数
	code := toolkits.GetRandomNum(6)
	msg := fmt.Sprintf(pkg.SMSTPL,param.Mobile,code)
	err := pkg.SendSms(param.Mobile, msg)
	if err != nil {
		middleware.Logger.Errorf("send sms to mobile %v failed:%v\n",param.Mobile,err)
		return fmt.Sprintf("短信验证码发送失败"),err
	}
	
	err = pkg.SmsSet(param.Mobile,code,60)
	if err != nil {
		middleware.Logger.Errorf(" mobile %v set sms code %v to redis failed:%v\n",param.Mobile,code,err)
		return fmt.Sprintf("短信验证码发送失败"),err
	}
	return fmt.Sprintf("短信验证码发送成功"),err
}
