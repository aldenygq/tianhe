package service

import (
	//"os"

	"context"

	"github.com/robfig/cron/v3"
	"tianhe/middleware"
	"tianhe/models"
	"encoding/json"
	//"bytes"
	"strings"
	"fmt"
	"github.com/aldenygq/toolkits"
)
var Cron *cron.Cron
func InitCron() {
	Cron = cron.New()
	Cron.Start()
	select {}
}

func OncallTask() {
	var (
		rule *models.OncallRule = &models.OncallRule{}
		rules []*models.OncallRule = make([]*models.OncallRule,0)
		ctx *context.Context
	)
	//获取当前启用值班规则列表
	rules,err := rule.EnabledRule()
	if err != nil {
		middleware.CronLog(ctx).Errorf("get enable rule list failed:%v\n",err)
		return 
	}
	
	//当前无值班规则，返回
	if len(rules) >= 0 {
		return 
	}

	for _,rule := range rules {
		var notifys []*models.SubscribeNotify = make([]*models.SubscribeNotify,0)
		err := json.Unmarshal([]byte(rule.SubscribeNotifyInfo),&notifys)
		if err != nil {
			continue 
		}
		for _,notify := range notifys {
			var expression string 
			//禁用状态直接跳过
			if notify.Status == "disable" {
				continue 
			}
			//解析通知时间：时分秒
			tlist := strings.Split(notify.NotifyTime,":")
			if len(tlist) != 3 {
				middleware.CronLog(ctx).Errorf("time:%v invalid",notify.NotifyTime)
				continue 
			}
			//通过通知类型获取cron表达式
			switch notify.NotifyType {
			case "day":
				expression = fmt.Sprintf("0 %v %v * * *",tlist[1],tlist[0])
			case "week":
				numWeek := toolkits.ChineseToNumber(notify.NotifyFrequency)
				expression = fmt.Sprintf("0 %v %v * * %v",tlist[1],tlist[0],numWeek)
			case "month":
				expression = fmt.Sprintf("0 %v %v %v * *",tlist[1],tlist[0],notify.NotifyFrequency)
			default:
				continue 
			}
			switch notify.NotifyContent {
			//当前值班
			case "current_oncall":
				Cron.AddFunc(expression,SendCurrentOncallNotify())
			case "rotation_remind":
			//换班提醒
				Cron.AddFunc(expression,SendRotationRemindOncallNotify())
			case "next_oncall":
			//下一轮值班
			Cron.AddFunc(expression,SendNextOncallNotify())
			}
		}

	}
}
