package events

import (
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

func (c Event) EventPing() (string, error) {
	c.logger.Info("Ping")
	msg, err := c.grpc.HandlePing(c.context, &grpc.PingRequest{})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}

	return msg.Message, nil
}
