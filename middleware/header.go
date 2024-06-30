package middleware

import (
	//"encoding/base64"
	"fmt"
	"net/http"
	//"net/url"
	//"strconv"
	"time"
	//"oncall/models"
	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
)

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
func NoRoute(c *gin.Context) {
	ctx := Context{Ctx: c}
	path := c.Request.URL.Path
	method := c.Request.Method
	
	ctx.Response(HTTP_NOT_FOUND_CODE, fmt.Sprintf("%s %s not found", method, path), nil)
}

//func GetParam(c *gin.Context,key string) string {
//	val := c.GetHeader(key)
//	if val != ""{
//		return val,true
//	}
	//val,err := c.Cookie(key)
	//if err != nil {
	//	return "",false
	//}
//	return "",true
//}


func Auth(c *gin.Context){
	//ctx := Context{Ctx: c}
	//var user *models.Users = &models.Users{}
	accessToken := c.GetHeader(ACCESS_TOKEN)
	if  accessToken == "" {
		//c.Abort()//组织调起其他函数
		c.Redirect(http.StatusMovedPermanently, "/tianhe/auth/v1/login")
		//ctx.Response(HTTP_NO_LOGIN, fmt.Sprintf("用户未登录"), nil)
		//return
	}
	ret,err:= ParseToken(accessToken)
	if err != nil {
		//
		//c.Abort()
		//ctx.Response(HTTP_TOKEN_INVALID, fmt.Sprintf(err.Error()), nil)
		//return
		c.Redirect(http.StatusMovedPermanently, "/tianhe/auth/v1/login")
	}
	//不存在代表未登录
	has := CheckToken(accessToken)
	if !has {
		//c.Abort()//组织调起其他函数
		//ctx.Response(HTTP_NO_LOGIN, fmt.Sprintf(err.Error()), nil)
		//return
		c.Redirect(http.StatusMovedPermanently, "/tianhe/auth/v1/login")
	}
	c.Set("User_En_Name",ret.UEnName)
	c.Next()
	return
}

func DoLogin(c *gin.Context,enname string)  error{
	customClaims :=&CustomClaims{
		UEnName:         enname,
		//Email: user.Email
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(MAXAGE)*time.Second).Unix(), // 过期时间，必须设置
		},
	}
	accessToken, err :=customClaims.MakeToken()
	if err != nil {
		return err
	}
	
	refreshClaims :=&CustomClaims{
		UEnName:        enname,
		//Email: user.Email
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(MAXAGE+1800)*time.Second).Unix(), // 过期时间，必须设置
		},
	}
	refreshToken, err :=refreshClaims.MakeToken()
	if err != nil {
		return err
	}
	
	//存储登陆状态
	//authUser := base64.StdEncoding.EncodeToString([]byte(user.EnName))
	//设置登陆状态
	err = RedisClient.Set(accessToken,enname,time.Duration(MAXAGE)*time.Second).Err()
	if err != nil {
		return err
	}
	
	c.Header(ACCESS_TOKEN,accessToken)
	c.Header(REFRESH_TOKEN,refreshToken)
	
	return nil
}
//判断是否https
func IsHttps(c *gin.Context) bool {
	if c.GetHeader("X-Forwarded-Proto") =="https" || c.Request.TLS!=nil{
		return true
	}
	return false
}


