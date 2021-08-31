This package is part of the [go-kit](https://github.com/Sef1995/go-kit).

# Config loader

The config loader standardizes and simplifies the way we load configuration settings.

The project's env prefix is passed to the loader. The loader uses it to read the `{prefix}_CONFIG_DIRECTORY` variable. From the directory that that variable points to, the *config.yml* file is read. If the env var is not set, the directory defaults to the current directory.

## Example

A project's `NewConfig()` method could look like the following snippet:

```
func NewConfig() (*Config, error) {
	loader := config.NewLoader(EnvPrefix)
	conf := &Config{}
	err := loader.Load(conf)

	return conf, err
}
```