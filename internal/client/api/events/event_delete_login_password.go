package events

import (
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
)

func (c Event) EventDeleteLoginPassword(token model.Token) error {
	c.logger.Info("Delete login password")

	return nil
}
