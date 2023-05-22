package main

import (
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/config/configagent"
)

func main() {
	log := logrus.New()
	config := configagent.NewConfigClient(log)
	log.SetLevel(config.DebugLevel)
}
