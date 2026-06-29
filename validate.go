package gormmom

import (
	"github.com/yylego/erero"
	"github.com/yylego/gormmom/internal/utils"
	"github.com/yylego/must"
)

// Validate checks GORM naming across the batch, returning the first issue it finds.
// It catches native-language table names, column names and index names that databases reject.
//
// Validate 校验整批的 GORM 命名，返回遇到的第一个问题
// 它能抓出数据库不接受的母语表名、列名和索引名
func (configs Configs) Validate() error {
	for _, config := range configs {
		if err := config.Validate(); err != nil {
			return erero.Wro(err)
		}
	}
	return nil
}

// Validate checks one struct's table name, column names and index names.
// Each name must fit the database-safe charset and length.
//
// Validate 校验单个结构体的表名、列名和索引名
// 每个名字都要落在数据库安全的字符集和长度内
func (cfg *Config) Validate() error {
	structName := cfg.structI.structName
	gormSchema := cfg.structI.gormSchema

	// The name length cap must be positive; a non-positive cap is a config bug.
	// 名字长度上限必须为正数；非正数说明配置有误
	must.True(cfg.options.MaxNameLength > 0)

	nameRegexp := utils.NewCommonRegexp(cfg.options.MaxNameLength)

	// Check the table name. A native-language table name produces a non-ASCII
	// name, so write a TableName() method that returns a database-safe name.
	// 检查表名。母语表名会产出非 ASCII 名字，请写一个 TableName() 方法返回数据库安全的名字
	if !nameRegexp.MatchString(gormSchema.Table) {
		return erero.Errorf("struct %s table %q is not database-safe, write a TableName() method", structName, gormSchema.Table)
	}

	// Check each column name.
	// 检查每个列名
	for _, field := range gormSchema.Fields {
		if field.DBName == "" {
			continue
		}
		if !nameRegexp.MatchString(field.DBName) {
			return erero.Errorf("struct %s field %s column %q is not database-safe, write an explicit column name", structName, field.Name, field.DBName)
		}
	}

	// Check each index name. When you do not write an explicit index name,
	// GORM builds one from field names, so native-language fields produce a
	// non-ASCII index name. This check forces you to write a safe index name.
	// 检查每个索引名。当你不写显式索引名时，GORM 会用字段名拼出索引名，
	// 母语字段因此会拼出非 ASCII 的索引名。这个检查强制你写一个安全的索引名
	for _, index := range gormSchema.ParseIndexes() {
		if !nameRegexp.MatchString(index.Name) {
			return erero.Errorf("struct %s index %q is not database-safe, write an explicit index name", structName, index.Name)
		}
	}

	return nil
}
