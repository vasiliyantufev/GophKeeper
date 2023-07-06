package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/card"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/login_password"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/metadata"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/text"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
)

type Handler struct {
	database      *database.DB
	config        *config.Config
	user          *user.User
	text          *text.Text
	card          *card.Card
	loginPassword *loginPassword.LoginPassword
	metadata      *metadata.Metadata
	storage       storage.Storage
	token         *token.Token
	logger        *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, config *config.Config, userRepository *user.User, textRepository *text.Text, cardRepository *card.Card,
	loginPasswordRepository *loginPassword.LoginPassword, metadataRepository *metadata.Metadata, storage storage.Storage,
	tokenRepository *token.Token, log *logrus.Logger) *Handler {
	return &Handler{database: db, config: config, user: userRepository, text: textRepository, card: cardRepository,
		loginPassword: loginPasswordRepository, metadata: metadataRepository, storage: storage, token: tokenRepository, logger: log}
}
