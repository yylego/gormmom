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
		&Estudiante{},
	}

	// Validate the naming with gormmom first, and generate the columns just when it passes
	// Primero se valida la nomenclatura con gormmom, y solo si pasa se generan las columnas
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
