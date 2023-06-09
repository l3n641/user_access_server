package applePay

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/apple_pay_bill", CreateApplePayBill)
}
