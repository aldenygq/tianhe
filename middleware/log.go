package middleware

import (
	"bufio"
	"tianhe/config"
	"os"
	"path"
	"time"
	
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/aldenygq/toolkits"
	"log"
)

var (
	apiLogger *logrus.Logger
	busLogger *logrus.Logger
)


const (
	LOG_PATH       = "log"
	LOG_LEVEL_INFO = "info"
	TIME_FORMAT    = "2006-01-02 15:04:05"
)

func InitApiLog() gin.HandlerFunc {
	apiLogger = logrus.New()
	logFilePath := LOG_PATH
	if !toolkits.IsDir(logFilePath) {
		if !toolkits.CreateDir(logFilePath) {
			log.Printf("create log dir fialed")
			os.Exit(-1)
		}
	}
	apilogFileName := config.Conf.Log.Logfile.Api
	apifileName := path.Join(logFilePath, apilogFileName)
	apiLogger.SetReportCaller(true)
	log.Printf("api log file:%v",apifileName)
	apifile, err := os.OpenFile(apifileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("write log file failed:%v\n", err)
		os.Exit(-1)
	}
	apiLogger.SetOutput(bufio.NewWriter(apifile))
	if config.Conf.Log.Loglevel == LOG_LEVEL_INFO {
		// 设置日志级别
		apiLogger.SetLevel(logrus.InfoLevel)
	} else {
		apiLogger.SetLevel(logrus.DebugLevel)
	}

	// 设置 rotatelogs
	apilogWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		apifileName+".%Y%m%d%H",

		// 生成软链，指向最新日志文件Loglevel
		rotatelogs.WithLinkName(apifileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(config.Conf.Log.Logmaxage*24*time.Hour),

		// 设置日志切割时间间隔(1小时)
		rotatelogs.WithRotationTime(time.Hour),
	)

	apiwriteMap := lfshook.WriterMap{
		logrus.InfoLevel:  apilogWriter,
		logrus.FatalLevel: apilogWriter,
		logrus.DebugLevel: apilogWriter,
		logrus.WarnLevel:  apilogWriter,
		logrus.ErrorLevel: apilogWriter,
		logrus.PanicLevel: apilogWriter,
	}
	apilfHook := lfshook.NewHook(apiwriteMap, &nested.Formatter{
		HideKeys:        true,
		NoFieldsColors:  false,
		CallerFirst:     false,
		TrimMessages:    true,
		TimestampFormat: TIME_FORMAT,
	})
	// 新增 Hook
	apiLogger.AddHook(apilfHook)
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		//开始时间
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		method := c.Request.Method
		statusCode := c.Writer.Status()
		ip := c.ClientIP()
        requestId,_:= c.Get("X-Request-Id")
		if config.Conf.Log.Loglevel == LOG_LEVEL_INFO {
			apiLogger.WithFields(logrus.Fields{
				//requestId
				"requestid": requestId,
				//客户端ip
				"clientIp": ip,
				//状态码
				"statusCode": statusCode,
				//接口请求方法
				"reqMethod": method,
				//请求接口
				"reqUri": uri,
				//请求耗时
				"latencyTime": latencyTime,
			}).Info()
		} else {
			now := time.Now().Format(TIME_FORMAT)
			apiLogger.Infof("%s | %s | %3d | %13v | %15s | %s  %s",
				requestId,
				now,
				statusCode,
				latencyTime,
				ip,
				method,
				uri,
			)
		}
	}
}

func InitLog() {
	busLogger = logrus.New()

	logFilePath := LOG_PATH
	if !toolkits.IsDir(logFilePath) {
		if !toolkits.CreateDir(logFilePath) {
			log.Printf("create log dir fialed")
			os.Exit(-1)
		}
	}
	
	buslogFileName := config.Conf.Log.Logfile.Bus
	// 日志文件
	busfileName := path.Join(logFilePath, buslogFileName)
	log.Printf("bus file name:%v\n", buslogFileName)
	//输出日志中添加文件名和方法信息
	busLogger.SetReportCaller(true)
	// 写入文件
	busfile, err := os.OpenFile(busfileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("write log file failed:%v\n", err)
		os.Exit(-1)
	}
	// 设置输出
	busLogger.SetOutput(bufio.NewWriter(busfile))

	if config.Conf.Log.Loglevel == LOG_LEVEL_INFO {
		// 设置日志级别
		busLogger.SetLevel(logrus.InfoLevel)
	} else {
		busLogger.SetLevel(logrus.DebugLevel)
	}
	buslogWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		busfileName+".%Y%m%d%H",

		// 生成软链，指向最新日志文件Loglevel
		rotatelogs.WithLinkName(busfileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(config.Conf.Log.Logmaxage*24*time.Hour),

		// 设置日志切割时间间隔(1小时)
		rotatelogs.WithRotationTime(time.Hour),
	)

	buswriteMap := lfshook.WriterMap{
		logrus.InfoLevel:  buslogWriter,
		logrus.FatalLevel: buslogWriter,
		logrus.DebugLevel: buslogWriter,
		logrus.WarnLevel:  buslogWriter,
		logrus.ErrorLevel: buslogWriter,
		logrus.PanicLevel: buslogWriter,
	}
	buslfHook := lfshook.NewHook(buswriteMap, &nested.Formatter{
		HideKeys:        true,
		NoFieldsColors:  false,
		CallerFirst:     false,
		TrimMessages:    true,
		TimestampFormat: TIME_FORMAT,
	})

	busLogger.AddHook(buslfHook)
}

func getRequestId(c *gin.Context) (value any) {
	requestid,_ := c.Get("X-Request-Id")
	return requestid
}
 
func  LogInfof(c *gin.Context,msg string) {
	busLogger.WithField("request_id",getRequestId(c)).Infof(msg)
}

func  LogErrorf(c *gin.Context,msg string) {
	busLogger.WithField("request_id",getRequestId(c)).Errorf(msg)
}