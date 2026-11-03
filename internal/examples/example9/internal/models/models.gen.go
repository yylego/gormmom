// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: models.gen_test.go:32 -> models.TestGenerateModels.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

func (T *HọcSinh) Columns() *HọcSinhColumns {
	return T.TableColumns(gormcnm.NewPlainDecoration())
}

func (T *HọcSinh) TableColumns(decoration gormcnm.ColumnNameDecoration) *HọcSinhColumns {
	return &HọcSinhColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:       gormcnm.Cmn(T.ID, "id", decoration),
		Tên:      gormcnm.Cmn(T.Tên, "name", decoration),
		Lớp:      gormcnm.Cmn(T.Lớp, "grade", decoration),
		CâuLạcBộ: gormcnm.Cmn(T.CâuLạcBộ, "club", decoration),
	}
}

type HọcSinhColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID       gormcnm.ColumnName[uint]
	Tên      gormcnm.ColumnName[string]
	Lớp      gormcnm.ColumnName[string]
	CâuLạcBộ gormcnm.ColumnName[string]
}
