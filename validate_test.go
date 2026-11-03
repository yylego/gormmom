package gormmom_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormmom"
)

// T学生 passes: the fields are in Chinese, and each column / index / table name is given in ASCII
// T学生 能通过：字段名用汉语，列名、索引名、表名都显式给了 ASCII 名字
type T学生 struct {
	ID  uint   `gorm:"primaryKey"`
	X姓名 string `gorm:"column:name;uniqueIndex:udx_students_name"`
	N年龄 int    `gorm:"column:age"`
	B班级 string `gorm:"column:grade;index:idx_students_grade"`
}

func (*T学生) TableName() string {
	return "students"
}

func TestConfigs_Validate(t *testing.T) {
	require.NoError(t, gormmom.NewConfigs([]any{&T学生{}}, gormmom.NewOptions()).Validate())

	// The batch returns the issue of the struct that fails
	// 整批校验返回不合格的那个结构体的问题
	err := gormmom.NewConfigs([]any{&T学生{}, &T学生_漏写列名{}}, gormmom.NewOptions()).Validate()
	require.ErrorContains(t, err, "struct T学生_漏写列名 field X姓名 column")
}

func TestConfig_Validate(t *testing.T) {
	require.NoError(t, gormmom.NewConfig(&T学生{}, gormmom.NewOptions()).Validate())
}

// T学生_漏写列名 leaves out the column name of a Chinese field, so GORM falls back to a non-ASCII column
// T学生_漏写列名 的汉语字段没写列名，GORM 会退回到非 ASCII 的列名
type T学生_漏写列名 struct {
	ID  uint
	X姓名 string
}

func (*T学生_漏写列名) TableName() string {
	return "no_column_names"
}

func TestConfig_Validate_NoColumnName(t *testing.T) {
	err := gormmom.NewConfig(&T学生_漏写列名{}, gormmom.NewOptions()).Validate()
	require.ErrorContains(t, err, "struct T学生_漏写列名 field X姓名 column")
	require.ErrorContains(t, err, "write an explicit column name")
}

// T学生_漏写索引名 gives an index with no name, so GORM builds a non-ASCII index name from the Chinese field
// T学生_漏写索引名 的索引没写名字，GORM 会用汉语字段名拼出非 ASCII 的索引名
type T学生_漏写索引名 struct {
	ID  uint
	X姓名 string `gorm:"column:name;index"`
}

func (*T学生_漏写索引名) TableName() string {
	return "no_index_names"
}

func TestConfig_Validate_NoIndexName(t *testing.T) {
	err := gormmom.NewConfig(&T学生_漏写索引名{}, gormmom.NewOptions()).Validate()
	require.ErrorContains(t, err, "struct T学生_漏写索引名 index")
	require.ErrorContains(t, err, "write an explicit index name")
}

// T学生_漏写表名 has a Chinese struct name and no TableName method, so the table name is non-ASCII
// T学生_漏写表名 的结构体名是汉语，又没写 TableName 方法，表名就是非 ASCII 的
type T学生_漏写表名 struct {
	ID  uint
	M名称 string `gorm:"column:name"`
}

func TestConfig_Validate_NoTableName(t *testing.T) {
	err := gormmom.NewConfig(&T学生_漏写表名{}, gormmom.NewOptions()).Validate()
	require.ErrorContains(t, err, "struct T学生_漏写表名 table")
	require.ErrorContains(t, err, "write a TableName() method")
}
