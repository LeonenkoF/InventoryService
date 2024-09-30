package config

import (
	"errors"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"dev"`
	HTTPServer `yaml:"http_server"`
	Db_connect `yaml:"db_connect"`
}

type HTTPServer struct {
	Port        string        `yaml:"port" env:"TODO_PORT" env-default:"7540"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

type Db_connect struct {
	User     string `yaml:"user" env-default:"7540"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

func New() (*Config, error) {
	cfg := &Config{}

	configPath := "config.yml"

	if configPath == "" {
		return nil, errors.New("config path is not set")
	}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, errors.New("cannot read config")
	}

	return cfg, nil

}
