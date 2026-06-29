package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormcngen"
	"github.com/yylego/gormmom"
	"github.com/yylego/runpath/runtestpath"
)

func TestGenGormMomAndCnm(t *testing.T) {
	objects := []interface{}{
		&Tユーザー{},
	}

	require.True(t, t.Run("GenGormMom", func(t *testing.T) {
		cfg := gormmom.NewConfigs(objects, gormmom.NewOptions())
		t.Log("第一步：使用 gormmom 生成日语字段的 mom 标签")

		require.NoError(t, cfg.Validate())
	}))

	require.True(t, t.Run("GenGormCnm", func(t *testing.T) {
		absPath := runtestpath.SrcPath(t)
		t.Log("第二步：使用 gormcngen 生成类型安全的列方法")

		options := gormcngen.NewOptions().
			WithColumnClassExportable(true).
			WithColumnsMethodRecvName("T").
			WithColumnsCheckFieldType(true)

		cfg := gormcngen.NewConfigs(objects, options, absPath)
		cfg.Gen()
		t.Log("gormcngen 生成完成")
	}))
}
