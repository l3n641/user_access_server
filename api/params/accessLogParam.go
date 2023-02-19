package params

type AccessLogPostParam struct {
	Domain    string `form:"Domain" json:"domain" binding:"required"`
	Uri       string `form:"uri" json:"uri"  binding:"required" `
	PageType  string `form:"page_type" json:"page_type"  binding:"required" `
	ClientIP  string `form:"client_ip" json:"client_ip"  binding:"required"`
	SessionID string `form:"session_id" json:"session_id" binding:"required"`
	UserAgent string `form:"user_agent" json:"user_agent"`
	Referer   string `form:"referer" json:"referer"`
}
