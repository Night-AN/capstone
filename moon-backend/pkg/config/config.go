package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port            int           `mapstructure:"port"`
	Mode            string        `mapstructure:"mode"`
	ReadTimeout     time.Duration `mapstructure:"readTimeout"`
	WriteTimeout    time.Duration `mapstructure:"writeTimeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdownTimeout"`

	LogLevel       string `mapstructure:"logLevel"`
	LogOutputPaths string `mapstructure:"logOutputPaths"`

	DatabaseDriver string `mapstructure:"databaseDriver"`
	DSN            string `mapstructure:"dsn"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	v.SetEnvPrefix("moon")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
