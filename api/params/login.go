package params

type LoginData struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password"  binding:"required" `
}
