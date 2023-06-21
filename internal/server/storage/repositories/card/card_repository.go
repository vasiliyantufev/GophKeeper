package card

import (
	"database/sql"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

type Card struct {
	db *database.DB
}

func New(db *database.DB) *Card {
	return &Card{
		db: db,
	}
}

func (c *Card) CreateCard(cardRequest *model.CreateCardRequest) (*model.Card, error) {
	card := &model.Card{}
	if err := c.db.Pool.QueryRow(
		"INSERT INTO card (user_id, card_data, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING card_id, card_data",
		cardRequest.UserID,
		cardRequest.CardData,
		time.Now(),
		time.Now(),
	).Scan(&card.ID, &card.CardData); err != nil {
		return nil, err
	}
	return card, nil
}

func (t *Card) GetListText(userId int64) ([]model.Card, error) {
	ListCard := []model.Card{}
	return ListCard, nil
}

func (c *Card) KeyExists(cardRequest *model.CreateCardRequest) (bool, error) {
	var exists bool
	row := c.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM metadata "+
		"inner join card on metadata.entity_id = card.card_id "+
		"inner join users on card.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3)", string(variables.Name), cardRequest.Name, cardRequest.UserID)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (c *Card) GetNodeCard(cardRequest *model.GetNodeCardRequest) (*model.Card, error) {
	card := &model.Card{}
	err := c.db.Pool.QueryRow("SELECT card.card_data FROM metadata "+
		"inner join card on metadata.entity_id = card.card_id "+
		"inner join users on card.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3",
		string(variables.Name), cardRequest.Value, cardRequest.UserID).Scan(
		&card.CardData,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return card, nil
}

func (c *Card) GetListCard(userId int64) ([]model.Card, error) {
	ListCard := []model.Card{}

	rows, err := c.db.Pool.Query("SELECT metadata.entity_id, metadata.key, text.text, metadata.value, text.created_at, "+
		"text.updated_at FROM metadata "+
		//rows, err := t.db.Pool.Query("SELECT text.text FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where users.user_id = $1", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		card := model.Card{}
		err = rows.Scan(&card.ID, &card.Key, &card.CardData, &card.Value, &card.CreatedAt, &card.UpdatedAt)
		if err != nil {
			return nil, err
		}
		ListCard = append(ListCard, card)
	}
	return ListCard, nil
}
