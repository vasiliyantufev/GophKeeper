package events

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	clientConfig "github.com/vasiliyantufev/gophkeeper/internal/client/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api/handlers"
	serverConfig "github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	grpcKeeper "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/entity"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/file"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func TestEvents(t *testing.T) {

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
		t.Fatalf("test containers failed: %v", err)
	}
	container.Start(context.Background())
	stopTime := time.Second
	defer container.Stop(context.Background(), &stopTime)
	databaseURI, err := container.ConnectionString(context.Background(), "sslmode=disable")

	// setting
	logger := logrus.New()
	db, err := database.New(&serverConfig.Config{DSN: databaseURI}, logger)
	if err != nil {
		t.Fatalf("db init failed: %v", err)
	}
	err = db.CreateTablesMigration("file://../../../../migrations")
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	// configs
	serverConfig := &serverConfig.Config{
		GRPC:                "localhost:50051",
		DSN:                 databaseURI,
		AccessTokenLifetime: 300 * time.Second,
		DebugLevel:          logrus.DebugLevel,
		FileFolder:          "../../../../data/test_keeper",
	}

	clientConfig := &clientConfig.Config{
		GRPC:       "localhost:50051",
		DebugLevel: logrus.DebugLevel,
		FileFolder: "../../../../data/test_keeper",
	}

	// repositories
	userRepository := user.New(db)
	fileRepository := file.New(db)
	storage := storage.New("/tmp")
	entityRepository := entity.New(db)
	tokenRepository := token.New(db)

	// setup server service
	handlerGrpc := *handlers.NewHandler(db, serverConfig, userRepository, fileRepository, &storage,
		entityRepository, tokenRepository, logger)
	lis, err := net.Listen("tcp", serverConfig.GRPC)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	grpcKeeper.RegisterGophkeeperServer(s, &handlerGrpc)

	go func() {
		if err = s.Serve(lis); err != nil {
			t.Fatalf("server exited with error: %v", err)
		}
	}()
	connectionServer, err := grpc.Dial(serverConfig.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("connection server with error: %v", err)
	}
	grpcKeeper.NewGophkeeperClient(connectionServer)

	// setup client service
	connectionClient, err := grpc.Dial(clientConfig.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("connection client with error: %v", err)
	}
	grpcClient := grpcKeeper.NewGophkeeperClient(connectionClient)
	client := NewEvent(context.Background(), clientConfig, logger, grpcClient)

	// -- TESTS --
	t.Run("ping db", func(t *testing.T) {
		_, err = client.Ping()
		assert.NoError(t, err, "failed ping db")
	})

}
