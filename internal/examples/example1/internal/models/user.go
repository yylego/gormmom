package models

// T用户 是汉语母语模型示范：字段名用汉语，列名/索引名给显式的数据库安全 ASCII 名字。
type T用户 struct {
	ID   uint   `gorm:"primaryKey"`
	U用户名 string `gorm:"column:username;uniqueIndex:udx_users_username"`
	E邮箱  string `gorm:"column:email;index:idx_users_email"`
	A年龄  int    `gorm:"column:age"`
	D电话  string `gorm:"column:phone"`
	J住所  string `gorm:"column:address"`
	S状态  string `gorm:"column:status;index:idx_users_status"`
}

func (*T用户) TableName() string {
	return "users"
}
