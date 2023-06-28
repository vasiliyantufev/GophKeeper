package events

import (
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
)

func (c Event) EventDeleteLoginPassword(row []string, token model.Token) error {
	c.logger.Info("Delete login password")

	logrus.Info(row)

	return nil
}
