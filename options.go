package gormmom

import "gorm.io/gorm/schema"

// defaultNameLength is the database name length cap (Postgres uses 63).
//
// defaultNameLength 是数据库标识符长度上限（Postgres 为 63）
const defaultNameLength = 63

// Options carries validation settings.
// MaxNameLength caps how long a table name, column name or index name can be;
// databases reject names beyond it.
// NamingStrategy is the GORM naming strategy used to derive the names that get
// validated. It defaults to GORM's own default; set it only when your project
// runs GORM with a custom naming strategy, so validation matches runtime.
//
// Options 承载校验设置。
// MaxNameLength 限制表名、列名或索引名的长度上限，超过该长度的名字数据库不接受。
// NamingStrategy 是推导待校验名字所用的 GORM 命名策略，默认与 GORM 自身的默认值相同；
// 只有当你的项目用了自定义命名策略时才需要设置它，让校验结果和运行时一致。
type Options struct {
	MaxNameLength  int          // Name length cap // 名字长度上限
	NamingStrategy schema.Namer // GORM naming strategy // GORM 命名策略
}

// NewOptions builds Options with the default name length cap and the default
// GORM naming strategy.
//
// NewOptions 构造带默认长度上限和默认 GORM 命名策略的 Options
func NewOptions() *Options {
	return &Options{
		MaxNameLength:  defaultNameLength,
		NamingStrategy: schema.NamingStrategy{},
	}
}

// WithMaxNameLength sets the name length cap and returns the options for chaining.
//
// WithMaxNameLength 设置名字长度上限，返回自身以便链式调用
func (options *Options) WithMaxNameLength(maxNameLength int) *Options {
	options.MaxNameLength = maxNameLength
	return options
}

// WithNamingStrategy sets the GORM naming strategy and returns the options for chaining.
//
// WithNamingStrategy 设置 GORM 命名策略，返回自身以便链式调用
func (options *Options) WithNamingStrategy(namingStrategy schema.Namer) *Options {
	options.NamingStrategy = namingStrategy
	return options
}
