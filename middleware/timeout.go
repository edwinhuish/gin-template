package middleware

import (
	"net/http"
	"time"

	"gin-template/global"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/timeout"
)

func timeOutResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "")
}

func Timeout() gin.HandlerFunc {

	return timeout.New(
		timeout.WithTimeout(time.Duration(global.CONFIG.System.Timeout)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeOutResponse),
	)
}
