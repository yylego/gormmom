// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: gormcnm_test.go:33 -> models.TestGenGormMomAndCnm.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

// 这个文件将被 gormcngen 自动生成和更新

func (T *Tユーザー) Columns() *TユーザーColumns {
	return &TユーザーColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:     gormcnm.Cnm(T.ID, "id"),
		Uユーザー名: gormcnm.Cnm(T.Uユーザー名, "username"),
		Eメール:   gormcnm.Cnm(T.Eメール, "email"),
		A年齢:    gormcnm.Cnm(T.A年齢, "age"),
		D電話:    gormcnm.Cnm(T.D電話, "phone"),
		J住所:    gormcnm.Cnm(T.J住所, "address"),
		Sステータス: gormcnm.Cnm(T.Sステータス, "status"),
	}
}

type TユーザーColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID     gormcnm.ColumnName[uint]
	Uユーザー名 gormcnm.ColumnName[string]
	Eメール   gormcnm.ColumnName[string]
	A年齢    gormcnm.ColumnName[int]
	D電話    gormcnm.ColumnName[string]
	J住所    gormcnm.ColumnName[string]
	Sステータス gormcnm.ColumnName[string]
}
