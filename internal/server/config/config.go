package config

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AddressGRPC         string        `env:"AddressGRPC"`
	AddressREST         string        `env:"AddressREST"`
	DSN                 string        `env:"DATABASE_DSN"`
	DebugLevel          logrus.Level  `env:"DEBUG_LEVEL"`
	AccessTokenLifetime time.Duration `env:"ACCESS_TOKEN_LIFETIME"`
	FileFolder          string        `env:"DATA_FOLDER"`
	TemplatePath        string        `env:"TEMPLATE_PATH"`
}

// NewConfig - creates a new instance with the configuration for the server
func NewConfig(log *logrus.Logger) *Config {
	// Set default values
	configServer := Config{
		AddressGRPC:         "localhost:8080",
		AddressREST:         "localhost:8088",
		DSN:                 "host=localhost port=5432 user=user password=password dbname=gophkeeper sslmode=disable",
		AccessTokenLifetime: 30000 * time.Second,
		FileFolder:          "./data/server_keeper",
		DebugLevel:          logrus.DebugLevel,
		TemplatePath:        "./web/templates/index.html",
	}

	flag.StringVar(&configServer.AddressGRPC, "g", configServer.AddressGRPC, "Server address GRPC")
	flag.StringVar(&configServer.AddressREST, "r", configServer.AddressREST, "Server address REST")
	flag.StringVar(&configServer.DSN, "d", configServer.DSN, "Database configuration")
	flag.StringVar(&configServer.FileFolder, "f", configServer.FileFolder, "File Folder")
	flag.Parse()
	err := env.Parse(&configServer)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(configServer)

	return &configServer
}
