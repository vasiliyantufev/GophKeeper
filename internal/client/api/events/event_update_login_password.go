package events

import (
	"context"
	"encoding/json"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) EventUpdateLoginPassword(name, passwordSecure, login, password string, token model.Token) error {
	c.logger.Info("Update login password")

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

	//updateLoginPassword, err := c.grpc.HandleUpdateLoginPassword(context.Background(), &grpc.UpdateLoginPasswordRequest{Name: name, Data: []byte(encryptLoginPassword),
	//	AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	//if err != nil {
	//	c.logger.Error(err)
	//	return err
	//}

	updatedLoginPasswordEntityID, err := c.grpc.HandleUpdateEntity(context.Background(),
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptLoginPassword), Type: variables.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	//c.logger.Debug(updateLoginPassword)
	c.logger.Debug(updatedLoginPasswordEntityID)
	return nil
}
