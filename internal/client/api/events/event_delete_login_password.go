package events

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventDeleteLoginPassword(loginPassword []string, token model.Token) error {
	c.logger.Info("Delete login password")

	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	deletedText, err := c.grpc.HandleDeleteLoginPassword(context.Background(),
		&grpc.DeleteLoginPasswordRequest{Name: loginPassword[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	logrus.Info(loginPassword)
	c.logger.Debug(deletedText)
	return nil
}
