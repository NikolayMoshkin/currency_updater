package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppConfig    AppConfig
	DBConfig     DBConfig
	SourceConfig SourceConfig
}

type AppConfig struct {
	UpdatePeriodSec int `env:"UPDATE_PERIOD_SEC" envDefault:"60"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	Password string `env:"DB_PASSWORD,required,notEmpty"`
	Name     string `env:"DB_NAME" envDefault:"main"`
}

type SourceConfig struct {
	SourceWithCredsApiKey string `env:"SOURCE_WITH_CREDS,required,notEmpty"`
}

func Load() (Config, error) {
	return env.ParseAs[Config]()
}
