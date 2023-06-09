package userAccess

import (
	"github.com/gin-gonic/gin"
	"user_accerss_server/api/middleware"
)

func RegisterRouter(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/user_access_log", AddAccessLog)
	apiGroup.GET("/user_access_domain_log", middleware.Authorization, GetAccessDomainList)
	apiGroup.GET("/user_access_user_log", middleware.Authorization, GetAccessUserList)
	apiGroup.GET("/user_access_user_detail", middleware.Authorization, GetAccessUserDetail)
}
