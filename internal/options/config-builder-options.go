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

	secrets map[string]string
}

func (c *ConfigBuilderOptions[T]) Merge(cfg *T) {
	mergeValues(
		c.Config,
		cfg,
	)
}

func (c *ConfigBuilderOptions[T]) Secrets() map[string]string {
	return c.secrets
}

func mergeValues(dst any, src any) {
	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)

	if dstValue.Kind() == reflect.Ptr {
		dstValue = dstValue.Elem()
	}

	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}

	for i := 0; i < dstValue.NumField(); i++ {
		if !dstValue.Type().Field(i).IsExported() {
			continue
		}

		dstField := dstValue.Field(i)
		srcField := srcValue.Field(i)

		if (srcField.Kind() != dstField.Kind()) ||
			(srcField.Kind() == reflect.Slice && srcField.IsNil()) ||
			srcField.IsZero() {
			continue
		}

		if dstField.Kind() == reflect.Struct {
			mergeValues(dstField.Addr().Interface(), srcField.Addr().Interface())
			continue
		} else if dstField.Kind() == reflect.Ptr && dstField.Elem().Kind() == reflect.Struct && !dstField.IsNil() {
			mergeValues(dstField.Interface(), srcField.Interface())
			continue
		}

		dstField.Set(srcField)
	}
}
