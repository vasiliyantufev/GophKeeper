package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventUpdateText(name, passwordSecure, text string, token model.Token) error {
	c.logger.Info("Update text")

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptText, err := encryption.Encrypt(string(text), secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	updateText, err := c.grpc.HandleUpdateText(context.Background(), &grpc.UpdateTextRequest{Name: name, Data: []byte(encryptText),
		AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	c.logger.Debug(updateText)
	return nil
}
