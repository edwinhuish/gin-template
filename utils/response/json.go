package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 描述
	Data    any    `json:"data"`    // 返回数据
}

func JSON(c *gin.Context, code int, message string, data any) {
	c.JSON(http.StatusOK, JsonResponse{
		code,
		message,
		data,
	})
}

func Ok(c *gin.Context, message string) {
	JSON(c, 0, "ok", nil)
}

func OkWithMessage(c *gin.Context, message string) {
	JSON(c, 0, message, nil)
}

func OkWithData(c *gin.Context, data any) {
	JSON(c, 0, "ok", data)
}

func OkWithDetail(c *gin.Context, message string, data any) {
	JSON(c, 0, message, data)
}

func Fail(c *gin.Context, message string) {
	JSON(c, 1, "error", nil)
}

func FailWithMessage(c *gin.Context, message string) {
	JSON(c, 1, message, nil)
}

func FailWithData(c *gin.Context, data any) {
	JSON(c, 1, "error", data)
}

func FailWithDetail(c *gin.Context, message string, data any) {
	JSON(c, 1, message, data)
}
