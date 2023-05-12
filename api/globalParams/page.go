package globalParams

import "time"

type PageParam struct {
	Page     int64 `form:"page" json:"page" binding:"required"`
	PageSize int64 `form:"page_size" json:"page_size" binding:"required"`
}

type MysqlPageParam struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}

type TimestampRange struct {
	DateStart int64 `form:"date_start" json:"date_start"`
	DateEnd   int64 `form:"date_end" json:"date_end"`
}

func (p *TimestampRange) GetStartLocalTime() time.Time {
	return time.Unix(p.DateStart/1000, 0).Local()
}

func (p *TimestampRange) GetEndLocalTime() time.Time {
	return time.Unix(p.DateEnd/1000, 0).Local()
}
