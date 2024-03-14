package options

type WithFunc[T any] struct {
	Func func(configObj *T)
}

func (opt *WithFunc[T]) Apply(options *ConfigBuilderOptions[T]) error {
	opt.Func(options.Config)
	return nil
}
