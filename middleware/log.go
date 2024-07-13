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
	gormLogger *logrus.Logger
)


const (
	LOG_PATH       = "log"
	LOG_LEVEL_INFO = "info"
	TIME_FORMAT    = "2006-01-02 15:04:05"
)

func initLog(logger *logrus.Logger,filepath string) *lfshook.LfsHook{
	logFilePath := LOG_PATH
	if !toolkits.IsDir(logFilePath) {
		if !toolkits.CreateDir(logFilePath) {
			log.Printf("create log dir fialed")
			os.Exit(-1)
		}
	}
	fileName := path.Join(logFilePath, filepath)
	logger.SetReportCaller(true)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("write log file failed:%v\n", err)
		os.Exit(-1)
	}	
	logger.SetOutput(bufio.NewWriter(file))
	if config.Conf.Log.Loglevel == LOG_LEVEL_INFO {
		// 设置日志级别
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(logrus.DebugLevel)
	}
		// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d%H",

		// 生成软链，指向最新日志文件Loglevel
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(config.Conf.Log.Logmaxage*24*time.Hour),

		// 设置日志切割时间间隔(1小时)
		rotatelogs.WithRotationTime(time.Hour),
	)
	logwriteMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(logwriteMap, &nested.Formatter{
		HideKeys:        true,
		NoFieldsColors:  false,
		CallerFirst:     false,
		TrimMessages:    true,
		NoColors: false,
		TimestampFormat: TIME_FORMAT,
	})

	return lfHook
}

func InitApiLog() gin.HandlerFunc {
	apiLogger = logrus.New()
	apiLogger.AddHook(initLog(apiLogger,config.Conf.Log.Logfile.Api))
	// 新增 Hook
	//apiLogger.AddHook(apilfHook)
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
	busLogger.AddHook(initLog(busLogger,config.Conf.Log.Logfile.Bus))
}
func InitDbLog() {
	gormLogger = logrus.New()
	gormLogger.AddHook(initLog(gormLogger,config.Conf.Log.Logfile.Sql))
}
 
func LogInfo(c *gin.Context) *logrus.Entry {
	return busLogger.WithField("request_id",GetRequestId(c))
}

func  LogErr(c *gin.Context) *logrus.Entry {
	return busLogger.WithField("request_id",GetRequestId(c))
}