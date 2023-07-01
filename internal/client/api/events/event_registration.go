package events

import (
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventRegistration(username, password string) (model.Token, error) {
	c.logger.Info("Registration")

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
	createdToken, _ := service.ConvertTimestampToTime(registeredUser.AccessToken.CreatedAt)
	endDateToken, _ := service.ConvertTimestampToTime(registeredUser.AccessToken.EndDateAt)
	token = model.Token{AccessToken: registeredUser.AccessToken.Token, UserID: registeredUser.AccessToken.UserId,
		CreatedAt: createdToken, EndDateAt: endDateToken}

	return token, nil
}
