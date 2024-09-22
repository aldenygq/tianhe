package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UEnName string
	jwt.StandardClaims
}
//产生token
func(cc *CustomClaims ) MakeToken() (string,error) {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	return token.SignedString([]byte(SECRETKEY))
}
//解析token
func ParseToken(tokenString string) (*CustomClaims,error)  {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims,nil
	} else {
		return nil,err
	}
}


//加入到黑名单
func DelToken(uname string ) error {
	err := RedisClient.Del(uname).Err()
	if err != nil {
		return err
	}
	return nil
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//LogInfo(c).Infof(fmt.Sprintf("access token:%v\n",ACCESS_TOKEN))
		auth := c.Request.Header.Get(ACCESS_TOKEN)
		if len(auth) == 0 {
			c.Abort()
			Log(c).Errorf(fmt.Sprintf("token %v invalid\n",auth))
			c.JSON(200, gin.H{
				"errno":"401",
				"errmsg": fmt.Sprintf("token invalid",auth),
				"requestid":GetRequestId(c),
				"data":"",
			})
			return
		}
		ret,err:= ParseToken(auth)
		if err != nil {
			c.Abort()
			Log(c).Errorf(fmt.Sprintf("user login info expired:%v",err))
			c.JSON(200, gin.H{
				"errno":"401",
				"errmsg": fmt.Sprintf("user login info expired:%v",err),
				"requestid":GetRequestId(c),
				"data":"",
			})
			return 
		}
		//LogInfo(c).Infof(fmt.Sprintf("user name:%v",ret.UEnName))
		val,err := RedisClient.Get(ret.UEnName).Result()
		//LogInfo(c).Infof(fmt.Sprintf("val:%v",val))
		if err != nil || val != auth{
			c.Abort()
			Log(c).Errorf(fmt.Sprintf("user %v not login",ret.UEnName))
			c.JSON(200, gin.H{
				"errno":"401",
				"errmsg": fmt.Sprintf("user %v not login",ret.UEnName),
				"requestid":GetRequestId(c),
				"data":"",
			})
			return 
		}
		c.Next()
	}
}

func DoLogin(c *gin.Context,enname string,expire int64) (string,error) {
	customClaims :=&CustomClaims{
		UEnName:         enname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expire)*time.Second).Unix(), // 过期时间，必须设置
			Issuer:    "alden",
		},
	}
	accessToken, err :=customClaims.MakeToken()
	if err != nil {
		Log(c).Errorf(fmt.Sprintf("create access token failed:%v",err))
		return "",err
	}
	Log(c).Infof("expire:%v\n",expire)
	//存储登陆状态
	err = RedisClient.Set(enname,accessToken,time.Duration(expire)*time.Second).Err()
	if err != nil {
		Log(c).Errorf(fmt.Sprintf("save access token to redis failed:%v",accessToken))
		return "",err
	}
	
	return accessToken,nil 
}

