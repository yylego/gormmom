// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: models.gen_test.go:32 -> models.TestGenerateModels.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

func (T *T학생) Columns() *T학생Columns {
	return T.TableColumns(gormcnm.NewPlainDecoration())
}

func (T *T학생) TableColumns(decoration gormcnm.ColumnNameDecoration) *T학생Columns {
	return &T학생Columns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:   gormcnm.Cmn(T.ID, "id", decoration),
		I이름:  gormcnm.Cmn(T.I이름, "name", decoration),
		H학년:  gormcnm.Cmn(T.H학년, "grade", decoration),
		D동아리: gormcnm.Cmn(T.D동아리, "club", decoration),
	}
}

type T학생Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID   gormcnm.ColumnName[uint]
	I이름  gormcnm.ColumnName[string]
	H학년  gormcnm.ColumnName[string]
	D동아리 gormcnm.ColumnName[string]
}
