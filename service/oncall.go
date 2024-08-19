package service

import (
	"encoding/json"
	"errors"
	"fmt"

	//"strings"
	"time"

	"tianhe/middleware"
	"tianhe/models"

	"github.com/gin-gonic/gin"
)

//默认信息
func DefaultInfo(c *gin.Context,param models.ParamDefaultInfo) (map[string]interface{},string,error) {
	var (
		data map[string]interface{} = make(map[string]interface{},0)
		err error
	)
	switch param.Operate {
	case "add_oncall_rule":
		data,err = GetAddOncallInfoDefaultInfo(c)
	}
	if err != nil {
		middleware.LogErr(c).Errorf("get default info by %v failed:%v\n",param.Operate,err)
		return nil,fmt.Sprintf("获取默认信息失败,失败原因：%v",err),err
	}
	
	return data,fmt.Sprintf("获取默认信息成功"),nil
}
//获取新增值班规则默认信息
func GetAddOncallInfoDefaultInfo(c *gin.Context) (map[string]interface{},error) {
	var (
		user *models.Users = &models.Users{}
		data map[string]interface{} = make(map[string]interface{},0)
	)
	_,list,err := user.List()
	if err != nil {
		middleware.LogErr(c).Errorf("get all users failed:%v\n",err)
		return nil,err
		
	}
	data["oncall_types"] = []string{"按天值班","按周值班","按月值班","自定义值班"}
	data["users"] = list
	data["is_skip_weekend"] = []string{"是","否"}
	data["is_temporary_oncall"] = []string{"是","否"}
	return data,nil
}

//新增值班规则
func AddOncall(c *gin.Context,param models.ParamAddOncallRule) (string,error) {
	var (
		//msg string
		err error
		oncall *models.OncallRule = &models.OncallRule{}
	)

	//开始日期不能小于当前日期
	middleware.LogErr(c).Infof("current day:%v\n",getDay())
	middleware.LogErr(c).Infof("start day:%v\n",param.StartDay)
	num := CompareTwoDay(getDay(),param.StartDay)
	if num != 0 {
		middleware.LogErr(c).Errorf("start day must greater than current day")
		return fmt.Sprintf("开始日期应大于当前日期"),errors.New("start day must greater than current day")
	}
	/*
	//自定义类型轮转天数不能小于0
	if param.OncallCycleType == "custom" {
		if param.PerRotationDays <= 0 {
			middleware.Logger.Errorf("custom type rotation days must be greater than 0")
			return fmt.Sprintf("自定义类型轮转天数必须大于0"),errors.New("custom type rotation days must be greater than 0")
		}
	}
	*/
	//临时值班开启
	if param.IsTemporaryOncall == 2 {
		if param.TemporaryOncallInfo == nil {
			middleware.LogErr(c).Errorf("temporary oncall info can not nil")
			return fmt.Sprintf("临时值班信息不能为空"),errors.New("temporary oncall info can not nil")
		}
	}
	
	oncall.IsSkipWeekend = param.IsSkipWeekend
	oncall.CnTitle = param.CnTitle
	oncall.IsTemporaryOncall = param.IsTemporaryOncall
	//oncall.PerRotationDays = param.PerRotationDays
	oncall.OncallCycleType = param.OncallCycleType
	data,_ := json.Marshal(param.OncallPeople)
	//oncall.OncallPeopleInfos = param.OncallPeople
	oncall.OncallPeopleInfos = string(data)
	oncall.StartDay = param.StartDay
	oncall.SubscribeGroups = param.SubscribeGroups
	oncall.EnTitle = param.EnTitle
	oncall.TemporaryOncallInfo = param.TemporaryOncallInfo
	oncall.Status = param.Status
	oncall.SubscribeNotifyInfo = param.SubscribeNotifyInfo
	oncall.RotationNum = param.RotationNum
	oncall.Creator = param.Creator
	oncall.CreateTime = time.Now().Unix()
	err = oncall.Create()
	if err != nil {
		middleware.LogErr(c).Errorf("add oncall rule failed:%v\n",err)
		return fmt.Sprintf("创建值班规则失败,失败原因:%v\n",err),err
	}
	return fmt.Sprintf("创建值班规则成功"),nil
}
/*
//修改值班规则
func ModifyOncall(param *models.ParamModifyOncallRule) (string,error) {
	var (
		//msg string
		err error
		oncall *models.OncallRule = &models.OncallRule{}
	)
	
	//开始日期不能小于当前日期
	num := CompareTwoDay(getDay(),param.StartDay)
	if num != 0 {
		middleware.Logger.Errorf("start day must greater than current day")
		return fmt.Sprintf("开始日期应大于当前日期"),errors.New("start day must greater than current day")
	}
	
	//自定义类型轮转天数不能小于0
	if param.OncallCycleType == "custom" {
		if param.PerRotationDays <= 0 {
			middleware.Logger.Errorf("custom type rotation days must be greater than 0")
			return fmt.Sprintf("自定义类型轮转天数必须大于0"),errors.New("custom type rotation days must be greater than 0")
		}
	}
	//临时值班开启
	if param.IsTemporaryOncall == 0 {
		if param.TemporaryOncallInfo == nil {
			middleware.Logger.Errorf("temporary oncall info can not nil")
			return fmt.Sprintf("临时值班信息不能为空"),errors.New("temporary oncall info can not nil")
		}
	}
	
	oncall.Id = param.Id
	oncall.IsSkipWeekend = param.IsSkipWeekend
	oncall.CnTitle = param.EnTitle
	oncall.IsTemporaryOncall = param.IsTemporaryOncall
	oncall.PerRotationDays = param.PerRotationDays
	oncall.OncallCycleType = param.OncallCycleType
	oncall.OncallPeopleInfos = param.OncallPeople
	oncall.StartDay = param.StartDay
	oncall.SubscribeGroups = param.SubscribeGroups
	oncall.EnTitle = param.EnTitle
	oncall.TemporaryOncallInfo = param.TemporaryOncallInfo
	oncall.Status = param.Status
	oncall.SubscribeNotifyInfo = param.SubscribeNotifyInfo
	oncall.RotationNum = param.RotationNum
	oncall.Creator = param.Updator
	oncall.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	err = oncall.Modify()
	if err != nil {
		middleware.Logger.Errorf("add oncall rule failed:%v\n",err)
		return fmt.Sprintf("修改值班规则失败,失败原因:%v\n",err),err
	}
	return fmt.Sprintf("修改值班规则成功"),nil
}

//值班规则列表
func OncallRules(param *models.ParamSearch) (*models.RespOncallRules,string,error) {
	var (
		ruleinfo *models.RespOncallRules = &models.RespOncallRules{}
		err error
		rule *models.OncallRule = &models.OncallRule{}
	)
	//rule.Id = param.Id
	rule.CnTitle = param.CnTitle
	rule.EnTitle = param.EnTitle
	rule.Creator = param.Creator
	total,rules,err := rule.List()
	if err != nil {
		middleware.Logger.Errorf("get oncall rule list failed:%v\n",err)
		return nil,fmt.Sprintf("获取值班规则列表失败,失败原因:%v\n",err),err
	}
	ruleinfo.Rules = rules
	ruleinfo.Count = total
	return ruleinfo,fmt.Sprintf("获取值班规则列表成功"),nil
}


func CurrrentDutyInfos(param models.ParamDutyPerson) ([]*models.RespDutyPerson,string,error){
	var (
		//ruleinfos []*models.RespDutyPerson = make([]*models.RespDutyPerson,0)
		err error
		rule *models.OncallRule = &models.OncallRule{}
		duty *models.CurrentDutyInfo = &models.CurrentDutyInfo{}
	)

	if param.RuleName != "" {
		switch IsContainChinese(param.RuleName) {
		case true:
			rule.CnTitle = param.RuleName
		case false:
			rule.EnTitle = param.RuleName
		default:
			middleware.Logger.Errorf("rule name invalid")
		}
		err = rule.Get()
		if err != nil {
			middleware.Logger.Errorf("get current duty info failed:%v\n",err)
			return nil,fmt.Sprintf("获取当前值班信息失败,失败原因:%v\n",err),err
		}
		duty.RuleId = rule.Id
	}
	if len(param.User) > 0 {
		duty.User = strings.Join(param.User,",")
	}
	
	dutys,err := duty.List()
	if err != nil {
		middleware.Logger.Errorf("get current duty info failed:%v\n",err)
		return nil,fmt.Sprintf("获取当前值班信息失败,失败原因:%v\n",err),err
	}
	return dutys,fmt.Sprintf("获取当前值班信息成功"),nil
}
*/
