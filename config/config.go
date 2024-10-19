package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Server Server
	Logger Logger
	Mysql  Mysql
}

type Server struct {
	Mode           string `envconfig:"MAG_ENV" default:"dev"`
	Port           string `envconfig:"MAG_SERVER_PORT" default:"8095"`
	TrustedProxies string `envconfig:"MAG_TRUSTED_PROXIES" default:"127.0.0.1/32"`
}

type Logger struct {
	Level       string `envconfig:"MAG_LOG_LEVEL" default:"debug"`
	Path        string `envconfig:"MAG_LOG_PATH" default:"./logs/access.log"`
	PrintStdOut bool   `envconfig:"MAG_STDOUT" default:"true"`
}

type Mysql struct {
	Host     string `envconfig:"MAG_MYSQL_HOST" default:"localhost:3306"`
	Driver   string `envconfig:"MAG_MYSQL_DRIVER" default:"mysql"`
	User     string `envconfig:"MAG_MYSQL_USER" default:"root"`
	Password string `envconfig:"MAG_MYSQL_PASSWORD" default:"1234"`
	Database string `envconfig:"MAG_MYSQL_DATABASE" default:"momssi"`
}

func LoadEnvConfig() (*EnvConfig, error) {
	var config EnvConfig
	if err := envconfig.Process("bac", &config); err != nil {
		return nil, err
	}

	if err := config.CheckValid(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *EnvConfig) CheckValid() error {
	return nil
}
