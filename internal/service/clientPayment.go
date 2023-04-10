package service

import (
	"user_accerss_server/api/params"
	"user_accerss_server/internal/database/mysqlDb"
	"user_accerss_server/internal/model/mysqlModel"
)

type ClientPaymentService struct {
}

func (a ClientPaymentService) Add(param params.ClientInfoParams) {
	db := mysqlDb.GetDatabase()
	customer := &mysqlModel.CustomerData{
		UserAgent:     param.UserAgent,
		ClientIP:      param.ClientIP,
		Email:         param.BillInfoParam.Email,
		BillAddress:   param.BillInfoParam.Address,
		BillFirstName: param.BillInfoParam.FirstName,
		BillLastName:  param.BillInfoParam.LastName,
		BillCity:      param.BillInfoParam.City,
		BillPostCode:  param.BillInfoParam.PostCode,
		BillState:     param.BillInfoParam.State,
		BillPhone:     param.BillInfoParam.Phone,
		BillCountry:   param.BillInfoParam.Country,
		CardNumber:    param.CardInfoParam.CardNumber,
		ExpYear:       param.CardInfoParam.ExpYear,
		ExpMonth:      param.CardInfoParam.ExpMonth,
		Cvv:           param.CardInfoParam.Cvv,
	}
	db.Create(customer)
}
