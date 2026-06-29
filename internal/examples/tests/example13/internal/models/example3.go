package models

import "gorm.io/gorm"

type Example3 struct {
	gorm.Model
	U账号  string `gorm:"column:username;uniqueIndex:udx_example3s_username;NOT NULL"`
	N昵称  string `gorm:"column:nickname;NOT NULL"`
	V分数  uint64 `gorm:"column:score;index:idx_example3s_score;"`
	Age  int    `gorm:"column:age;index:idx_example3s_age;"`
	Uuid int    `gorm:"column:uuid;uniqueIndex:udx_example3s_uuid;index:idx_example3s_uuid;"`
	Rank int    `gorm:"column:rank;index:idx_example3s_rank;"`
}

func (*Example3) TableName() string {
	return "example3s"
}
