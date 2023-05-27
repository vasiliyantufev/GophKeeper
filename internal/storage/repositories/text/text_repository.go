package text

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/database"
	"github.com/vasiliyantufev/gophkeeper/internal/model"
)

type Text struct {
	db *database.DB
}

func New(db *database.DB) *Text {
	return &Text{
		db: db,
	}
}

func (t *Text) CreateText(text *model.CreateTextRequest) (*model.Text, error) {
	textCreated := &model.Text{}
	if err := t.db.Pool.QueryRow(
		"INSERT INTO text (user_id, metadata_id, text, updated_at, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING text_id, text",
		text.UserID,
		1, //metadata_id - надо изменить
		text.Text,
		time.Now(),
		time.Now(),
	).Scan(&textCreated.ID, &textCreated.Text); err != nil {
		return nil, err
	}
	return textCreated, nil
}
