package events

import (
	"context"
	"encoding/json"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventCreateLoginPassword(name, description, passwordSecure, login, password string, token model.Token) error {
	c.logger.Info("Create text")

	loginPassword := model.LoginPassword{Login: login, Password: password}
	jsonLoginPassword, err := json.Marshal(loginPassword)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptLoginPassword, err := encryption.Encrypt(string(jsonLoginPassword), secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	createdLoginPassword, err := c.grpc.HandleCreateLoginPassword(context.Background(),
		&grpc.CreateLoginPasswordRequest{Name: name, Description: description, Data: []byte(encryptLoginPassword),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	c.logger.Debug(createdLoginPassword.Data)
	return nil
}
