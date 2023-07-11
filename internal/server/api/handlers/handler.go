package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/binary"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/entity"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
)

type Handler struct {
	database *database.DB
	config   *config.Config
	user     *user.User
	binary   *binary.Binary
	storage  *storage.Storage
	entity   *entity.Entity
	token    *token.Token
	logger   *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, config *config.Config, userRepository *user.User,
	binaryRepository *binary.Binary, storage *storage.Storage, entityRepository *entity.Entity, tokenRepository *token.Token, log *logrus.Logger) *Handler {
	return &Handler{database: db, config: config, user: userRepository, binary: binaryRepository, storage: storage,
		entity: entityRepository, token: tokenRepository, logger: log}
}
