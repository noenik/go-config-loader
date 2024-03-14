package configbuilder

type ConfigBuilder[T any] interface {
	Build(cfgObj *T) error

	// WithConfigFile sets a specific config file to be loaded.
	//WithConfigFile(configFile string) ConfigBuilder[T]

	// WithDefaultConfigFiles indicates that the default config should be loaded. This is config.local.json for local dev and
	// config.json for other cases.
	WithDefaultConfigFiles() ConfigBuilder[T]

	// WithDefaults explicitly sets all configuration to the default. The default is reading a json config file from ./config.json,
	// reading environment specific config from ./config.<env>.json, getting version info from ./version, adding environment
	// variables, and reading Docker secrets.
	// This is equivalent to doing the following:
	//	builder.
	//	  WithDefaultConfigFiles().
	//	  WithDefaultVersionFile().
	//	  WithEnvironmentVariables().
	//	  WithDockerSecrets()
	WithDefaults() ConfigBuilder[T]

	// WithDefaultVersionFile specifies that version information should be loaded from the default version file ./version
	WithDefaultVersionFile() ConfigBuilder[T]

	// WithDockerSecrets loads any Docker secrets. By default, this option will look for files inside /var/run/secrets.
	// Any number of directories can be provided to override the location where secrets will be searched for.
	WithDockerSecrets(dirs ...string) ConfigBuilder[T]

	// WithEnvironmentVariables adds configuration by environment variables. Fields must be tagged with `env`. Optionally, a
	// prefix can be set. By default, this is "ST_". To disable prefix pass the value "-". Please note that the prefix will
	// be stripper from the value key and thus should not be included in the tag value.
	WithEnvironmentVariables(prefix ...string) ConfigBuilder[T]

	// WithFunc specifies a function to apply to the configuration. Please note that this option should be added at the end
	// so that other options may apply their effects first if the intention is to use fields populated by other options.
	WithFunc(fn func(configObj *T)) ConfigBuilder[T]

	// WithJsonFile specifies a json file that should be loaded.
	WithJsonFile(path string) ConfigBuilder[T]

	// WithJsonString specifies a serialized json string that should be loaded. This is useful in cases where for instance
	// the config is embedded at build time, or any other scenario where the configuration is provided as anything other
	// than a file
	WithJsonString(config string) ConfigBuilder[T]

	// WithVersionFile specifies a file that contains the version and build of the software. This is a simple text file
	// with two lines in key=value format:
	//	version=x.y.z
	//	build=a.b.c
	WithVersionFile(path string) ConfigBuilder[T]

	// WithVersionString specifies a string containing the version information. Please note that it needs to be in the same
	// format as the version file (with newline). See WithVersionFile.
	WithVersionString(version string) ConfigBuilder[T]

	// WithStringConfig adds a JSON formatted string to be parsed as the configuration. Useful for embedding or custom loading.
	//WithStringConfig(config string) ConfigBuilder[T]

	// WithStringVersion adds a JSON formatted string to be parsed as the configuration. Useful for embedding or custom loading.
	//WithStringVersion(version string) ConfigBuilder[T]
}
