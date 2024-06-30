package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_NOT_FOUND_CODE = 404
	HTTP_NO_LOGIN = 1001
	HTTP_TOKEN_INVALID = 1002
	HTTP_FAIL_CODE = 1003
	HTTP_SUCCESS_CODE = 200
	HTTP_USER_IS_DISABLE = 1004
	HTTP_MOBILE_INVALID = 1005
	HTTP_LOGIN_TYPE_INVALID = 1006
	HTTP_CHECK_FAILED = 1007
	
)

type Context struct {
	Ctx *gin.Context
}

type Response struct {
	Errno  int         `json:"errno"`
	Data   interface{} `json:"data"`
	ErrMsg string      `json:"errmsg"`
}

func (c *Context) Response(errno int, errmsg string, data interface{}) {

	c.Ctx.JSON(http.StatusOK, Response{
		Errno:  errno,
		Data:   data,
		ErrMsg: errmsg,
	})
}
