package gormmom_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormmom"
)

func TestNewConfigs(t *testing.T) {
	configs := gormmom.NewConfigs([]any{&T学生{}, &T学生_漏写列名{}}, gormmom.NewOptions())
	require.Len(t, configs, 2)
	require.ErrorContains(t, configs.Validate(), "struct T学生_漏写列名 field X姓名 column")
}
