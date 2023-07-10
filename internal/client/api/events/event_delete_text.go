package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) EventDeleteText(text []string, token model.Token) error {
	c.logger.Info("Delete text")

	createdToken, err := service.ConvertTimeToTimestamp(token.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	endDateToken, err := service.ConvertTimeToTimestamp(token.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	deletedText, err := c.grpc.HandleDeleteText(context.Background(),
		&grpc.DeleteTextRequest{Name: text[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	deletedTextEntityID, err := c.grpc.HandleDeleteEntity(context.Background(),
		&grpc.DeleteEntityRequest{Name: text[0], Type: variables.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(text)
	c.logger.Debug(deletedText)
	c.logger.Debug(deletedTextEntityID)
	return nil
}
