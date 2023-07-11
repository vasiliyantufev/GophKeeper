﻿package handlers

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/randomizer"
	serverConfig "github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/binary"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/entity"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	grpcKeeper "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func TestHandlers(t *testing.T) {

	// -- SETUP --
	// initiate postgres container
	container, err := postgres.RunContainer(context.Background(),
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatalf("Test containers failed: %v", err)
	}

	container.Start(context.Background())
	stopTime := time.Second
	defer container.Stop(context.Background(), &stopTime)

	databaseURI, err := container.ConnectionString(context.Background(), "sslmode=disable")

	logger := logrus.New()
	db, err := database.New(&serverConfig.Config{DSN: databaseURI}, logger)
	if err != nil {
		t.Fatalf("Db init failed: %v", err)
	}

	err = db.CreateTablesMigration("file://../../../../migrations")
	if err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	config := &serverConfig.Config{
		GRPC:                "localhost:8080",
		DSN:                 databaseURI,
		AccessTokenLifetime: 300 * time.Second,
		DebugLevel:          logrus.DebugLevel,
		FileFolder:          "../../../../data/test_keeper",
	}

	// repositories
	userRepository := user.New(db)
	binaryRepository := binary.New(db)
	storage := storage.New("/tmp")
	entityRepository := entity.New(db)
	tokenRepository := token.New(db)

	// setup server service
	handlerGrpc := *NewHandler(db, config, userRepository, binaryRepository, &storage,
		entityRepository, tokenRepository, logger)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	grpcKeeper.RegisterGophkeeperServer(s, &handlerGrpc)

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()

	// -- TEST DATA --
	var authenticatedUser *grpcKeeper.AuthenticationResponse
	data := randomizer.RandStringRunes(10)
	dataUpdate := randomizer.RandStringRunes(10)
	username := randomizer.RandStringRunes(10)
	password, _ := encryption.HashPassword("Password-00")
	name := randomizer.RandStringRunes(10)
	description := randomizer.RandStringRunes(10)
	metadata := model.MetadataEntity{Name: name, Description: description, Type: variables.Text.ToString()}
	jsonMetadata, _ := json.Marshal(metadata)

	// -- TESTS --
	t.Run("ping db", func(t *testing.T) {
		_, err = handlerGrpc.Ping(context.Background(), &grpcKeeper.PingRequest{})
		assert.NoError(t, err, "failed ping db")
	})

	t.Run("registration", func(t *testing.T) {
		_, err = handlerGrpc.Registration(context.Background(), &grpcKeeper.RegistrationRequest{Username: username, Password: password})
		assert.NoError(t, err, "registration failed")
	})

	t.Run("user exist", func(t *testing.T) {
		_, err = handlerGrpc.UserExist(context.Background(), &grpcKeeper.UserExistRequest{Username: username})
		assert.NoError(t, err, "user exist failed")
	})

	t.Run("authentication", func(t *testing.T) {
		authenticatedUser, err = handlerGrpc.Authentication(context.Background(), &grpcKeeper.AuthenticationRequest{Username: username, Password: password})
		assert.NoError(t, err, "authentication failed")
	})

	t.Run("create entity", func(t *testing.T) {
		_, err = handlerGrpc.HandleCreateEntity(context.Background(),
			&grpcKeeper.CreateEntityRequest{Data: []byte(data), Metadata: string(jsonMetadata),
				AccessToken: &grpcKeeper.Token{Token: authenticatedUser.AccessToken.Token, UserId: authenticatedUser.AccessToken.UserId,
					CreatedAt: authenticatedUser.AccessToken.CreatedAt, EndDateAt: authenticatedUser.AccessToken.EndDateAt}})
		assert.NoError(t, err, "create entity failed")
	})

	t.Run("update entity", func(t *testing.T) {
		_, err = handlerGrpc.HandleUpdateEntity(context.Background(),
			&grpcKeeper.UpdateEntityRequest{Name: name, Data: []byte(dataUpdate), Type: variables.Text.ToString(),
				AccessToken: &grpcKeeper.Token{Token: authenticatedUser.AccessToken.Token, UserId: authenticatedUser.AccessToken.UserId,
					CreatedAt: authenticatedUser.AccessToken.CreatedAt, EndDateAt: authenticatedUser.AccessToken.EndDateAt}})
		assert.NoError(t, err, "update entity failed")
	})

	t.Run("get list entity", func(t *testing.T) {
		_, err = handlerGrpc.HandleGetListEntity(context.Background(),
			&grpcKeeper.GetListEntityRequest{Type: variables.Text.ToString(),
				AccessToken: &grpcKeeper.Token{Token: authenticatedUser.AccessToken.Token, UserId: authenticatedUser.AccessToken.UserId,
					CreatedAt: authenticatedUser.AccessToken.CreatedAt, EndDateAt: authenticatedUser.AccessToken.EndDateAt}})
		assert.NoError(t, err, "get list failed")
	})

	t.Run("delete entity", func(t *testing.T) {
		_, err = handlerGrpc.HandleDeleteEntity(context.Background(),
			&grpcKeeper.DeleteEntityRequest{Name: name, Type: variables.Text.ToString(),
				AccessToken: &grpcKeeper.Token{Token: authenticatedUser.AccessToken.Token, UserId: authenticatedUser.AccessToken.UserId,
					CreatedAt: authenticatedUser.AccessToken.CreatedAt, EndDateAt: authenticatedUser.AccessToken.EndDateAt}})
		assert.NoError(t, err, "delete entity failed")
	})
}
