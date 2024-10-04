package configs

type (
	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Host string `yaml:"host" env:"HTTP_HOST" env-default:"0.0.0.0"`
		Port string `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
	}
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}
)
