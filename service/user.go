package service

import (
	"errors"
	"fmt"
	"tianhe/config"
	"tianhe/middleware"
	"tianhe/models"
	"tianhe/pkg"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/aldenygq/toolkits"
)

// 用户注册
func UserRegister(c *gin.Context,param models.ParamUserRegister) (string, error) {
	var (
		user *models.Users = &models.Users{}
		//expire *models.UserTokenExpire = &models.UserTokenExpire{}
	)
	//校验密码复杂度
	err := pkg.ValidatePassword(param.PassWord) 
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("password %v complexity invalid:%v", param.PassWord,err))
		return fmt.Sprintf(err.Error()), err
	}
	//密码加密
	encryptpwd, err := toolkits.Encrypt([]byte(param.PassWord), config.Conf.Util.InitKey)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("encrypt password error: %v\n", err))
		return fmt.Sprintf("密码不合法"), err
	}
	user.EnName = param.EnName
	user.Mobile = param.Mobile
	user.Status = 1
	user.Password = encryptpwd
	user.Email = param.Email
	user.Ctime = time.Now().Unix()
	user.CreateType = "register"
	user.Creator = param.EnName
	user.ExpireTime = config.Conf.Route.AuthTokenExpire
	//middleware.Logf(c,fmt.Sprintf("requestid:%v",c.Get("X-Request-Id")))
	//用户是否存在
	u,result,err := user.IsExist()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("%v", err))
		return fmt.Sprintf("%v", err), err
	}
	if result {
		//logrus.Errorf("user %v exist",u)
		middleware.Log(c).Errorf(fmt.Sprintf("user %v exist",u))
		return fmt.Sprintf("user %v exist",u), errors.New(fmt.Sprintf("user %v exist",u))
	}

	//middleware.Logf(c,fmt.Sprintf("requestid:%v",c.Get("X-Request-Id")))
	//创建用户
	err = user.Create()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("user %v register failed:%v\n", user.EnName, err))
		return  err.Error(), errors.New(fmt.Sprintf("user %v register failed:%v\n", user.EnName, err))
	}
	return fmt.Sprintf("注册成功"), nil
}

// 获取个人信息
func UserInfo(c *gin.Context,token string) (*models.Users, string, error) {
	var user *models.Users = &models.Users{}
	ret, err := middleware.ParseToken(token)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("parse token failed:%v", err))
		return nil, fmt.Sprintf("解析token失败,失败原因:%v\n", err), err
	}
	middleware.Log(c).Infof(fmt.Sprintf("uname:%v",ret.UEnName))
	user.EnName = ret.UEnName
	err = user.GetByUname()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("get user info failed:%v", err))
		return nil, fmt.Sprintf("获取个人信息失败,失败原因:%v", err), err
	}
	user.Mobile = string([]byte(user.Mobile)[0:3]) + "****" + string([]byte(user.Mobile)[6:])
	user.Password = "************"
	return user, fmt.Sprintf("获取个人信息成功"), nil
}
func SetTokenExpire(c *gin.Context,accessToken string,param models.ParamSetUserTokenExpire) (string, error) {
	var user *models.Users = &models.Users{}
	ret, err := middleware.ParseToken(accessToken)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("parse token failed:%v", err))
		return fmt.Sprintf("解析token失败,失败原因:%v", err), err
	}
	middleware.Log(c).Infof(fmt.Sprintf("uname:",ret.UEnName))
	user.EnName = ret.UEnName
	user.Mtime = time.Now().Unix()
	user.ExpireTime = param.ExpireTime
	err = user.SetUserTokenExpire()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("set user %v token expire time failed:%v",ret.UEnName,err))
		return fmt.Sprintf("set user %v token expire time failed:%v",ret.UEnName,err),err 
	}
	return fmt.Sprintf("set token expire success"),nil 
}
// 修改密码
func ModifyPassword(c *gin.Context,token, pwd string) (string, error) {
	var user *models.Users = &models.Users{}
	ret, err := middleware.ParseToken(token)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("parse token failed:%v\n", err))
		return fmt.Sprintf("解析token失败,失败原因:%v\n", err), err
	}
	err = pkg.ValidatePassword(pwd) 
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("password %v complexity invalid:%v", pwd,err))
		return fmt.Sprintf(err.Error()), err
	}
	encryptpwd, err := toolkits.Encrypt([]byte(pwd), config.Conf.Util.InitKey)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("encrypt password error: %v", err))
		return fmt.Sprintf("密码不合法"), err
	}
	user.EnName = ret.UEnName
	user.Password = encryptpwd
	err = user.UpdateByEnName()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("get user info failed:%v\n", err))
		return fmt.Sprintf("修改密码失败,失败原因:%v\n", err), err
	}

	return fmt.Sprintf("修改密码成功"), nil
}

// 忘记密码
func ForgotPassword(c *gin.Context,param models.ParamForgotPassword) (string, error) {
	var (
		user *models.Users = &models.Users{}
	)
	user.Mobile = param.Mobile

	//用户是否存在
	u,result,err := user.IsExist()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("%v", err))
		return fmt.Sprintf("%v", err), err
	}
	if !result {
		middleware.Log(c).Errorf(fmt.Sprintf("user %v not exist",u))
		return fmt.Sprintf("user %v not exist",u), errors.New(fmt.Sprintf("user %v not exist",u))
	}

	err = pkg.ValidatePassword(param.PassWord) 
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("password %v complexity invalid:%v", param.PassWord,err))
		return fmt.Sprintf(err.Error()), err
	}
	encryptpwd, err := toolkits.Encrypt([]byte(param.PassWord), config.Conf.Util.InitKey)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("encrypt password error: %v", err))
		return fmt.Sprintf("密码不合法"), err
	}
	user.Password = encryptpwd
	user.Mtime = time.Now().Unix()
	err = user.UpdateByMobile()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("update user %v password failed: %v\n", user.Mobile, err))
		return fmt.Sprintf("修改密码失败,请联系系统管理员查询。"), err
	}

	return fmt.Sprintf("修改密码成功,请前往登录页面登陆。"), nil
}

// 启用/禁用用户
func ModifyUserStatus(c *gin.Context,param models.ParamModifyUserStatus) (string, error) {
	var (
		user               *models.Users = &models.Users{}
		opreate, opreatecn string
	)
	user.Status = param.Status
	user.EnName = param.EnName
	switch param.Status {
	case 1:
		opreate = "enable"
		opreatecn = "启用"
	case 2:
		opreate = "disable"
		opreatecn = "禁用"
	case 3:
		opreate = "delete"
		opreatecn = "删除"
	default:
		middleware.Log(c).Errorf(fmt.Sprintf("opreate %v invalid",opreate))
		return fmt.Sprintf("opreate %v invalid",opreate), errors.New(fmt.Sprintf("opreate %v invalid",opreate))
	}
	err := user.UpdateByEnName()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("%v user %v password failed: %v\n", opreate, user.EnName, err))
		return fmt.Sprintf("%v用户失败",opreatecn), err
	}

	err = middleware.DelToken(user.EnName)
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("user logout failed:%v\n", err))
	}

	middleware.Log(c).Infof(fmt.Sprintf("%v user %v password success", opreate, user.EnName, err))
	return fmt.Sprintf("%v用户成功",opreatecn), err
}


// 获取用户列表
func UserList(c *gin.Context,param models.ParamUserList) (map[string]interface{}, string, error) {
	var (
		users map[string]interface{} = make(map[string]interface{}, 0)
		user  *models.Users          = &models.Users{}
	)
	user.Mobile = param.Mobile
	user.EnName = param.EnName
	user.Status = param.Status
	count, us, err := user.List()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("get user list failed: %v\n", err))
		return nil, fmt.Sprintf("获取用户列表失败"), err
	}
	users["count"] = count
	users["users"] = us

	return users, fmt.Sprintf("获取用户列表成功"), nil
}

func ModifyUserInfo(c *gin.Context,param models.ParamModifyUserInfo) (string, error) {
	var user *models.Users = &models.Users{}
	user.Mobile = param.Mobile
	user.EnName = param.EnName
	user.Email = param.Email
	err := user.UpdateByEnName()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("modify user %v info failed: %v\n", user.EnName, err))
		return fmt.Sprintf("修改个人信息失败"), err
	}

	middleware.Log(c).Infof(fmt.Sprintf("modify user %v info success", user.EnName, err))
	return fmt.Sprintf("修改个人信息成功"), err
}


func Unregister(c *gin.Context,accessToken string) (string,error) {
	var user *models.Users = &models.Users{}
	ret,err := middleware.ParseToken(accessToken)
	if err != nil || ret == nil {
		middleware.Log(c).Errorf(fmt.Sprintf("user not login:%v",err))
		return fmt.Sprintf("user not login"),errors.New("user not login") 
	}

	user.EnName = ret.UEnName
	user.Status = 3
	user.Mtime = time.Now().Unix()
	err = user.UpdateByEnName()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("update user %v status failed:%v",err))
		return fmt.Sprintf("unregister user %v failed:%v",user.EnName,err),err 
	}

	err = middleware.RedisClient.Del(user.EnName).Err()
	if err != nil {
		middleware.Log(c).Errorf(fmt.Sprintf("delete user %v token failed:%v",user.EnName,err))
	}

	return fmt.Sprintf("user %v unregister success"),nil 
}