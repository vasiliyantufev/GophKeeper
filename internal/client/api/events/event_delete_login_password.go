package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventDeleteLoginPassword(loginPassword []string, token model.Token) error {
	c.logger.Info("Delete login password")

	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	deletedText, err := c.grpc.HandleDeleteLoginPassword(context.Background(),
		&grpc.DeleteLoginPasswordRequest{Name: loginPassword[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(loginPassword)
	c.logger.Debug(deletedText)
	return nil
}
