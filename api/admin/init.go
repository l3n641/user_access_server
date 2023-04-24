package admin

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(Group *gin.RouterGroup) {
	Group.POST("/session", Session)
	Group.DELETE("/session", Logout)
	Group.GET("/user_info", GetUserInfo)
}
