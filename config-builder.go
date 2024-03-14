package configbuilder

import (
	"errors"
	"github.com/noenik/go-config-loader/internal/options"
	"os"
	"reflect"
)

func NewConfigBuilder[T any]() ConfigBuilder[T] {
	return &configBuilder[T]{}
}

type configBuilder[T any] struct {
	options []options.BuilderOption[T]
}

func (builder *configBuilder[T]) Build(cfgObj *T) error {
	ps := reflect.ValueOf(cfgObj)
	if ps.Elem().Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	if len(builder.options) == 0 {
		builder.WithDefaults()
	}

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}

	opts := options.ConfigBuilderOptions[T]{
		Config:      cfgObj,
		Environment: env,
	}

	envField := ps.Elem().FieldByName("Environment")
	if envField.IsValid() && envField.CanSet() && envField.Kind() == reflect.String {
		envField.SetString(opts.Environment)
	}

	for _, opt := range builder.options {
		err := opt.Apply(&opts)
		if err != nil {
			return err
		}
	}

	return nil
}
