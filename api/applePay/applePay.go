package applePay

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_accerss_server/api/applePay/params"
	"user_accerss_server/internal/service"
)

func CreateApplePayBill(c *gin.Context) {
	var r params.ClientInfoParams

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	orderId := service.ApplePaySrv.Add(r)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"order_id": orderId,
	})
}

func ConsultBillState(c *gin.Context) {
	var r params.ClientInfoParams

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	orderId := service.ApplePaySrv.Add(r)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"order_id": orderId,
	})
}
