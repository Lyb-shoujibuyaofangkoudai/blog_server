package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"` // 0成功
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	SUCCESS = 0
	ERROR   = -1
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

// OkWith 只返回最基本的成功
func OkWith(c *gin.Context) {
	Result(SUCCESS, map[string]string{}, "success", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]string{}, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]string{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrMap[code]
	if ok {
		Result(int(code), map[string]string{}, msg, c)
		return
	}
	Result(ERROR, map[string]string{}, "未匹配到对应的error code", c)
}

type FileUpload struct {
	FileName  string
	Url       string
	IsSuccess bool
	ErrMsg    string
}
