package gormmom

import (
	"sync"

	"github.com/yylego/rese"
	"github.com/yylego/syntaxgo/syntaxgo_reflect"
	"gorm.io/gorm/schema"
)

// Config holds one target struct, parsed into a GORM schema, plus its options.
// It reads nothing from disk; the schema comes from reflection alone.
//
// Config 保存单个目标结构体（已解析成 GORM 模式）及其选项
// 不读取任何磁盘文件；模式完全来自反射
type Config struct {
	structName string         // Struct name // 结构体名称
	gormSchema *schema.Schema // GORM schema // GORM 模式
	options    *Options       // Validation settings // 校验设置
}

// NewConfig builds a Config from an object plus options, parsing the object by reflection.
// The naming strategy in the options decides how field names turn into database names.
//
// NewConfig 用对象和选项构造 Config，通过反射解析对象
// 选项里的命名策略决定字段名如何转成数据库名字
func NewConfig(object any, options *Options) *Config {
	return &Config{
		structName: syntaxgo_reflect.GetTypeNameV3(object),
		gormSchema: rese.P1(schema.Parse(object, &sync.Map{}, options.namingStrategy)),
		options:    options,
	}
}
