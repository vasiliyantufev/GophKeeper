package handlers

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	grpcKeeper "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	// initiate postgres container
	container, err := postgres.RunContainer(context.Background(),
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal(err)
	}

	container.Start(context.Background())
	stopTime := time.Second
	defer container.Stop(context.Background(), &stopTime)

	databaseURI, err := container.ConnectionString(context.Background(), "sslmode=disable")

	logger := logrus.New()
	db, err := database.New(&config.Config{DSN: databaseURI}, logger)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	err = db.CreateTablesMigration("file://../../../../migrations")
	if err != nil {
		log.Fatal(err)
	}

	config := &config.Config{
		GRPC:                "localhost:8080",
		DSN:                 databaseURI,
		AccessTokenLifetime: 300 * time.Second,
	}

	var handlerGrpc = NewHandler(db, config, nil, nil, nil, nil,
		nil, nil, nil, nil, logger)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	grpcKeeper.RegisterGophkeeperServer(s, handlerGrpc)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufferDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestHandlers(t *testing.T) {

	//ctx := context.Background()
	//conn, err := grpc.DialContext(ctx, "bufferNet", grpc.WithContextDialer(bufferDialer), grpc.WithInsecure())
	//if err != nil {
	//	t.Fatalf("Failed to dial bufferNet: %v", err)
	//}
	//defer conn.Close()
	//client := grpcKeeper.NewGophkeeperClient(conn)
	//resp, err := client.HandlePing(ctx, &grpcKeeper.PingRequest{})
	//if err != nil {
	//	t.Fatalf("Ping failed: %v", err)
	//}
	//log.Printf("Response: %+v", resp)
	// Test for output here.
}
