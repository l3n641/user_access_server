package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_accerss_server/api/params"
	"user_accerss_server/internal/service"
)

func AddAccessLog(c *gin.Context) {

	var r params.AccessLogPostParam

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}

	service.UserAccessSrv.AddRecord(r)
	c.JSON(http.StatusOK, gin.H{})
}
