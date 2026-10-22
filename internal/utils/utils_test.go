package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm/schema"
)

type sampleStruct struct {
	ID  uint
	V名称 string `gorm:"column:name"`
}

func TestParseSchema(t *testing.T) {
	// Without a custom naming scheme, GORM's default naming applies
	// 命名策略传 nil 时回退到 GORM 默认策略
	gormSchema := ParseSchema(&sampleStruct{}, nil)
	require.Equal(t, "sample_structs", gormSchema.Table)
	require.Equal(t, "name", gormSchema.LookUpField("V名称").DBName)

	// A custom naming scheme decides the table name
	// 自定义命名策略决定表名
	gormSchema = ParseSchema(&sampleStruct{}, schema.NamingStrategy{TablePrefix: "t_"})
	require.Equal(t, "t_sample_structs", gormSchema.Table)
}

func TestNewCommonRegexp(t *testing.T) {
	nameRegexp := NewCommonRegexp(5)
	require.True(t, nameRegexp.MatchString("a_1B"))
	require.True(t, nameRegexp.MatchString("abcde"))   // at the cap
	require.False(t, nameRegexp.MatchString("abcdef")) // beyond the cap
	require.False(t, nameRegexp.MatchString(""))       // blank
	require.False(t, nameRegexp.MatchString(" name"))  // leading space, rejected on purpose
	require.False(t, nameRegexp.MatchString("名称"))     // non-ASCII
	require.False(t, nameRegexp.MatchString("a-b"))    // hyphen
}
