package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"user_accerss_server/api"
)

func Authorization(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "验证失败",
		})
		c.Abort()
		return
	}

	state := strings.HasPrefix(auth, "Bearer ")
	if state == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "验证失败",
		})
		c.Abort()
		return
	}
	result := strings.Split(auth, " ")

	if api.Token == "" || result[1] != api.Token {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "验证失败",
		})
		c.Abort()
		return
	}
	c.Next()
}
