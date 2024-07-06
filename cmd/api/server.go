package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"tianhe/config"
	"tianhe/pkg"
	"tianhe/middleware"
	"tianhe/routers"

	
	"github.com/spf13/cobra"
)

var (
	conf   string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start Http server",
		Example: "oncall  server -c config/gin-demo.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&conf, "config", "c", "config/conf.yaml", "Start server with provided configuration file")
}

func usage() {
	usageStr := `starting http server`
	log.Printf("%s\n", usageStr)
}

func setup() {
	// 1. 读取配置
	config.InitConfig()
	// 2.初始化log
	middleware.InitLog()
	// 3. 初始化数据链接
	middleware.InitDB()
	//4、初始化sms客户端
	pkg.InitSms()
	// 4. http客户端初始化
	//httpclient.Init()
	// 5. 启动异步任务队列
	//go task.Start()
	//6.启动定时任务客户端
	//go service.InitCron()
	//7.初始化任务
	//go service.OncallTask()
}

func run() error {
	r := router.InitRouter()
	
	//停服之前关闭数据库连接
	defer func() {
		err := middleware.Sql.Close()
		if err != nil {
			log.Printf("close mysql connection failed:%v",err)
		}
	}()
	
	
	//启动http服务
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d",config.Conf.Server.Host,config.Conf.Server.Port),
		Handler: r,
	}
	
	go func() {
		// 服务连接
		if config.Conf.Server.IsHttps {
			if err := srv.ListenAndServeTLS(config.Conf.Server.Ssl.Pem, config.Conf.Server.Ssl.Key); err != nil && err != http.ErrServerClosed {
				log.Printf("listen: %s\n", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Printf("listen: %s\n", err)
			}
		}
	}()
	/*
	log.Printf("%s Server Run http://%s:%s/ \r\n",
		tools.GetCurrntTimeStr(),
		config.Conf.Server.Host,
		config.Conf.Server.Port)
	
	log.Printf("%s Swagger URL http://%s:%s/swagger/index.html \r\n",
		tools.GetCurrntTimeStr(),
		config.Conf.Server.Host,
		config.Conf.Server.Port)
	log.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrntTimeStr())
	*/
	
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt,syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//log.Printf("%s Shutdown Server ... \r\n", tools.GetCurrntTimeStr())
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown:", err)
	}
	log.Printf("Server exiting")
	return nil
}

