package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/database"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

type Handler struct {
	database *database.DB
	logger   *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, log *logrus.Logger) *Handler {
	return &Handler{database: db, logger: log}
}
