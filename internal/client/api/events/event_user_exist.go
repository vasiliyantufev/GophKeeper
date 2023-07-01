package events

import grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"

func (c Event) EventUserExist(username string) (bool, error) {
	c.logger.Info("User exist")

	user, err := c.grpc.HandleUserExist(c.context, &grpc.UserExistRequest{Username: username})
	if err != nil {
		c.logger.Error(err)
		return user.Exist, err
	}

	return user.Exist, nil
}
