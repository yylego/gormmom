package gormmom

import (
	"github.com/yylego/erero"
	"gorm.io/gorm/schema"
)

// Options carries the validation settings, set via NewOptions and the With methods.
//
// Options 承载校验设置，通过 NewOptions 和 With 系列方法设置
type Options struct {
	maxIdentLength int          // Length cap of the database identifiers: table, column and index names // 数据库标识符（表名、列名、索引名）的长度上限
	namingStrategy schema.Namer // GORM naming strategy // GORM 命名策略
}

// NewOptions builds Options with the defaults: a name length cap of 63 (as in Postgres),
// and the default GORM naming strategy.
//
// NewOptions 用默认值构造 Options：名字长度上限 63（与 Postgres 相同），以及 GORM 默认的命名策略
func NewOptions() *Options {
	return &Options{
		maxIdentLength: 63,
		namingStrategy: schema.NamingStrategy{},
	}
}

// WithMaxNameLength sets how long a table name, column name and index name can be, since databases reject names beyond it.
// The cap must be positive.
//
// WithMaxNameLength 设置表名、列名、索引名的长度上限，超过上限的名字数据库不接受
// 上限必须为正数
func (options *Options) WithMaxNameLength(maxIdentLength int) *Options {
	if maxIdentLength <= 0 {
		panic(erero.Errorf("WithMaxNameLength needs a positive cap, got %d", maxIdentLength))
	}
	options.maxIdentLength = maxIdentLength
	return options
}

// WithNamingStrategy sets the GORM naming strategy that derives the names to validate.
// Set it when the project runs GORM with a custom strategy, so the validation matches runtime.
//
// WithNamingStrategy 设置推导待校验名字所用的 GORM 命名策略
// 项目用了自定义命名策略时设置它，让校验结果和运行时一致
func (options *Options) WithNamingStrategy(namingStrategy schema.Namer) *Options {
	if namingStrategy == nil {
		panic(erero.New("WithNamingStrategy needs the strategy the project runs GORM with, got nil"))
	}
	options.namingStrategy = namingStrategy
	return options
}
