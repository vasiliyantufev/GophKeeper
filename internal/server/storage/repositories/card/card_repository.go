package card

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
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
