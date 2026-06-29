package models

import "time"

type Example struct {
	ID        int32  `gorm:"column:id;primaryKey;" json:"id"`
	V名称       string `gorm:"column:name;type:text"`
	V字段       string `gorm:"column:field;"`
	V性别       string `gorm:"column:gender;"`
	V特殊       string `gorm:"column:special;type:int32"`
	V年龄       int    `gorm:"column:age;" json:"age"` //理论上不要直接给model添加json标签，因为那是view层的逻辑，但实际上假如非这样做也能处理
	Rank      int32
	V身高       int32     `gorm:"column:height;"`
	V体重       int32     `gorm:"column:weight;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
