package events

import (
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
)

func (c Event) EventDeleteCard(card []string, token model.Token) error {
	c.logger.Info("Delete card")

	return nil
}
