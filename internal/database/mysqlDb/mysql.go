package mysqlDb

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func NewDatabase() {
	dsn := viper.GetString("database_mysql.dsn")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDatabase() (c *gorm.DB) {
	return db
}
