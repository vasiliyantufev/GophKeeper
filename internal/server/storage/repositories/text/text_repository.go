package text

import (
	"database/sql"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
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
		"INSERT INTO text (user_id, data, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING text_id, data",
		textRequest.UserID,
		textRequest.Data,
		time.Now(),
		time.Now(),
	).Scan(&text.ID, &text.Data); err != nil {
		return nil, err
	}
	return text, nil
}

func (t *Text) GetNodeText(textRequest *model.GetNodeTextRequest) (*model.Text, error) {
	text := &model.Text{}
	err := t.db.Pool.QueryRow("SELECT text.data FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4 and text.deleted_at IS NULL",
		string(variables.Name), textRequest.Value, textRequest.UserID, string(variables.Text)).
		Scan(&text.Data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return text, nil
}

func (t *Text) GetListText(userId int64) ([]model.Text, error) {
	listText := []model.Text{}

	rows, err := t.db.Pool.Query("SELECT metadata.entity_id, metadata.key, text.data, metadata.value, text.created_at, "+
		"text.updated_at FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where users.user_id = $1 and metadata.type = $2 and text.deleted_at IS NULL",
		userId, string(variables.Text))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		text := model.Text{}
		err = rows.Scan(&text.ID, &text.Key, &text.Data, &text.Value, &text.CreatedAt, &text.UpdatedAt)
		if err != nil {
			return nil, err
		}
		listText = append(listText, text)
	}
	return listText, nil
}

func (t *Text) GetIdText(value string, userID int64) (int64, error) {
	var textID int64
	err := t.db.Pool.QueryRow("SELECT text.text_id FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4 and text.deleted_at IS NULL",
		string(variables.Name), value, userID, string(variables.Text)).
		Scan(&textID)
	if err != nil {
		if err == sql.ErrNoRows {
			return textID, errors.ErrRecordNotFound
		} else {
			return textID, err
		}
	}
	return textID, nil
}

func (t *Text) KeyExists(textRequest *model.CreateTextRequest) (bool, error) {
	var exists bool
	row := t.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4 and text.deleted_at IS NULL)",
		string(variables.Name), textRequest.Name, textRequest.UserID, string(variables.Text))
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (t *Text) DeleteText(textID int64) error {
	var id int64
	layout := "01/02/2006 15:04:05"
	if err := t.db.Pool.QueryRow("UPDATE text SET deleted_at = $1 WHERE text_id = $2 RETURNING text_id",
		time.Now().Format(layout),
		textID,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (t *Text) UpdateText(textID int64, data []byte) error {
	var id int64
	layout := "01/02/2006 15:04:05"
	if err := t.db.Pool.QueryRow("UPDATE text SET data = $1, updated_at = $2 WHERE text_id = $3 RETURNING text_id",
		data,
		time.Now().Format(layout),
		textID,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}
