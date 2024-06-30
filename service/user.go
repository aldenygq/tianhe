package service

import (
	"fmt"
	"time"
	"errors"
	"oncall/models"
	"oncall/pkg"
	"oncall/config"
	"oncall/middleware"
	"github.com/aldenygq/toolkits"
)

//用户注册
func UserRegister(param models.ParamUserRegister) (string,error) {
	var (
		user *models.Users = &models.Users{}
	)
	if !pkg.ValidatePassword(param.PassWord) {
		fmt.Printf("password %v complexity invalid",param.PassWord)
		return fmt.Sprintf("密码复杂度不够"),errors.New(fmt.Sprintf("password %v complexity invalid",param.PassWord))
	}
	encryptpwd, err := toolkits.Encrypt([]byte(param.PassWord), config.Conf.Util.InitKey)
	if err != nil {
		fmt.Printf("encrypt password error: %v\n", err)
		return fmt.Sprintf("密码不合法"),err
	}
	user.EnName = param.EnName
	user.Mobile = param.Mobile
	user.CnName = param.CnName
	user.Mtime = time.Now().Unix()
	user.Status = 1
	user.Password = encryptpwd
	//用户是否禁用
	result,err := user.IsDisable()
	if result {
		middleware.Logger.Errorf("user %v is disabled",u)
		return fmt.Sprintf("用户已禁用"),errors.New("用户已禁用")
	}
	if err != nil {
		middleware.Logger.Errorf("user authentication failed:%v\n",err)
		return fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err
	}
	//用户是否存在
	has,err := user.Exist()
	if has {
		middleware.Logger.Errorf("user en name %v exist\n",user.EnName)
		return errors.New(fmt.Sprintf("user en name %v exist",user.EnName)),fmt.Sprintf("用户名已存在",err)
	}
	if err != nil {
		middleware.Logger.Errorf("check user en name %v is exist failed:%v\n",user.EnName,err)
		return errors.New(fmt.Sprintf("check user en name %v is exist failed:%v",user.EnName,err)),fmt.Sprintf("用户校验失败,失败原因",err)
	}
	
	err = user.Create()
	if err != nil {
		middleware.Logger.Errorf("user %v register failed:%v\n",user.EnName,err)
		return errors.New(fmt.Sprintf("user %v register failed:%v\n",user.EnName,err)),fmt.Sprintf("用户注册失败,失败原因",err)
	}
	
	return fmt.Sprintf("注册成功"),nil
}

//获取个人信息
func UserInfo(token string) (*models.Users,string,error) {
	var user *models.Users = &models.Users{}
	ret, err := middleware.ParseToken(token)
	if err != nil {
		middleware.Logger.Errorf("parse token failed:%v\n", err)
		return nil,fmt.Sprintf("解析token失败,失败原因:%v\n",err),err
	}
	
	user.EnName = ret.UEnName
	err = user.Get()
	if err != nil {
		middleware.Logger.Errorf("get user info failed:%v\n", err)
		return nil,fmt.Sprintf("获取个人信息失败,失败原因:%v\n",err),err
	}
	user.Mobile = string([]byte(user.Mobile)[0:3])+"****"+string([]byte(user.Mobile)[6:])
	user.Password = "************"
	return user,fmt.Sprintf("获取个人信息成功"),nil
}

//修改密码
func ModifyPassword(token,pwd string) (string,error) {
	var user *models.Users = &models.Users{}
	ret, err := middleware.ParseToken(token)
	if err != nil {
		middleware.Logger.Errorf("parse token failed:%v\n", err)
		return fmt.Sprintf("解析token失败,失败原因:%v\n",err),err
	}
	if !pkg.ValidatePassword(pwd) {
		fmt.Printf("password %v complexity invalid",pwd)
		return fmt.Sprintf("密码复杂度不够"),errors.New(fmt.Sprintf("password %v complexity invalid",pwd))
	}
	encryptpwd, err := toolkits.Encrypt([]byte(pwd), config.Conf.Util.InitKey)
	if err != nil {
		fmt.Printf("encrypt password error: %v\n", err)
		return fmt.Sprintf("密码不合法"),err
	}
	user.EnName = ret.UEnName
	user.Password = encryptpwd
	err = user.UpdateByEnName()
	if err != nil {
		middleware.Logger.Errorf("get user info failed:%v\n", err)
		return fmt.Sprintf("修改密码失败,失败原因:%v\n",err),err
	}
	
	return fmt.Sprintf("修改密码成功"),nil
}

//忘记密码
func ForgotPassword(param models.ParamForgotPassword) (string,error) {
	var (
		user *models.Users = &models.Users{}
	)
	user.Mobile = param.Mobile
	//用户是否禁用
	result,err := user.IsDisable()
	if result {
		middleware.Logger.Errorf("user %v is disabled",u)
		return fmt.Sprintf("用户已禁用"),errors.New("用户已禁用")
	}
	if err != nil {
		middleware.Logger.Errorf("user authentication failed:%v\n",err)
		return fmt.Sprintf("用户身份验证失败,请联系系统管理员查看"),err
	}
	
	has,err := user.Exist()
	if !has {
		middleware.Logger.Errorf("user mobile %v not exist\n",user.Mobile)
		return fmt.Sprintf("用户不存在"),errors.New(fmt.Sprintf("user mobile %v not exist\n",user.Mobile))
	}
	if err != nil {
		middleware.Logger.Errorf("check user mobile %v is exist failed:%v\n",user.Mobile,err)
		return fmt.Sprintf("校验用户信息失败,请联系系统管理员查询。"),errors.New(fmt.Sprintf("check user mobile %v failed:%v\n",user.Mobile,err))
	}
	if !pkg.ValidatePassword(param.PassWord) {
		fmt.Printf("password %v complexity invalid",param.PassWord)
		return fmt.Sprintf("密码复杂度不够"),errors.New(fmt.Sprintf("password %v complexity invalid",param.PassWord))
	}
	encryptpwd, err := toolkits.Encrypt([]byte(param.PassWord), config.Conf.Util.InitKey)
	if err != nil {
		middleware.Logger.Errorf(fmt.Printf("encrypt password error: %v\n", err))
		return fmt.Sprintf("密码不合法"),err
	}
	user.Password = encryptpwd
	err = user.UpdateByMobile()
	if err != nil {
		middleware.Logger.Errorf(fmt.Printf("update user %v password failed: %v\n",user.Mobile,err))
		return fmt.Sprintf("修改密码失败,请联系系统管理员查询。"),err
	}
	
	return fmt.Sprintf("修改密码成功,请前往登录页面登陆。"),nil
}

//启用/禁用用户
func ModifyUserStatus(param models.ParamModifyUserStatus) (string,error) {
	var (
		user *models.Users = &models.Users{}
		opreate,opreatecn string
	)
	user.Status = param.Status
	user.CnName = param.EnName
	switch param.Status {
	case 1:
		opreate = "enable"
		opreatecn = "启用"
	case 2:
		opreate = "disable"
		opreatecn = "禁用"
	}
	err := user.UpdateByEnName()
	if err != nil {
		middleware.Logger.Errorf(fmt.Printf("%v user %v password failed: %v\n",opreate,user.EnName,err))
		return fmt.Sprintf("%v用户失败"),err
	}
	
	middleware.Logger.Errorf(fmt.Printf("%v user %v password success",opreate,user.EnName,err))
	return fmt.Sprintf("%v用户成功"),err
}

//删除用户，软删除
func DeleteUser(param models.ParamUserEnName) (string,error) {
	var (
		user *models.Users = &models.Users{}
	)
	user.EnName = param.EnName
	user.Status = 3
	err := user.UpdateByEnName()
	if err != nil {
		middleware.Logger.Errorf(fmt.Printf("delete user %v failed: %v\n",user.EnName,err))
		return fmt.Sprintf("删除用户失败"),err
	}
	
	middleware.Logger.Errorf(fmt.Printf("delete user %v success",user.EnName,err))
	return fmt.Sprintf("删除用户成功"),err
}

//获取用户列表
func UserList(param models.ParamUserList) (map[string]interface{},string,error) {
	var (
		users map[string]interface{} = make(map[string]interface{},0)
		user *models.Users = &models.Users{}
		
	)
	user.Mobile = param.Mobile
	user.CnName = param.CnName
	user.EnName = param.EnName
	user.Status = param.Status
	count,us,err := user.List()
	if err != nil {
		middleware.Logger.Errorf(fmt.Printf("get user list failed: %v\n",err))
		return nil,fmt.Sprintf("获取用户列表失败"),err
	}
	users["count"] = count
	users["users"] = us
	
	return users,fmt.Sprintf("获取用户列表成功"),nil
}

func ModifyUserInfo(param models.ParamModifyUserInfo) (string,error) {
	var user *models.Users = &models.Users{}
	user.Mobile = param.Mobile
	user.CnName = param.CnName
	user.EnName = param.EnName
	user.Email = param.Email
	err := user.UpdateByEnName()
	if err != nil {
		middleware.Logger.Errorf(fmt.Printf("modify user %v info failed: %v\n",user.EnName,err))
		return fmt.Sprintf("修改个人信息失败"),err
	}
	
	middleware.Logger.Errorf(fmt.Printf("modify user %v info success",user.EnName,err))
	return fmt.Sprintf("修改个人信息成功"),err
}