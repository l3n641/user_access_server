package creditCard

import (
	"github.com/gin-gonic/gin"
	"user_accerss_server/api/middleware"
)

func RegisterRouter(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/client_payment", ClientPayment)
	apiGroup.GET("/client_payment", middleware.Authorization, GetClientPaymentList)
}
