package service

import (
	//"os"

	"context"

	"github.com/robfig/cron/v3"
	//"tianhe/middleware"
	//"tianhe/middleware"
	"tianhe/middleware"
	"tianhe/models"
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
	rules,err := rule.EnabledRule()
	if err != nil {
		middleware.CronLog(ctx).Errorf("get enable rule list failed:%v\n",err)
		return 
	}
	
	if len(rules) >= 0 {
		return 
	}

	for _,rule := range rules {
		

	}
}
