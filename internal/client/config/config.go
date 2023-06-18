package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type Config struct {
	GRPC       string       `env:"GRPC"`
	DebugLevel logrus.Level `env:"DEBUG_LEVEL" envDefault:"debug"`
}

// NewConfig - creates a new instance with the configuration for the client
func NewConfig(log *logrus.Logger) *Config {
	// Set default values
	configClient := Config{
		GRPC: "localhost:8080",
	}

	flag.StringVar(&configClient.GRPC, "g", configClient.GRPC, "Server address")
	flag.Parse()
	err := env.Parse(&configClient)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(configClient)

	return &configClient
}
