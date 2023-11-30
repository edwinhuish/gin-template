package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		acceptToken := "21bf0228-e3d2-40b4-9dd7-d71c49a37fb4-25cac0d4-cf21-4bf4-8c2f-0b22aa95e2ea-38f8a25d-89ba"

		token := c.GetHeader("x-proxy-token")

		if token != acceptToken {
			// c.AbortWithStatus(http.StatusNotFound)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无权访问"})
			return
		}

		// 处理请求
		c.Next()
	}
}
