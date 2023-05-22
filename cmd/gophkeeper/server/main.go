package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/api/server"
	grpcHandler "github.com/vasiliyantufev/gophkeeper/internal/api/server/handlers"
	"github.com/vasiliyantufev/gophkeeper/internal/config/configserver"
)

func main() {
	log := logrus.New()
	config := configserver.NewConfigServer(log)
	log.SetLevel(config.DebugLevel)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	handlerGrpc := grpcHandler.NewHandler()
	go server.StartService(handlerGrpc, config, log)

	<-ctx.Done()
	log.Info("server shutdown on signal with:", ctx.Err())
}
