package gormmom_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormmom"
)

// Product1 是一个合法的母语模型示范：字段名用母语写，
// 但列名、索引名、表名都显式给了数据库安全的 ASCII 名字，把关器应当放行。
type Product1 struct {
	ID  uint    `gorm:"primaryKey"`
	P名称 string  `gorm:"column:name;type:varchar(100);not null;uniqueIndex:udx_product1_name"`
	P价格 float64 `gorm:"column:price;type:decimal(10,2);not null;index:idx_product1_price"`
	S库存 int     `gorm:"column:stock;type:int;default:0"`
	D描述 string  `gorm:"column:description;type:text"`
	S状态 string  `gorm:"column:status;type:varchar(20);default:'active';index:idx_product1_status"`
}

// TableName 指定 Product1 的 ASCII 表名
func (*Product1) TableName() string {
	return "product1s"
}

func TestConfigs_Validate(t *testing.T) {
	objects := []interface{}{
		&Product1{},
	}

	cfg := gormmom.NewConfigs(objects, gormmom.NewOptions())
	t.Log(cfg)

	require.NoError(t, cfg.Validate())
}

// 下面三个反例各踩一条失败分支：列名 / 索引名 / 表名，把关器都应当报错。
// 表名都给个最简单的 "a"，把注意力留给被测的那条分支。

// badColumn 的母语字段没给 column，列名会落到非 ASCII。
type badColumn struct {
	ID  uint
	X姓名 string // 无 column → 列名非 ASCII
}

func (*badColumn) TableName() string { return "a" }

// badIndex 的母语字段给了 index 但没给索引名，GORM 用字段名拼出非 ASCII 索引名。
type badIndex struct {
	ID  uint
	Y地址 string `gorm:"column:address;index"` // index 无名 → 索引名非 ASCII
}

func (*badIndex) TableName() string { return "a" }

// badTable 的 TableName 直接返回一个非法表名 "-"（连字符不在允许字符集内）。
type badTable struct {
	ID uint
}

func (*badTable) TableName() string { return "-" }

func TestConfigs_Validate_CatchesBadNaming(t *testing.T) {
	for name, object := range map[string]interface{}{
		"bad-column": &badColumn{},
		"bad-index":  &badIndex{},
		"bad-table":  &badTable{},
	} {
		t.Run(name, func(t *testing.T) {
			err := gormmom.NewConfigs([]interface{}{object}, gormmom.NewOptions()).Validate()
			require.Error(t, err)
			t.Log(err)
		})
	}
}
