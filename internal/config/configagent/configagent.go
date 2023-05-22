package configagent

import "github.com/sirupsen/logrus"

type ConfigClient struct {
	GRPC       string       `env:"GRPC"`
	DebugLevel logrus.Level `env:"DEBUG_LEVEL" envDefault:"debug"`
}

// NewConfigClient - creates a new instance with the configuration for the client
func NewConfigClient() *ConfigClient {
	configClient := ConfigClient{}
	return &configClient
}
