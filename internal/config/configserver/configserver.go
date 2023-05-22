package config_server

import (
	"github.com/sirupsen/logrus"
)

type ConfigServer struct {
	GRPC           string       `env:"GRPC"`
	DebugLevel     logrus.Level `env:"DEBUG_LEVEL" envDefault:"debug" json:"debug_level"`
	DSN            string       `env:"DATABASE_DSN" json:"dsn"`
	MigrationsPath string       `env:"ROOT_PATH" envDefault:"file://./migrations"`
}

// NewConfigServer - creates a new instance with the configuration for the server
func NewConfigServer() *ConfigServer {
	configServer := ConfigServer{}
	return &configServer
}
