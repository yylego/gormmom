package utils

import (
	"regexp"
	"strconv"
	"sync"

	"github.com/yylego/rese"
	"gorm.io/gorm/schema"
)

// ParseSchema parses an object into a GORM schema by reflection, using the given
// naming strategy to derive the table / column / index names. A nil namer falls
// back to GORM's default strategy.
//
// ParseSchema 通过反射把对象解析成 GORM 模式，用传入的命名策略推导表名/列名/索引名。
// namer 为 nil 时回退到 GORM 的默认策略。
func ParseSchema(object any, namer schema.Namer) *schema.Schema {
	if namer == nil {
		namer = schema.NamingStrategy{}
	}
	return rese.P1(schema.Parse(object, &sync.Map{}, namer))
}

// NewCommonRegexp builds the regexp that checks a name's length and charset:
// only ASCII letters, digits and underscore, 1..maxLen characters.
// A leading space such as `gorm:"column: name;"` is accepted by GORM but rejected
// here on purpose, to keep names clean and avoid downstream surprises.
//
// NewCommonRegexp 创建检查名字长度和字符的正则：只允许 ASCII 字母、数字、下划线，长度 1..maxLen。
// 当列名前带个前导空格 比如 `gorm:"column: name;"` 时，在 gorm 中也是可以用的，
// 但该规则里故意不允许这种情况，保持名字干净、避免后续出别的问题。
func NewCommonRegexp(maxLen int) *regexp.Regexp {
	return regexp.MustCompile(`^[a-zA-Z0-9_]{1,` + strconv.Itoa(maxLen) + `}$`)
}
