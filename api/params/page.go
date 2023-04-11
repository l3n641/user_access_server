package params

type PageParam struct {
	Page     int64 `form:"page" json:"page" binding:"required"`
	PageSize int64 `form:"page_size" json:"page_size" binding:"required"`
}
