package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"user_accerss_server/api"
	"user_accerss_server/api/middleware"
	"user_accerss_server/internal/database/mongoDb"
)
import "github.com/spf13/viper"

func init() {
	initConfig()
	initFileLog()
	mongoDb.InitDb()
}

func main() {
	router := gin.Default()
	debug := viper.GetString("app.debug")
	if debug != "" {
		gin.SetMode(gin.DebugMode)
	}
	router.Use(middleware.Cors)

	apiGroup := router.Group("/api")

	apiGroup.POST("/session", api.Session)
	apiGroup.DELETE("/session", api.Logout)
	apiGroup.GET("/user_info", api.GetUserInfo)

	apiGroup.POST("/user_access_log", api.AddAccessLog)
	apiGroup.GET("/user_access_log", middleware.Authorization, api.GetAccessList)

	httpPort := viper.GetString("app.httpPort")
	http.ListenAndServe(":"+httpPort, router)
}

func initConfig() {
	viper.SetConfigName("configs/app")
	viper.AddConfigPath(".") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

// 初始化日志
func initFileLog() {
	gin.DisableConsoleColor()
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}
