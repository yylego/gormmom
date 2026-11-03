// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: models.gen_test.go:32 -> models.TestGenerateModels.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

func (T *T學生) Columns() *T學生Columns {
	return T.TableColumns(gormcnm.NewPlainDecoration())
}

func (T *T學生) TableColumns(decoration gormcnm.ColumnNameDecoration) *T學生Columns {
	return &T學生Columns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:  gormcnm.Cmn(T.ID, "id", decoration),
		X姓名: gormcnm.Cmn(T.X姓名, "name", decoration),
		N年級: gormcnm.Cmn(T.N年級, "grade", decoration),
		S社團: gormcnm.Cmn(T.S社團, "club", decoration),
	}
}

type T學生Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID  gormcnm.ColumnName[uint]
	X姓名 gormcnm.ColumnName[string]
	N年級 gormcnm.ColumnName[string]
	S社團 gormcnm.ColumnName[string]
}
