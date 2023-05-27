package text

import (
	"database/sql"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/database"
	"github.com/vasiliyantufev/gophkeeper/internal/model"
	"github.com/vasiliyantufev/gophkeeper/internal/storage/errors"
)

type Text struct {
	db *database.DB
}

func New(db *database.DB) *Text {
	return &Text{
		db: db,
	}
}

func (t *Text) CreateText(textRequest *model.CreateTextRequest) (*model.Text, error) {
	text := &model.Text{}
	if err := t.db.Pool.QueryRow(
		"INSERT INTO text (user_id, metadata_id, text, updated_at, created_at) VALUES ($1, $2, $3, $4, $5) "+
			"RETURNING text_id, text",
		textRequest.UserID,
		textRequest.MetadataID,
		textRequest.Text,
		time.Now(),
		time.Now(),
	).Scan(&text.ID, &text.Text); err != nil {
		return nil, err
	}
	return text, nil
}

func (t *Text) GetNodeText(textRequest *model.GetNodeTextRequest) (*model.GetNodeTextResponse, error) {
	text := &model.GetNodeTextResponse{}
	err := t.db.Pool.QueryRow("SELECT metadata.name, text.text FROM metadata "+
		"inner join text on metadata.metadata_id = text.metadata_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where metadata.name = $1",
		textRequest.Name).Scan(
		&text.Name,
		&text.Text,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return text, nil
}

func (t *Text) NameExists(name string) (bool, error) {
	var exists bool
	row := t.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM metadata "+
		"inner join text on metadata.metadata_id = text.metadata_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where metadata.name = $1)", name)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}
