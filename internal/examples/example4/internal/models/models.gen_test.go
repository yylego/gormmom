package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormcngen"
	"github.com/yylego/gormmom"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath/runtestpath"
)

//go:generate go test -run ^TestGenerateModels$ -tags=gormcngen_generate
func TestGenerateModels(t *testing.T) {
	objects := []any{
		&T학생{},
	}

	// Validate the naming with gormmom first, and generate the columns just when it passes
	// 먼저 gormmom 으로 이름을 검증하고, 통과했을 때만 열을 생성한다
	require.True(t, t.Run("ValidateNaming", func(t *testing.T) {
		require.NoError(t, gormmom.NewConfigs(objects, gormmom.NewOptions()).Validate())
	}))

	require.True(t, t.Run("GenerateColumns", func(t *testing.T) {
		absPath := osmustexist.FILE(runtestpath.SrcPath(t))
		t.Log(absPath)

		options := gormcngen.NewOptions().
			WithRecvName("T")

		cfg := gormcngen.NewConfigs(objects, options, absPath)
		cfg.WithGoBuildDirective(gormcngen.DirectiveSkipGenerate)
		cfg.Gen()
	}))
}
