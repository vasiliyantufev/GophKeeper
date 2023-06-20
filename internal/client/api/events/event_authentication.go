package events

import (
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) Authentication(username, password string) (model.Token, error) {
	c.logger.Info("authentication")
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
	created, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.CreatedAt)
	endDate, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.EndDateAt)
	token = model.Token{AccessToken: authenticatedUser.AccessToken.Token, UserID: authenticatedUser.AccessToken.UserId,
		CreatedAt: created, EndDateAt: endDate}
	return token, nil
}
