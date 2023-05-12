package params

import "user_accerss_server/api/globalParams"

type ClientInfoParams struct {
	Domain        string         `form:"domain" json:"domain" binding:"required"`
	Site          string         `form:"site" json:"site" binding:"required"`
	UserAgent     string         `form:"user_agent" json:"user_agent" binding:"required"`
	Lang          string         `form:"lang" json:"lang" `
	ClientIP      string         `form:"ip" json:"ip" binding:"required"`
	NotifyUrl     string         `form:"notify_url" json:"notify_url" binding:"required"`
	RedirectUrl   string         `form:"redirect_url" json:"redirect_url" binding:"required"`
	BillInfoParam *BillInfoParam `json:"bill_info" form:"bill_info" binding:"required"`
	ShipInfoParam *ShipInfoParam `json:"ship_info" form:"ship_info" binding:"required"`
	CardInfoParam *CardInfoParam `json:"card_info" form:"card_info" binding:"required"`
	Order         OrderParam     `json:"order" form:"order" binding:"required"`
}

type CardInfoParam struct {
	CardNumber string `form:"card_number" json:"card_number" binding:"required"`
	ExpYear    string `form:"exp_year" json:"exp_year" binding:"required"`
	ExpMonth   string `form:"exp_month" json:"exp_month" binding:"required"`
	Cvv        string `form:"cvv" json:"cvv" binding:"required"`
}

type BillInfoParam struct {
	Address   string `form:"address" json:"address" binding:"required"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	City      string `form:"city" json:"city" binding:"required"`
	State     string `form:"state" json:"state"`
	PostCode  string `form:"post_code" json:"post_code" binding:"required"`
	Phone     string `form:"phone" json:"phone" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Country   string `form:"country" json:"country" binding:"required"`
}

type ShipInfoParam struct {
	Address   string `form:"address" json:"address" binding:"required"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	City      string `form:"city" json:"city" binding:"required"`
	State     string `form:"state" json:"state"`
	PostCode  string `form:"post_code" json:"post_code" binding:"required"`
	Country   string `form:"country" json:"country" binding:"required"`
}

type OrderParam struct {
	OrderCreatedAt string          `form:"order_created_at" json:"order_created_at" binding:"required"`
	OrderId        string          `form:"order_id" json:"order_id" binding:"required"`
	Amount         float32         `form:"amount" json:"amount" binding:"required"`
	Currency       string          `form:"currency" json:"currency" binding:"required"`
	Products       []*ProductParam `form:"products" json:"products" binding:"required"`
}

type ProductParam struct {
	Title              string  `form:"title" json:"title" binding:"required"`
	Quantity           float32 `form:"quantity" json:"quantity" binding:"required"`
	Price              float32 `form:"price" json:"price" binding:"required"`
	ProductUrl         string  `form:"product_url" json:"product_url" binding:"required"`
	ImageUrl           string  `form:"image_url" json:"image_url" binding:"required"`
	ProductsAttributes string  `form:"products_attributes" json:"products_attributes" `
}

type ClientPaymentListParam struct {
	PageParam      globalParams.MysqlPageParam
	TimestampRange globalParams.TimestampRange
}

type ConsultBillStateParam struct {
	OrderId string `form:"order_id" json:"order_id" binding:"required"`
}
