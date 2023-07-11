package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) EventDeleteLoginPassword(loginPassword []string, token model.Token) error {
	c.logger.Info("Delete login password")

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

	//deletedLoginPassword, err := c.grpc.HandleDeleteLoginPassword(context.Background(),
	//	&grpc.DeleteLoginPasswordRequest{Name: loginPassword[0],
	//		AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	//if err != nil {
	//	c.logger.Error(err)
	//	return err
	//}

	deletedLoginPasswordEntityID, err := c.grpc.HandleDeleteEntity(context.Background(),
		&grpc.DeleteEntityRequest{Name: loginPassword[0], Type: variables.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	//c.logger.Debug(deletedLoginPassword.Id)
	c.logger.Debug(deletedLoginPasswordEntityID)
	return nil
}
