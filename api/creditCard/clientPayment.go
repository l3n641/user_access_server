package creditCard

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_accerss_server/api/creditCard/params"
	"user_accerss_server/internal/service"
)

func ClientPayment(c *gin.Context) {
	var r params.ClientInfoParams

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	service.ClientPaymentSrv.Add(r)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"html": "",
	})
}

func GetClientPaymentList(c *gin.Context) {

	var r params.ClientPaymentListParam

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	data, total := service.ClientPaymentSrv.GetList(r)
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"total": total,
	})
}
