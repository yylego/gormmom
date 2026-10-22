package gormmom

// Config holds one target struct plus its options.
//
// Config 保存单个目标结构体及其选项
type Config struct {
	structI *gormStruct
	options *Options
}

// NewConfig builds a Config from an object plus options, parsing the object by reflection.
//
// NewConfig 用对象和选项构造 Config，通过反射解析对象
func NewConfig(object any, options *Options) *Config {
	return &Config{
		structI: parseObject(object, options.NamingStrategy),
		options: options,
	}
}
