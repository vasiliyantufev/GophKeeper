package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/database"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/storage/repositories/metadata"
	"github.com/vasiliyantufev/gophkeeper/internal/storage/repositories/text"
	"github.com/vasiliyantufev/gophkeeper/internal/storage/repositories/user"
)

type Handler struct {
	database *database.DB
	user     *user.User
	text     *text.Text
	metadata *metadata.Metadata
	logger   *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, userRepository *user.User, textRepository *text.Text, metadataRepository *metadata.Metadata, log *logrus.Logger) *Handler {
	return &Handler{database: db, user: userRepository, text: textRepository, metadata: metadataRepository, logger: log}
}
