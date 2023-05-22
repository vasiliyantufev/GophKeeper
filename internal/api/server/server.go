package server

import (
	"net"

	"github.com/sirupsen/logrus"
	grpcHandler "github.com/vasiliyantufev/gophkeeper/internal/api/server/handlers"
	"github.com/vasiliyantufev/gophkeeper/internal/config/configserver"
	grpcGophkeeper "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"google.golang.org/grpc"
)

// StartService - starts the gophkeeper server
func StartService(grpcHandler *grpcHandler.Handler, config *configserver.ConfigServer, log *logrus.Logger) {
	log.Info("Start gophkeeper application %v\n", config.GRPC)

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
