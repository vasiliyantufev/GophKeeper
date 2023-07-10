package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) EventUpdateText(name, passwordSecure, text string, token model.Token) error {
	c.logger.Info("Update text")

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptText, err := encryption.Encrypt(text, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

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
	updateText, err := c.grpc.HandleUpdateText(context.Background(), &grpc.UpdateTextRequest{Name: name, Data: []byte(encryptText),
		AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	updatedTextEntity, err := c.grpc.HandleUpdateEntity(context.Background(),
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptText), Type: variables.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(updateText)
	c.logger.Debug(updatedTextEntity)
	return nil
}
