package middleware

import (
	"net/http"
	"tianhe/config"

	"errors"
	"time"
	"strings"
	"strconv"
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")
		
		// Create request id with UUID4
		if requestId == "" {
			requestId = GenerateUuid()
		}
		
		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)
		
		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}

func GetRequestId(c *gin.Context) (value any) {
	requestid,has := c.Get("X-Request-Id")
	if !has {
			u4 := uuid.NewV4()
			requestid = u4.String()
	}
	return requestid
}

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}

//请求方式校验
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodConnect:
			c.Next()
		case http.MethodPost:
			c.Next()
		case http.MethodGet:
			c.Next()
		case http.MethodDelete:
			c.Next()
		case http.MethodHead:
			c.Next()
		case http.MethodPatch:
			c.Next()
		case http.MethodPut:
			c.Next()
		case http.MethodOptions:
			c.Next()
		case http.MethodTrace:
			c.Next()
		default:
			c.AbortWithError(http.StatusMethodNotAllowed, errors.New("The request method is not allowed"))
            return
		}
    }
}

//限速
func LimitHandler() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(float64(config.Conf.Route.Qps), nil)
	lmt.SetMessage("您访问过于频繁，系统安全检查认为恶意攻击。")
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "您的操作太频繁，请稍后再试！",
				"result":  nil,
			})
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

//判断是否https
func IsHttps(c *gin.Context) bool {
	if c.GetHeader("X-Forwarded-Proto") =="https" || c.Request.TLS != nil{
		return true
	}
	return false
}

func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			
			if c.IsAborted() {
				c.Status(200)
			}
			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "CustomError" {
					statusCode, e := strconv.Atoi(p[1])
					if e != nil {
						break
					}
					c.Status(statusCode)
					fmt.Println(
						time.Now().Format("\n 2006-01-02 15:04:05.9999"),
						"[ERROR]",
						c.Request.Method,
						c.Request.URL,
						statusCode,
						c.Request.RequestURI,
						c.ClientIP(),
						p[2],
					)
					c.JSON(http.StatusOK, gin.H{
						"code": statusCode,
						"msg":  p[2],
					})
				}
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}
