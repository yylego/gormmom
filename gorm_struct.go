package gormmom

import (
	"github.com/yylego/gormmom/internal/utils"
	"github.com/yylego/must"
	"github.com/yylego/syntaxgo/syntaxgo_reflect"
	"gorm.io/gorm/schema"
)

// gormStruct holds a parsed GORM struct: its name and schema.
//
// gormStruct 保存解析后的 GORM 结构体：名称与模式
type gormStruct struct {
	structName string         // Struct name // 结构体名称
	gormSchema *schema.Schema // GORM schema // GORM 模式
}

// newGormStruct builds a gormStruct from a struct name and its schema.
//
// newGormStruct 用结构体名和模式构造 gormStruct
func newGormStruct(structName string, gormSchema *schema.Schema) *gormStruct {
	return &gormStruct{
		structName: must.Nice(structName),
		gormSchema: must.Full(gormSchema),
	}
}

// parseObject builds a gormStruct from an object instance via reflection.
// It reads nothing from disk; the schema comes from reflection alone.
// The namer decides how field names turn into database names, so the same
// strategy your project runs GORM with produces the names that get validated.
//
// parseObject 通过反射从对象实例构造 gormStruct
// 不读取任何磁盘文件；模式完全来自反射
// namer 决定字段名如何转成数据库名字，因此用和项目里 GORM 相同的策略才能校验到一致的名字
func parseObject(object any, namer schema.Namer) *gormStruct {
	return newGormStruct(syntaxgo_reflect.GetTypeNameV3(object), utils.ParseSchema(object, namer))
}
