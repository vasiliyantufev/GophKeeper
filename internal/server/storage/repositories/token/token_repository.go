package token

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
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

func (t Token) Create(userID int64) (string, error) {
	token := encryption.GenerateAccessToken(lengthToken)
	currentTime := time.Now()

	var accessToken string
	return token, t.db.Pool.QueryRow(
		"INSERT INTO access_token (access_token, user_id, created_at, end_date_at) VALUES ($1, $2, $3, $4) RETURNING access_token",
		token,
		userID,
		currentTime,
		currentTime.Add(time.Hour+lifetimeToken),
	).Scan(&accessToken)
}

func (t *Token) Validate(token string) (bool, *model.Token, error) {
	currentTime := time.Now()

	tokenUser := &model.Token{}
	if err := t.db.Pool.QueryRow("SELECT access_token, user_id, end_date_at FROM access_token where access_token = $1", token).Scan(
		&tokenUser.AccessToken,
		&tokenUser.UserID,
		&tokenUser.EndDateAt,
	); err != nil {
		return false, nil, err
	}

	//timestamp.Timestamp to Time
	//timeValue := time.Unix(tokenUser.EndDateAt.GetSeconds(), int64(tokenUser.EndDateAt.GetNanos())).UTC()

	if currentTime.After(tokenUser.EndDateAt) {
		return false, tokenUser, nil
	}

	return true, tokenUser, nil
}
