package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventDeleteText(text []string, token model.Token) error {
	c.logger.Info("Delete text")

	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	deletedText, err := c.grpc.HandleDeleteText(context.Background(),
		&grpc.DeleteTextRequest{Name: text[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	c.logger.Debug(text)
	c.logger.Debug(deletedText)
	return nil
}
