package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventCreateText(name, description, password, plaintext string, token model.Token) error {
	c.logger.Info("Create text")

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText, err := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	createdText, err := c.grpc.HandleCreateText(context.Background(),
		&grpc.CreateTextRequest{Name: name, Description: description, Text: []byte(encryptText),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	c.logger.Debug(createdText.Text)
	return nil
}
