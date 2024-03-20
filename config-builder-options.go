package configbuilder

import (
	"github.com/noenik/go-config-loader/internal/options"
)

func (builder *configBuilder[T]) WithDefaultConfigFiles() ConfigBuilder[T] {
	builder.options = append(
		builder.options,
		&options.WithJsonFile[T]{Required: true},
		&options.WithEnvJsonFile[T]{},
	)
	return builder
}

func (builder *configBuilder[T]) WithDefaults() ConfigBuilder[T] {
	return builder.WithDefaultConfigFiles().
		WithDefaultVersionFile().
		WithEnvironmentVariables().
		WithDockerSecrets()
}

func (builder *configBuilder[T]) WithDefaultVersionFile() ConfigBuilder[T] {
	builder.options = append(builder.options, &options.WithVersionFile[T]{})
	return builder
}

func (builder *configBuilder[T]) WithDockerSecrets(dirs ...string) ConfigBuilder[T] {
	builder.options = append(builder.options,
		&options.WithDockerSecrets[T]{
			Dirs: dirs,
		},
	)
	return builder
}

func (builder *configBuilder[T]) WithEnvironmentVariables(prefix ...string) ConfigBuilder[T] {
	var pf string

	if len(prefix) > 1 {
		panic("there should be at most one prefix")
	}

	if len(prefix) == 1 {
		pf = prefix[0]
	}

	builder.options = append(builder.options,
		&options.WithEnvironmentVariables[T]{
			Prefix: pf,
		},
	)

	return builder
}

func (builder *configBuilder[T]) WithFunc(fn func(configObj *T)) ConfigBuilder[T] {
	builder.options = append(builder.options,
		&options.WithFunc[T]{
			Func: fn,
		},
	)
	return builder
}

func (builder *configBuilder[T]) WithJsonFile(path string, optional ...bool) ConfigBuilder[T] {
	var isOptional bool
	if len(optional) > 0 {
		isOptional = optional[0]
	}
	builder.options = append(builder.options, &options.WithJsonFile[T]{FilePath: path, Required: !isOptional})
	return builder
}

func (builder *configBuilder[T]) WithJsonString(config string) ConfigBuilder[T] {
	builder.options = append(builder.options, &options.WithJsonString[T]{ConfigJson: config})
	return builder
}

func (builder *configBuilder[T]) WithVersionFile(path string) ConfigBuilder[T] {
	builder.options = append(builder.options, &options.WithVersionFile[T]{FilePath: path})
	return builder
}

func (builder *configBuilder[T]) WithVersionString(version string) ConfigBuilder[T] {
	builder.options = append(builder.options, &options.WithVersionString[T]{Version: version})
	return builder
}
