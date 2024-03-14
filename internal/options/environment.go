package options

import (
	"github.com/noenik/go-config-loader/internal/utils"
	"os"
	"reflect"
	"strings"
)

const tag = "env"

type WithEnvironmentVariables[T any] struct {
	Prefix string
}

func (opt *WithEnvironmentVariables[T]) Apply(options *ConfigBuilderOptions[T]) error {
	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.Split(env, "=")
		if len(parts) != 2 {
			continue
		}

		var (
			key   = parts[0]
			value = parts[1]
		)

		if opt.Prefix == "" || opt.Prefix == "-" {
			envVars[key] = value
			continue
		}

		if !strings.HasPrefix(key, opt.Prefix) || len(key) == len(opt.Prefix) {
			continue
		}

		key = key[len(opt.Prefix):]
		envVars[key] = value
	}

	utils.FindAndApplyValuesByTag(envVars, reflect.ValueOf(options.Config).Elem(), tag, 1)

	return nil
}
