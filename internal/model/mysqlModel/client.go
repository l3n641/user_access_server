package mysqlModel

type Client struct {
	Id             int        `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt      *LocalTime `gorm:"column:create_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at"`
	LastAccessTime *LocalTime `gorm:"column:last_access_time;type:TIMESTAMP" json:"last_access_time"`
	Domain         string     `gorm:"column:user_agent;type:varchar(200);NOT NULL" json:"domain"`
	SessionID      string     `gorm:"column:user_agent;type:varchar(200);NOT NULL" json:"session_id"`
	ClientIP       string     `gorm:"column:client_ip;type:varchar(64);NOT NULL" json:"client_ip"`
	ClientCountry  string     `gorm:"column:client_country;type:varchar(20);NOT NULL" json:"client_country"`
	UserAgent      string     `gorm:"column:user_agent;type:varchar(200);NOT NULL" json:"user_agent"`
	Lang           string     `gorm:"column:lang;type:varchar(200);NOT NULL" json:"lang"`
}
