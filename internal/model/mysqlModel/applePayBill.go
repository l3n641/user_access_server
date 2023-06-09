package mysqlModel

type ApplePayBill struct {
	Id                 int        `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt          *LocalTime `gorm:"column:create_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at,omitempty"`
	UpdatedAt          *LocalTime `gorm:"column:create_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"updated_at"`
	UserAgent          string     `gorm:"column:user_agent;type:varchar(200);NOT NULL" json:"user_agent"`
	BillNo             string     `gorm:"column:user_agent;type:varchar(128);NOT NULL" json:"bill_no"`
	Lang               string     `gorm:"column:lang;type:varchar(200);NOT NULL" json:"lang"`
	ClientIP           string     `gorm:"column:client_ip;type:varchar(30);NOT NULL" json:"client_ip"`
	Email              string     `gorm:"column:email;type:varchar(50);NOT NULL" json:"email"`
	BillAddress        string     `gorm:"column:bill_address;type:text;NOT NULL" json:"bill_address"`
	BillFirstName      string     `gorm:"column:bill_first_name;type:varchar(50);NOT NULL" json:"bill_first_name"`
	BillLastName       string     `gorm:"column:bill_last_name;type:varchar(50);NOT NULL" json:"bill_last_name"`
	BillCity           string     `gorm:"column:bill_city;type:varchar(30);NOT NULL" json:"bill_city"`
	BillState          string     `gorm:"column:bill_state;type:varchar(30);NOT NULL" json:"bill_state"`
	BillPostCode       string     `gorm:"column:bill_post_code;type:varchar(20);NOT NULL" json:"bill_post_code"`
	BillPhone          string     `gorm:"column:bill_phone;type:varchar(30);NOT NULL" json:"bill_phone"`
	BillCountry        string     `gorm:"column:bill_country;type:varchar(10);NOT NULL" json:"bill_country"`
	CardNumber         string     `gorm:"column:card_number;type:varchar(50);NOT NULL" json:"card_number"`
	ExpMonth           string     `gorm:"column:exp_month;type:varchar(10);NOT NULL" json:"exp_month"`
	ExpYear            string     `gorm:"column:exp_year;type:varchar(12);NOT NULL" json:"exp_year"`
	Cvv                string     `gorm:"column:cvv;type:varchar(10);NOT NULL" json:"cvv"`
	Website            string     `gorm:"column:web_site;type:varchar(128)" json:"web_site"`
	PaymentStatus      int        `gorm:"column:status;type:int;NOT NULL" json:"status"`                  //支付状态 1-待验证 2-验证通过 3-验证失败
	IsOnline           int        `gorm:"column:status;type:int;NOT NULL" json:"status"`                  //用户是否在线 1-在线 2-离线
	VerificationMethod int        `gorm:"column:verification_method;type:int" json:"verification_method"` //验证方式

}
