package gormmom_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormmom"
)

func TestNewConfig(t *testing.T) {
	cfg := gormmom.NewConfig(&T学生{}, gormmom.NewOptions())
	require.NoError(t, cfg.Validate())
}
