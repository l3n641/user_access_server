package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"user_accerss_server/api/admin"
	"user_accerss_server/api/creditCard"
	"user_accerss_server/api/middleware"
	"user_accerss_server/api/userAccess"
	"user_accerss_server/internal/database/mongoDb"
	"user_accerss_server/internal/database/mysqlDb"
	"user_accerss_server/internal/model/mysqlModel"
)
import "github.com/spf13/viper"

func init() {
	initConfig()
	initFileLog()
	initMysql()
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
	admin.RegisterRouter(apiGroup)
	userAccess.RegisterRouter(apiGroup)
	creditCard.RegisterRouter(apiGroup)

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

func initMysql() {
	mysqlDb.NewDatabase()
	db := mysqlDb.GetDatabase()
	db.AutoMigrate(&mysqlModel.CustomerData{})
}
