package handlers

import (
	"context"
	"log"
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
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
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
	tokenRepository := token.New(db)

	// setup server service
	handlerGrpc := *NewHandler(db, config, userRepository, nil, nil, nil,
		nil, nil, nil, tokenRepository, logger)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	grpcKeeper.RegisterGophkeeperServer(s, &handlerGrpc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// -- TEST DATA --
	username := randomizer.RandStringRunes(10)
	password := "Password-00"
	password, err = encryption.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	// -- TESTS --
	t.Run("ping db", func(t *testing.T) {
		err = handlerGrpc.database.Ping()
		assert.NoError(t, err, "failed ping db")
	})

	t.Run("registration", func(t *testing.T) {
		_, err = handlerGrpc.HandleRegistration(context.Background(), &grpcKeeper.RegistrationRequest{Username: username, Password: password})
		assert.NoError(t, err, "registration failed")
	})

	t.Run("authentication", func(t *testing.T) {
		_, err = handlerGrpc.HandleAuthentication(context.Background(), &grpcKeeper.AuthenticationRequest{Username: username, Password: password})
		assert.NoError(t, err, "authentication failed")
	})

}
