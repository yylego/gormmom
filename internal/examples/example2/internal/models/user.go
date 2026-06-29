package models

// Tユーザー 是日语母语模型示范：字段名用日语，列名/索引名给显式的数据库安全 ASCII 名字。
type Tユーザー struct {
	ID     uint   `gorm:"primaryKey"`
	Uユーザー名 string `gorm:"column:username;uniqueIndex:udx_users_username"`
	Eメール   string `gorm:"column:email;index:idx_users_email"`
	A年齢    int    `gorm:"column:age"`
	D電話    string `gorm:"column:phone"`
	J住所    string `gorm:"column:address"`
	Sステータス string `gorm:"column:status;index:idx_users_status"`
}

func (*Tユーザー) TableName() string {
	return "users"
}
