package gormmom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Example1 struct {
	Name     string `gorm:"column:name"`
	Age      int    `gorm:"column:age"`
	Enrolled bool   `gorm:"column:enrolled"`
}

type Example2 struct {
	V姓名 string `gorm:"column:name"`
	V年龄 int    `gorm:"column:age"`
	V在读 bool   `gorm:"column:enrolled"`
}

func TestNewConfigs(t *testing.T) {
	configs := NewConfigs([]any{&Example1{}, &Example2{}}, NewOptions())
	require.Len(t, configs, 2)
	require.Equal(t, "Example1", configs[0].structI.structName)
	require.Equal(t, "Example2", configs[1].structI.structName)
}
