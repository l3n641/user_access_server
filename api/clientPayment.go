package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_accerss_server/api/params"
	"user_accerss_server/internal/service"
)

func ClientPayment(c *gin.Context) {
	var r params.ClientInfoParams

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}

	service.ClientPaymentSrv.Add(r)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"html": "",
	})
}
