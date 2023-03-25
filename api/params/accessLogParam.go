package params

import "time"

type AccessLogPostParam struct {
	Domain    string `form:"Domain" json:"domain" binding:"required"`
	Uri       string `form:"uri" json:"uri"  binding:"required" `
	PageType  string `form:"page_type" json:"page_type"  binding:"required" `
	ClientIP  string `form:"client_ip" json:"client_ip"  binding:"required"`
	SessionID string `form:"session_id" json:"session_id" binding:"required"`
	UserAgent string `form:"user_agent" json:"user_agent"`
	Referer   string `form:"referer" json:"referer"`
}

type AccessDomainLogGetParam struct {
	PageParam
	Domain    string `form:"domain" json:"domain"`
	DateStart int64  `form:"date_start" json:"date_start"  binding:"required"`
	DateEnd   int64  `form:"date_end" json:"date_end"  binding:"required"`
}

func (p *AccessDomainLogGetParam) GetStartLocalTime() time.Time {
	return time.Unix(p.DateStart/1000, 0).Local()
}

func (p *AccessDomainLogGetParam) GetEndLocalTime() time.Time {
	return time.Unix(p.DateEnd/1000, 0).Local()
}

type AccessUserLogGetParam struct {
	PageParam
	Domain    string `form:"domain" json:"domain"`
	DateStart int64  `form:"date_start" json:"date_start"  binding:"required"`
	DateEnd   int64  `form:"date_end" json:"date_end"  binding:"required"`
	SessionID string `form:"session_id" json:"session_id" `
}

func (p *AccessUserLogGetParam) GetStartLocalTime() time.Time {
	return time.Unix(p.DateStart/1000, 0).Local()
}

func (p *AccessUserLogGetParam) GetEndLocalTime() time.Time {
	return time.Unix(p.DateEnd/1000, 0).Local()
}
