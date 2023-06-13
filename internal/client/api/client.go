package api

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

type Client struct {
	grpc    gophkeeper.GophkeeperClient
	logger  *logrus.Logger
	context context.Context
	gophkeeper.UnimplementedGophkeeperServer
}

// NewClient - creates a new grpc client instance
func NewClient(ctx context.Context, log *logrus.Logger, client gophkeeper.GophkeeperClient) *Client {
	return &Client{context: ctx, logger: log, grpc: client}
}

func (c Client) Ping() (string, error) {
	msg, err := c.grpc.HandlePing(c.context, &gophkeeper.PingRequest{})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	return msg.Message, nil
}

func (c Client) UserExist(username string) (bool, error) {
	user, err := c.grpc.HandleUserExist(c.context, &gophkeeper.UserExistRequest{Username: username})
	if err != nil {
		c.logger.Error(err)
		return user.Exist, err
	}
	return user.Exist, nil
}

func (c Client) Registration(username, password string) (model.User, error) {
	user := model.User{}
	registeredUser, err := c.grpc.HandleRegistration(c.context, &gophkeeper.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return user, err
	}
	user = model.User{ID: registeredUser.User.UserId, Username: registeredUser.User.Username, AccessToken: registeredUser.AccessToken}
	return user, nil
}

func (c Client) Authentication(username, password string) (model.User, error) {
	user := model.User{}
	authenticatedUser, err := c.grpc.HandleAuthentication(c.context, &gophkeeper.AuthenticationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return user, err
	}
	user = model.User{ID: authenticatedUser.User.UserId, Username: authenticatedUser.User.Username, AccessToken: authenticatedUser.AccessToken}
	return user, nil
}

func (c Client) CreateText(key, value, password, text, accessToken string) error {
	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText := encryption.Encrypt(text, secretKey)
	_, err := c.grpc.HandleCreateText(c.context, &gophkeeper.CreateTextRequest{Key: key, Value: value, Text: []byte(encryptText), AccessToken: accessToken})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	return nil
}

func (c Client) GetNodeText(key, value, password, accessToken string) (string, error) {
	var plaintext string
	secretKey := encryption.AesKeySecureRandom([]byte(password))
	getNodeText, err := c.grpc.HandleGetNodeText(c.context, &gophkeeper.GetNodeTextRequest{Key: key, Value: value, AccessToken: accessToken})
	if err != nil {
		c.logger.Error(err)
		return plaintext, err
	}
	plaintext = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
	if err != nil {
		c.logger.Error(err)
		return plaintext, err
	}
	return plaintext, nil
}

func (c Client) GetListText(accessToken string) (*gophkeeper.GetListTextResponse, error) {
	getListText, err := c.grpc.HandleGetListText(c.context, &gophkeeper.GetListTextRequest{AccessToken: accessToken})
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return getListText, nil
}
