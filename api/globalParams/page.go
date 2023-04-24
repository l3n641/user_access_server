package globalParams

type PageParam struct {
	Page     int64 `form:"page" json:"page" binding:"required"`
	PageSize int64 `form:"page_size" json:"page_size" binding:"required"`
}

type MysqlPageParam struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}
