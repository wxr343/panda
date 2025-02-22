package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"panda/global"
)

// Response 响应结构体
type Response struct {
	ErrorCode int         `json:"error_code"` // 自定义错误码
	Data      interface{} `json:"data"`       // 数据
	Message   string      `json:"message"`    // 信息
}

// Success 响应成功
func Success(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, Response{
		ErrorCode: 0,
		Data:      data,
		Message:   "Success",
	})
}

// Fail 响应失败，ErrorCode
func Fail(c *gin.Context, errorCode int, msg string) {

	c.JSON(http.StatusOK, Response{
		ErrorCode: errorCode,
		Data:      nil,
		Message:   msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}
