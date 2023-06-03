package token

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/database"
	"github.com/vasiliyantufev/gophkeeper/internal/model"
	"github.com/vasiliyantufev/gophkeeper/internal/service/encryption"
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

func (t Token) Create(userID int) (string, error) {
	token := encryption.GenerateAccessToken(lengthToken)
	currentTime := time.Now()

	var accessToken string
	return token, t.db.Pool.QueryRow(
		"INSERT INTO token (access_token, user_id, created_at, end_date_at) VALUES ($1, $2, $3, $4) RETURNING access_token",
		userID,
		token,
		currentTime,
		currentTime.Add(time.Hour+lifetimeToken),
	).Scan(&accessToken)
}
