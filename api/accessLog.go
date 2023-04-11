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

func GetAccessDomainList(c *gin.Context) {

	var r params.AccessDomainLogGetParam

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	data, total := service.UserAccessDomainLogSrv.GetList(r)
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"total": total,
	})
}

func GetAccessUserList(c *gin.Context) {

	var r params.AccessUserLogGetParam

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	data, total := service.UserAccessSrv.GetList(r)
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"total": total,
	})
}

func GetAccessUserDetail(c *gin.Context) {

	var r params.AccessUserLogGetParam

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	data, total := service.UserAccessDetailSrv.GetList(r)
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"total": total,
	})
}
