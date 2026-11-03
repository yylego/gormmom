// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: models.gen_test.go:32 -> models.TestGenerateModels.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

func (T *Estudiante) Columns() *EstudianteColumns {
	return T.TableColumns(gormcnm.NewPlainDecoration())
}

func (T *Estudiante) TableColumns(decoration gormcnm.ColumnNameDecoration) *EstudianteColumns {
	return &EstudianteColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:     gormcnm.Cmn(T.ID, "id", decoration),
		Nombre: gormcnm.Cmn(T.Nombre, "name", decoration),
		Año:    gormcnm.Cmn(T.Año, "grade", decoration),
		Club:   gormcnm.Cmn(T.Club, "club", decoration),
	}
}

type EstudianteColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID     gormcnm.ColumnName[uint]
	Nombre gormcnm.ColumnName[string]
	Año    gormcnm.ColumnName[string]
	Club   gormcnm.ColumnName[string]
}
