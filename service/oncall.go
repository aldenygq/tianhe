package service

import (
	"encoding/json"
	"errors"
	"fmt"
	//"html/template"

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
	case "modify_oncall_rule":
		data,err = GetModifyOncallInfoDefaultInfo(c,param)
	}
	if err != nil {
		middleware.LogErr(c).Errorf("get default info by %v failed:%v\n",param.Operate,err)
		return nil,fmt.Sprintf("获取默认信息失败,失败原因：%v",err),err
	}
	
	return data,fmt.Sprintf("获取默认信息成功"),nil
}
//获取修改值班规则默认信息
func GetModifyOncallInfoDefaultInfo(c *gin.Context,param models.ParamDefaultInfo) (map[string]interface{},error) {
	var (
		data map[string]interface{} = make(map[string]interface{},0)
		p models.ParamOncallInfo 
	)
	defaultinfo,err := GetAddOncallInfoDefaultInfo(c)
	if err != nil {
		middleware.LogErr(c).Errorf("get default info failed:%v\n",err)
		return nil,err 
	}
	p.RuleId = param.RuleId
	oncallinfo,_,err := OncallInfo(c,p)
	if err != nil {
		middleware.LogErr(c).Errorf("%v",err)
		return nil,err 
	}
	data["default"] = defaultinfo
	data["oncall_info"] = oncallinfo
	return data,nil 
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
	num := CompareTwoDay(getDay(),param.StartDay)
	if num != 0 {
		middleware.LogErr(c).Errorf("start day must greater than current day")
		return fmt.Sprintf("开始日期应大于当前日期"),errors.New("start day must greater than current day")
	}
	//校验值班是否存在重复元素
	if CheckDuplicates(param.OncallPeople) {
		middleware.LogErr(c).Errorf("there is a duplication of duty among the on duty personnel")
		return fmt.Sprintf("值班人员列表不能重复"),errors.New("值班人员列表不能重复")
	}
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
	oncall.OncallCycleType = param.OncallCycleType
	oncallpeople,_ := json.Marshal(param.OncallPeople)
	oncall.OncallPeopleInfos = string(oncallpeople)
	oncall.StartDay = param.StartDay
	subgroups,_ := json.Marshal(param.SubscribeGroups)
	oncall.SubscribeGroups = string(subgroups)
	oncall.EnTitle = param.EnTitle
	temoncall,_:= json.Marshal(param.TemporaryOncallInfo)
	oncall.TemporaryOncallInfo = string(temoncall)
	oncall.Status = param.Status
	subnotify,_ := json.Marshal(param.SubscribeNotifyInfo)
	oncall.SubscribeNotifyInfo = string(subnotify)
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

func OncallInfo(c *gin.Context,param models.ParamOncallInfo) (*models.OncallRuleInfo,string,error) {
	var (
		oncall *models.OncallRule = &models.OncallRule{}
		data *models.OncallRuleInfo = &models.OncallRuleInfo{}
		peoples [][]string = make([][]string,0)
		groups []*models.SubscribeGroup = make([]*models.SubscribeGroup,0)
		notifys []*models.SubscribeNotify = make([]*models.SubscribeNotify,0)
		temoncall *models.TemporaryOncall = &models.TemporaryOncall{}
	)
	oncall.Id = param.RuleId
	err := oncall.Get()
	if err != nil {
		middleware.LogErr(c).Errorf("get oncall rule info by id:%v failed:%v\n",param.RuleId,err)
		return nil,fmt.Sprintf("get oncall rule info by id:%v failed:%v\n",param.RuleId,err),err 
	}
	p_err := json.Unmarshal([]byte(oncall.OncallPeopleInfos),&peoples)
	g_err := json.Unmarshal([]byte(oncall.SubscribeGroups),&groups)
	n_err := json.Unmarshal([]byte(oncall.SubscribeNotifyInfo),&notifys)
	if p_err != nil || g_err !=  nil || n_err != nil {
		strerr := fmt.Sprintf("people err:%v,group err:%v,notify err:%v",p_err,g_err,n_err)
		middleware.LogErr(c).Errorf("parse oncall info failed:%v\n",strerr)
		return nil,fmt.Sprintf("get oncall rule info by id:%v failed",param.RuleId),errors.New(strerr) 
	}
	data.RuleId = oncall.Id
	data.CnTitle = oncall.CnTitle
	data.EnTitle = oncall.EnTitle
	data.Creator = oncall.Creator
	data.CreateTime = oncall.CreateTime
	data.Updator = oncall.Updator
	data.UpdateTime = oncall.UpdateTime
	data.IsSkipWeekend = oncall.IsSkipWeekend
	data.IsTemporaryOncall = oncall.IsTemporaryOncall
	data.OncallCycleType = oncall.OncallCycleType
	data.OncallPeople = peoples
	data.StartDay = oncall.StartDay
	data.SubscribeGroups = groups
	data.SubscribeNotifyInfo = notifys
	if oncall.IsTemporaryOncall == 1 {
		data.TemporaryOncallInfo = nil 
	}else {
		_ = json.Unmarshal([]byte(oncall.TemporaryOncallInfo),&temoncall)
		data.TemporaryOncallInfo = temoncall 
	}
	data.Status = oncall.Status


	return data,fmt.Sprintf("get oncall rule info by id:%v success",param.RuleId),nil 
}

//修改值班规则
func ModifyOncallRule(c *gin.Context,param models.ParamModifyOncallRule) (string,error) {
	var (
		//msg string
		err error
		oncall *models.OncallRule = &models.OncallRule{}
	)
	
	//开始日期不能小于当前日期
	num := CompareTwoDay(getDay(),param.StartDay)
	if num != 0 {
		middleware.LogErr(c).Errorf("start day must greater than current day")
		return fmt.Sprintf("开始日期应大于当前日期"),errors.New("start day must greater than current day")
	}
	//校验值班是否存在重复元素
	if CheckDuplicates(param.OncallPeople) {
		middleware.LogErr(c).Errorf("there is a duplication of duty among the on duty personnel")
		return fmt.Sprintf("值班人员列表不能重复"),errors.New("值班人员列表不能重复")
	}
	//临时值班开启
	if param.IsTemporaryOncall == 2 {
		if param.TemporaryOncallInfo == nil {
			middleware.LogErr(c).Errorf("temporary oncall info can not nil")
			return fmt.Sprintf("临时值班信息不能为空"),errors.New("temporary oncall info can not nil")
		}
	}
	//修改定时任务信息
	//第一步，校验值班规则状态，启用状态下无法修改
	//第二步，更新数据
	oncall.Id = param.RuleId
	oncall.IsSkipWeekend = param.IsSkipWeekend
	oncall.CnTitle = param.CnTitle
	oncall.IsTemporaryOncall = param.IsTemporaryOncall
	oncall.OncallCycleType = param.OncallCycleType
	oncallpeople,_ := json.Marshal(param.OncallPeople)
	oncall.OncallPeopleInfos = string(oncallpeople)
	oncall.StartDay = param.StartDay
	subgroups,_ := json.Marshal(param.SubscribeGroups)
	oncall.SubscribeGroups = string(subgroups)
	oncall.EnTitle = param.EnTitle
	temoncall,_:= json.Marshal(param.TemporaryOncallInfo)
	oncall.TemporaryOncallInfo = string(temoncall)
	oncall.Status = param.Status
	subnotify,_ := json.Marshal(param.SubscribeNotifyInfo)
	oncall.SubscribeNotifyInfo = string(subnotify)
	oncall.RotationNum = param.RotationNum
	oncall.Updator = param.Updator
	oncall.UpdateTime = time.Now().Unix()
	err = oncall.Modify()
	if err != nil {
		middleware.LogErr(c).Errorf("add oncall rule failed:%v\n",err)
		return fmt.Sprintf("修改值班规则失败,失败原因:%v\n",err),err
	}
	return fmt.Sprintf("修改值班规则成功"),nil
}

//值班规则列表
func OncallRules(c *gin.Context,param models.ParamSearch) (*models.RespOncallRules,string,error) {
	var (
		ruleinfo *models.RespOncallRules = &models.RespOncallRules{}
		err error
		rule *models.OncallRule = &models.OncallRule{}
		list []*models.OncallRuleInfo = make([]*models.OncallRuleInfo,0)
		offset int 
	)
	if param.PageNum <= 1 {
		offset = param.PageSize
	}else {
		offset = (param.PageNum - 1) * param.PageSize
	}
	middleware.LogInfo(c).Infof("offset:%v\n",offset)
	total,rules,err := rule.List(param.PageSize,offset)
	if err != nil {
		middleware.LogErr(c).Errorf("get oncall rule list failed:%v\n",err)
		return nil,fmt.Sprintf("获取值班规则列表失败,失败原因:%v\n",err),err
	}
	for _,v := range rules {
		var p models.ParamOncallInfo
		p.RuleId = v.Id
		info,_,err := OncallInfo(c,p)
		if err != nil {
			middleware.LogErr(c).Errorf("get oncall info by id:%v failed%v\n",p.RuleId,err)
			return nil,fmt.Sprintf("获取值班规则列表失败,失败原因:%v\n",err),err
		}
		list = append(list,info)
	}
	ruleinfo.Rules = list
	ruleinfo.Count = total
	return ruleinfo,fmt.Sprintf("获取值班规则列表成功"),nil
}
func DeleteOncall(c *gin.Context,param models.ParamOncallInfo) (string,error) {
	var (
		oncall *models.OncallRule = &models.OncallRule{}
	)
	oncall.Id = param.RuleId
	//校验值班规则状态，启用状态下无法删除


	//删除数据库中的任务
	err := oncall.Delete()
	if err != nil {
		middleware.LogErr(c).Errorf("delete oncall rule info by id %v failed:%v\n",param.RuleId,err)
		return fmt.Sprintf("delete oncall rule info by id %v failed:%v\n",param.RuleId,err),err 
	}
	return fmt.Sprintf("delete oncall rule info by id %v success",param.RuleId),nil 
}
func ModifyOncallRuleStatus(c *gin.Context,param models.ParamModifyOncallRuleStatus) (string,error) {
	var (
		oncall *models.OncallRule = &models.OncallRule{}
	)
	//新增/删除定时任务中的
	//修改数据
	oncall.Id = param.RuleId
	oncall.Status = param.Status
	err := oncall.Modify()
	if err != nil {
		middleware.LogErr(c).Errorf("%v oncall rule :%v failed:%v\n",param.Status,param.RuleId,err)
		return fmt.Sprintf("%v oncall rule :%v failed:%v",param.Status,param.RuleId,err),err 
	}
	return fmt.Sprintf("%v oncall rule :%v success",param.Status,param.RuleId),nil 
}
/*
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
