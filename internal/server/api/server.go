package api

import (
	"net"

	"github.com/sirupsen/logrus"
	grpcHandler "github.com/vasiliyantufev/gophkeeper/internal/server/api/handlers"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	grpcGophkeeper "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"google.golang.org/grpc"
)

// StartService - starts the gophkeeper server
func StartService(grpcHandler *grpcHandler.Handler, config *config.Config, log *logrus.Logger) {
	log.Infof("Start gophkeeper application %s ", config.GRPC)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", config.GRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcGophkeeper.RegisterGophkeeperServer(grpcServer, grpcHandler)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("gprc server: %v", err)
	}
}
