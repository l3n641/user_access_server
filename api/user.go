package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	name := viper.GetString("app.admin_name")
	avatar := viper.GetString("app.admin_avatar")
	c.JSON(http.StatusOK, gin.H{
		"name":   name,
		"avatar": avatar,
		"roles":  "admin",
	})

}
