package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api"
	grpcHandler "github.com/vasiliyantufev/gophkeeper/internal/server/api/handlers"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/binary"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/entity"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
)

func main() {
	logger := logrus.New()
	config := config.NewConfig(logger)
	logger.SetLevel(config.DebugLevel)

	db, err := database.New(config, logger)
	if err != nil {
		logger.Fatal(err)
	} else {
		defer db.Close()
		db.CreateTablesMigration("file://../migrations")
	}

	userRepository := user.New(db)
	binaryRepository := binary.New(db)
	storage := storage.New("/tmp")
	entityRepository := entity.New(db)
	tokenRepository := token.New(db)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	var handlerGrpc = grpcHandler.NewHandler(db, config, userRepository, binaryRepository,
		&storage, entityRepository, tokenRepository, logger)
	go api.StartService(handlerGrpc, config, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
