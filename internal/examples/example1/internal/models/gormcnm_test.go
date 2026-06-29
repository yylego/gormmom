package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormcngen"
	"github.com/yylego/gormmom"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath/runtestpath"
)

func TestGen(t *testing.T) {
	objects := []interface{}{
		&T用户{},
	}

	require.True(t, t.Run("GenGormMom", func(t *testing.T) {
		cfg := gormmom.NewConfigs(objects, gormmom.NewOptions())
		t.Log(cfg)

		require.NoError(t, cfg.Validate())
	}))

	// 使用 require.True(t, t.Run(---)) 限制只有前一步成功才执行后一步的

	require.True(t, t.Run("GenGormCnm", func(t *testing.T) {
		absPath := osmustexist.FILE(runtestpath.SrcPath(t))
		t.Log(absPath)

		options := gormcngen.NewOptions().
			WithColumnClassExportable(true). //中间类型名称的样式为导出的 T学生Columns
			WithColumnsMethodRecvName("T").
			WithColumnsCheckFieldType(true)

		cfg := gormcngen.NewConfigs(objects, options, absPath)
		cfg.Gen()
	}))
}
