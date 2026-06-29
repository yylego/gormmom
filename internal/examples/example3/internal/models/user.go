package models

// T사용자 是韩语母语模型示范：字段名用韩语，列名/索引名给显式的数据库安全 ASCII 名字。
type T사용자 struct {
	ID    uint   `gorm:"primaryKey"`
	U사용자명 string `gorm:"column:username;uniqueIndex:udx_users_username"`
	E이메일  string `gorm:"column:email;index:idx_users_email"`
	N나이   int    `gorm:"column:age"`
	J전화   string `gorm:"column:phone"`
	J주소   string `gorm:"column:address"`
	S상태   string `gorm:"column:status;index:idx_users_status"`
}

func (*T사용자) TableName() string {
	return "users"
}
