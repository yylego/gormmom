package gormmom

import (
	"fmt"
	"regexp"
	"strconv"
)

// Validate checks GORM naming across the batch, and returns the first issue it finds.
// It catches native-language table names, column names and index names that databases reject.
//
// Validate 校验整批的 GORM 命名，返回找到的第一个问题
// 它能抓出数据库不接受的母语表名、列名和索引名
func (configs Configs) Validate() error {
	for _, config := range configs {
		if err := config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate checks one struct's table name, column names and index names, and returns the first issue it finds.
// Each name must fit the database-safe charset and length.
// The issue is an expected result, so Validate returns it with no log; the test that calls Validate shows it.
//
// Validate 校验单个结构体的表名、列名和索引名，返回找到的第一个问题
// 每个名字都要落在数据库安全的字符集和长度内
// 这个问题是预期内的结果，只返回、不打日志，由调用 Validate 的测试来展示
func (cfg *Config) Validate() error {
	structName := cfg.structName
	gormSchema := cfg.gormSchema
	nameRegexp := newNameRegexp(cfg.options.maxIdentLength)

	// Check the table name. A native-language table name produces a non-ASCII
	// name, so write a TableName() method that returns a database-safe name.
	// 检查表名。母语表名会产出非 ASCII 名字，请写一个 TableName() 方法返回数据库安全的名字
	if !nameRegexp.MatchString(gormSchema.Table) {
		return fmt.Errorf("struct %s table %q is not database-safe, write a TableName() method", structName, gormSchema.Table)
	}

	// Check each column name.
	// 检查每个列名
	for _, field := range gormSchema.Fields {
		if field.DBName == "" {
			continue
		}
		if !nameRegexp.MatchString(field.DBName) {
			return fmt.Errorf("struct %s field %s column %q is not database-safe, write an explicit column name", structName, field.Name, field.DBName)
		}
	}

	// Check each index name. When you do not write an explicit index name,
	// GORM builds one from field names, so native-language fields produce a
	// non-ASCII index name. This check forces you to write a safe index name.
	// 检查每个索引名。当你不写显式索引名时，GORM 会用字段名拼出索引名，
	// 母语字段因此会拼出非 ASCII 的索引名。这个检查强制你写一个安全的索引名
	for _, index := range gormSchema.ParseIndexes() {
		if !nameRegexp.MatchString(index.Name) {
			return fmt.Errorf("struct %s index %q is not database-safe, write an explicit index name", structName, index.Name)
		}
	}

	return nil
}

// newNameRegexp builds the regexp that checks a name's length and charset:
// just ASCII letters, digits and underscore, 1..maxLength characters.
// A leading space such as `gorm:"column: name;"` passes in GORM but not here, on purpose,
// to keep the names clean.
//
// newNameRegexp 创建检查名字长度和字符的正则：只允许 ASCII 字母、数字、下划线，长度 1..maxLength
// 列名前带空格（如 `gorm:"column: name;"`）在 GORM 里能用，这里故意不放行，保持名字干净
func newNameRegexp(maxLength int) *regexp.Regexp {
	return regexp.MustCompile(`^[a-zA-Z0-9_]{1,` + strconv.Itoa(maxLength) + `}$`)
}
