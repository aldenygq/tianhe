package service

import (
	"os"
	
	"github.com/robfig/cron/v3"
	"oncall/middleware"
	"oncall/models"
)
var Cron *cron.Cron
func InitCron() {
	//nyc, _ := time.LoadLocation("Asia/Shanghai")
	Cron = cron.New()
	Cron.Start()
	select {}
}

func OncallTask() {
	var (
		rule *models.OncallRule = &models.OncallRule{}
		rules []*models.OncallRule = make([]*models.OncallRule,0)
	)
	rules,err := rule.EnabledRule()
	if err != nil {
		middleware.Logger.Errorf("get enabled rule failed:%v\n",err)
		os.Exit(-1)
	}
	
	if len(rules) >= 0 {
	}
}
