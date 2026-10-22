package gormmom

// Configs is a batch of Config instances, checked as a group.
//
// Configs 是一批 Config 实例，一起校验
type Configs []*Config

// NewConfigs builds a batch from objects plus shared options, parsing each by reflection.
//
// NewConfigs 用对象集合和共享选项构造一批配置，逐个通过反射解析
func NewConfigs(objects []any, options *Options) Configs {
	var configs = make([]*Config, 0, len(objects))
	for _, object := range objects {
		configs = append(configs, NewConfig(object, options))
	}
	return configs
}
