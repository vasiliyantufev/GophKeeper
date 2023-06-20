package events

import (
	"context"

	"github.com/sirupsen/logrus"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

type Event struct {
	grpc    grpc.GophkeeperClient
	logger  *logrus.Logger
	context context.Context
	grpc.UnimplementedGophkeeperServer
}

// NewEvent - creates a new grpc client instance
func NewEvent(ctx context.Context, log *logrus.Logger, client grpc.GophkeeperClient) *Event {
	return &Event{context: ctx, logger: log, grpc: client}
}

//	func (c Event) GetNodeText(key, value, password, accessToken string) (string, error) {
//		var plaintext string
//		secretKey := encryption.AesKeySecureRandom([]byte(password))
//		getNodeText, err := c.grpc.HandleGetNodeText(c.context, &gophkeeper.GetNodeTextRequest{Key: key, Value: value, AccessToken: accessToken})
//		if err != nil {
//			c.logger.Error(err)
//			return plaintext, err
//		}
//		plaintext = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
//		if err != nil {
//			c.logger.Error(err)
//			return plaintext, err
//		}
//		return plaintext, nil
//	}
