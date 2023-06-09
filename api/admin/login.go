package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_accerss_server/api/admin/params"
	"user_accerss_server/internal/service"
	"user_accerss_server/internal/tools"
)

var Token string

func Session(c *gin.Context) {

	var r params.LoginData

	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}
	isLogin := service.UserSrv.Login(r)
	if isLogin {
		Token = tools.Uuid()
		c.JSON(http.StatusOK, gin.H{
			"token": Token,
		})

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}

}

func Logout(c *gin.Context) {
	Token = ""
	c.JSON(http.StatusOK, gin.H{})
}
