package token

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

const lengthToken = 32
const lifetimeToken = 100 * time.Hour

type TokenRepository interface {
	Create(user *model.User) (string, error)
}

type Token struct {
	db *database.DB
}

func New(db *database.DB) *Token {
	return &Token{
		db: db,
	}
}

func (t Token) Create(userID int64) (*model.Token, error) {
	token := &model.Token{}
	accessToken := encryption.GenerateAccessToken(lengthToken)
	currentTime := time.Now()

	if err := t.db.Pool.QueryRow(
		"INSERT INTO access_token (access_token, user_id, created_at, end_date_at) VALUES ($1, $2, $3, $4) RETURNING access_token, user_id, created_at, end_date_at",
		accessToken,
		userID,
		currentTime,
		currentTime.Add(time.Hour+lifetimeToken),
	).Scan(&token.AccessToken, &token.UserID, &token.CreatedAt, &token.EndDateAt); err != nil {
		return nil, err
	}
	return token, nil
}

func (t *Token) Validate(token *grpc.Token) bool {
	currentTime := time.Now()
	endDate, _ := service.ConvertTimestampToTime(token.EndDateAt)
	if currentTime.After(endDate) {
		return false
	}
	return true
}
