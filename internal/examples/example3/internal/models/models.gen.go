// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: models.gen_test.go:32 -> models.TestGenerateModels.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

func (T *T生徒) Columns() *T生徒Columns {
	return T.TableColumns(gormcnm.NewPlainDecoration())
}

func (T *T生徒) TableColumns(decoration gormcnm.ColumnNameDecoration) *T生徒Columns {
	return &T生徒Columns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:   gormcnm.Cmn(T.ID, "id", decoration),
		N名前:  gormcnm.Cmn(T.N名前, "name", decoration),
		G学年:  gormcnm.Cmn(T.G学年, "grade", decoration),
		Tクラブ: gormcnm.Cmn(T.Tクラブ, "club", decoration),
	}
}

type T生徒Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID   gormcnm.ColumnName[uint]
	N名前  gormcnm.ColumnName[string]
	G学年  gormcnm.ColumnName[string]
	Tクラブ gormcnm.ColumnName[string]
}
