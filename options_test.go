package gormmom_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormmom"
	"gorm.io/gorm/schema"
)

func TestNewOptions(t *testing.T) {
	// The defaults: GORM's default naming strategy, and a name length cap of 63
	// 默认值：GORM 默认的命名策略，名字长度上限 63
	require.NoError(t, gormmom.NewConfig(&T学生{}, gormmom.NewOptions()).Validate())
}

func TestOptions_WithMaxIdentLength(t *testing.T) {
	// "idx_students_grade", the longest name of T学生, has 18 characters: it passes a cap of 18, and fails a cap of 17
	// T学生 里最长的名字 "idx_students_grade" 有 18 个字符：上限 18 能通过，上限 17 就不行
	require.NoError(t, gormmom.NewConfig(&T学生{}, gormmom.NewOptions().WithMaxIdentLength(18)).Validate())

	err := gormmom.NewConfig(&T学生{}, gormmom.NewOptions().WithMaxIdentLength(17)).Validate()
	require.ErrorContains(t, err, `index "idx_students_grade"`)

	require.Panics(t, func() {
		gormmom.NewOptions().WithMaxIdentLength(0)
	})
}

func TestOptions_WithNamingStrategy(t *testing.T) {
	// Pass the strategy the project runs GORM with, here GORM's default one
	// 传入项目里 GORM 用的命名策略，这里就是 GORM 默认的策略
	options := gormmom.NewOptions().WithNamingStrategy(schema.NamingStrategy{})
	require.NoError(t, gormmom.NewConfig(&T学生{}, options).Validate())

	require.Panics(t, func() {
		gormmom.NewOptions().WithNamingStrategy(nil)
	})
}
