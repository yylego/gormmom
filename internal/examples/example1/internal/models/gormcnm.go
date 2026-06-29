// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: gormcnm_test.go:36 -> models.TestGen.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

func (T *T用户) Columns() *T用户Columns {
	return &T用户Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:   gormcnm.Cnm(T.ID, "id"),
		U用户名: gormcnm.Cnm(T.U用户名, "username"),
		E邮箱:  gormcnm.Cnm(T.E邮箱, "email"),
		A年龄:  gormcnm.Cnm(T.A年龄, "age"),
		D电话:  gormcnm.Cnm(T.D电话, "phone"),
		J住所:  gormcnm.Cnm(T.J住所, "address"),
		S状态:  gormcnm.Cnm(T.S状态, "status"),
	}
}

type T用户Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID   gormcnm.ColumnName[uint]
	U用户名 gormcnm.ColumnName[string]
	E邮箱  gormcnm.ColumnName[string]
	A年龄  gormcnm.ColumnName[int]
	D电话  gormcnm.ColumnName[string]
	J住所  gormcnm.ColumnName[string]
	S状态  gormcnm.ColumnName[string]
}
