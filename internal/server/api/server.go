package api

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	grpcHandler "github.com/vasiliyantufev/gophkeeper/internal/server/api/grpc"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	grpcGophkeeper "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"google.golang.org/grpc"
)

// StartGRPCService - starts the GRPC gophkeeper server
func StartGRPCService(grpcHandler *grpcHandler.Handler, config *config.Config, log *logrus.Logger) {
	log.Infof("Starting GRPC server %s ", config.AddressGRPC)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", config.AddressGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcGophkeeper.RegisterGophkeeperServer(grpcServer, grpcHandler)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed gprc server: %v", err)
	}
}

// StartRESTService - starts the REST gophkeeper server
func StartRESTService(r *chi.Mux, config *config.Config, log *logrus.Logger) {
	log.Infof("Starting REST server %v\n", config.AddressREST)
	if lis := http.ListenAndServe(config.AddressREST, r); lis != nil {
		log.Fatalf("failed to listen: %v", lis)
	}
}
