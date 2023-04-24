package service

import (
	"github.com/spf13/viper"
	"user_accerss_server/api/admin/params"
)

type UserService struct {
}

func (p *UserService) Login(data params.LoginData) bool {
	name := viper.GetString("app.admin_name")
	password := viper.GetString("app.admin_password")
	if data.Password == password && name == data.UserName {
		return true
	}
	return false
}
