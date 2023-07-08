package events

import (
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventAuthentication(username, password string) (model.Token, error) {
	c.logger.Info("Authentication")

	token := model.Token{}
	password, err := encryption.HashPassword(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	authenticatedUser, err := c.grpc.HandleAuthentication(c.context, &grpc.AuthenticationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return token, err
	}

	createdToken, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.CreatedAt)
	endDateToken, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.EndDateAt)
	token = model.Token{AccessToken: authenticatedUser.AccessToken.Token, UserID: authenticatedUser.AccessToken.UserId,
		CreatedAt: createdToken, EndDateAt: endDateToken}

	err = service.CreateStorageNotExistsUser(c.config.FileFolder, token.UserID)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}

	return token, nil
}
