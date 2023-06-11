package configserver

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type ConfigServer struct {
	GRPC                string        `env:"GRPC"`
	DSN                 string        `env:"DATABASE_DSN"`
	MigrationsPath      string        `env:"ROOT_PATH" envDefault:"file://./migrations"`
	DebugLevel          logrus.Level  `env:"DEBUG_LEVEL" envDefault:"debug"`
	AccessTokenLifetime time.Duration `env:"ACCESS_TOKEN_LIFETIME"`
}

// NewConfigServer - creates a new instance with the configuration for the server
func NewConfigServer(log *logrus.Logger) *ConfigServer {
	// Set default values
	configServer := ConfigServer{
		GRPC:                "localhost:8080",
		DSN:                 "host=localhost port=5432 user=user password=password dbname=gophkeeper sslmode=disable",
		AccessTokenLifetime: 300 * time.Second,
	}

	flag.StringVar(&configServer.GRPC, "g", configServer.GRPC, "Server address")
	flag.StringVar(&configServer.DSN, "d", configServer.DSN, "Database configuration")
	flag.Parse()
	err := env.Parse(&configServer)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(configServer)

	return &configServer
}
