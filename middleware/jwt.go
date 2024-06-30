package middleware

import (
	"fmt"
	
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UEnName string
	//EmployeeNum string
	//Email string
	//UCnName string
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
func DelToken(token string ) error {
	// 从池里获取连接
	_,err := RedisClient.Get(token).Result()
	if err != nil {
		return err
	}
	// 用完后将连接放回连接池
	defer RedisClient.Close()
	err = RedisClient.Del(token).Err()
	if err != nil {
		return err
	}
	return nil
}
//检查token是否存在
func CheckToken(token string) bool  {
	 err := RedisClient.Get(token).Err()
	 if err != nil {
	 	return false
	 }
	// 用完后将连接放回连接池
	defer RedisClient.Close()
	return true
}
