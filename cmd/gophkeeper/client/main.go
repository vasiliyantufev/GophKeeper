package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/config/configagent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcClient "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

func main() {
	log := logrus.New()
	config := configagent.NewConfigClient(log)
	log.SetLevel(config.DebugLevel)

	conn, err := grpc.Dial(config.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := grpcClient.NewGophkeeperClient(conn)
	resp, err := client.HandlePing(context.Background(), &grpcClient.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(resp)
}
