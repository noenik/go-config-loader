package options

import "github.com/noenik/go-config-loader/interfaces"

type WithFunc[T any] struct {
	Func func(configObj *T, c interfaces.ContextAccessor)
}

func (opt *WithFunc[T]) Apply(options *ConfigBuilderOptions[T]) error {
	opt.Func(options.Config, options)
	return nil
}
