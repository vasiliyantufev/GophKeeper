package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api/grpc"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api/rest"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api/router"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/entity"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/file"
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
	binaryRepository := file.New(db)
	storage := storage.New("/tmp")
	entityRepository := entity.New(db)
	tokenRepository := token.New(db)

	handlerRest := resthandler.NewHandler(db, config, userRepository, tokenRepository, logger)
	routerService := router.Route(handlerRest)
	rs := chi.NewRouter()
	rs.Mount("/", routerService)

	handlerGrpc := grpchandler.NewHandler(db, config, userRepository, binaryRepository,
		&storage, entityRepository, tokenRepository, logger)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go api.StartGRPCService(handlerGrpc, config, logger)
	go api.StartRESTService(rs, config, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
