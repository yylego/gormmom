// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: gormcnm_test.go:33 -> models.TestGenGormMomAndCnm.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

// 这个文件将被 gormcngen 自动生成和更新

func (T *T사용자) Columns() *T사용자Columns {
	return &T사용자Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:    gormcnm.Cnm(T.ID, "id"),
		U사용자명: gormcnm.Cnm(T.U사용자명, "username"),
		E이메일:  gormcnm.Cnm(T.E이메일, "email"),
		N나이:   gormcnm.Cnm(T.N나이, "age"),
		J전화:   gormcnm.Cnm(T.J전화, "phone"),
		J주소:   gormcnm.Cnm(T.J주소, "address"),
		S상태:   gormcnm.Cnm(T.S상태, "status"),
	}
}

type T사용자Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID    gormcnm.ColumnName[uint]
	U사용자명 gormcnm.ColumnName[string]
	E이메일  gormcnm.ColumnName[string]
	N나이   gormcnm.ColumnName[int]
	J전화   gormcnm.ColumnName[string]
	J주소   gormcnm.ColumnName[string]
	S상태   gormcnm.ColumnName[string]
}
