package events

import (
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) Registration(username, password string) (model.Token, error) {
	c.logger.Info("registration")

	token := model.Token{}
	password, err := encryption.HashPassword(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	registeredUser, err := c.grpc.HandleRegistration(c.context, &grpc.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	createdToken, err := service.ConvertTimestampToTime(registeredUser.AccessToken.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	endDateToken, err := service.ConvertTimestampToTime(registeredUser.AccessToken.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	token = model.Token{AccessToken: registeredUser.AccessToken.Token, UserID: registeredUser.AccessToken.UserId,
		CreatedAt: createdToken, EndDateAt: endDateToken}

	err = service.CreateStorageUser(c.config.FileFolder, token.UserID)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}

	return token, nil
}
