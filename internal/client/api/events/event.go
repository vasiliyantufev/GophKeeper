package events

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/config"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

type Event struct {
	grpc    grpc.GophkeeperClient
	config  *config.Config
	logger  *logrus.Logger
	context context.Context
	grpc.UnimplementedGophkeeperServer
}

// NewEvent - creates a new grpc client instance
func NewEvent(ctx context.Context, config *config.Config, log *logrus.Logger, client grpc.GophkeeperClient) *Event {
	return &Event{context: ctx, config: config, logger: log, grpc: client}
}
