package events

import (
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
)

func (c Event) EventDeleteText(token model.Token) error {
	c.logger.Info("Delete text")

	return nil
}
