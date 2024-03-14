package options

import (
	"reflect"
)

type BuilderOption[T any] interface {
	Apply(options *ConfigBuilderOptions[T]) error
}

type ConfigBuilderOptions[T any] struct {
	Config      *T
	Environment string
	VersionFile string
}

func (c *ConfigBuilderOptions[T]) Merge(cfg *T) {
	mergeValues(
		c.Config,
		cfg,
	)
}

func mergeValues(dst any, src any) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < dstValue.NumField(); i++ {
		dstField := dstValue.Field(i)
		srcField := srcValue.Field(i)

		if (srcField.Kind() != dstField.Kind()) ||
			(srcField.Kind() == reflect.Slice && srcField.IsNil()) ||
			srcField.IsZero() {
			continue
		}

		dstField.Set(srcField)
	}
}
