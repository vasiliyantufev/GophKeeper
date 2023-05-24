package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/api/server"
	grpcHandler "github.com/vasiliyantufev/gophkeeper/internal/api/server/handlers"
	"github.com/vasiliyantufev/gophkeeper/internal/config/configserver"
	"github.com/vasiliyantufev/gophkeeper/internal/database"
)

func main() {
	logger := logrus.New()
	config := configserver.NewConfigServer(logger)
	logger.SetLevel(config.DebugLevel)

	db, err := database.New(config, logger)
	if err != nil {
		logger.Error(err)
	} else {
		defer db.Close()
	}

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	handlerGrpc := grpcHandler.NewHandler(db, logger)
	go server.StartService(handlerGrpc, config, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
