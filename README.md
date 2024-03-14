# Go Config Loader

This is a config loader for Go tailored to my own needs. It is mainly intended to be used in Docker services that follow 
the standardized structure I am using.

## Examples

### Basic
The minimum needed is the following
```go
type MyConfig struct {}

func main() {
    var cfg MyConfig
    
    err := cb.NewConfigBuilder[MyConfig]().Build(&cfg)
    if err != nil {
        panic(err)
    }   
}
```

This will run with all the defaults, meaning it will
* Load config from the json file `config.json`
* Add an environment specific config if it exists (and `GO_ENV` has a value), `config.<env>.json`
* Add version from the default version file
* Add environment variables
* Add any Docker secrets found in `/var/run/secrets`

### Complete

```go
type Version struct {
    Build string
    Version string
}

type MyConfig struct {
    Environment string
    MySetting   string `json:"mySetting"`
    MySecret    string `secret:"This_Is_My_Secret"`
    Version     *Version
}

func main() {
    var cfg MyConfig
    err := cb.NewConfigBuilder[MyConfig].
            WithDefaults(). // Explicitly set defaults to preserve them in case any other option is added
            // Alternatively add any option explicitly
            // WithJsonFile("my_json_file.json").
            // WithEnvironmentVariables().
            // WithDockerSecrets().
            Build(&cfg)
    if err != nil {
        panic(err)
    }   
}
```

This example assumes that the following files are available:

`./my_json_config.json` with content:

```json
{
  "mySetting": "my value"
}
```

`./version` with content similar to:

```
version=1.0.0
build=local
```

And finally `/var/run/secrets/This_Is_My_Secret`

Code completion will show the complete list of available options.

## Caveats

This is not really built to be used by others, so there is currently no support for other configuration file formats
like yaml, toml, ini and so on. I might add this on request if anyone is interested.