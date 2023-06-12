package api

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

type Client struct {
	grpc   gophkeeper.GophkeeperClient
	Logger *logrus.Logger
	gophkeeper.UnimplementedGophkeeperServer
}

// NewClient - creates a new grpc client instance
func NewClient(log *logrus.Logger, client gophkeeper.GophkeeperClient) *Client {
	return &Client{Logger: log, grpc: client}
}

func (c Client) Ping(ctx context.Context, username, password string) (string, error) {
	msg, err := c.grpc.HandlePing(context.Background(), &gophkeeper.PingRequest{})
	if err != nil {
		c.Logger.Error(err)
		return msg.Message, err
	}
	return msg.Message, nil
}

func (c Client) Registration(ctx context.Context, username, password string) (model.User, error) {
	user := model.User{}
	registeredUser, err := c.grpc.HandleRegistration(context.Background(), &gophkeeper.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		c.Logger.Error(err)
		return user, err
	}
	user = model.User{ID: registeredUser.User.UserId, Username: registeredUser.User.Username, AccessToken: registeredUser.AccessToken}
	return user, nil
}

func (c Client) Authentication(ctx context.Context, username, password string) (model.User, error) {
	user := model.User{}
	authenticatedUser, err := c.grpc.HandleAuthentication(context.Background(),
		&gophkeeper.AuthenticationRequest{Username: username, Password: password})
	if err != nil {
		c.Logger.Error(err)
		return user, err
	}
	user = model.User{ID: authenticatedUser.User.UserId, Username: authenticatedUser.User.Username, AccessToken: authenticatedUser.AccessToken}
	return user, nil
}

func (c Client) CreateText(ctx context.Context, key, value, password, text, accessToken string) error {
	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText := encryption.Encrypt(text, secretKey)
	_, err := c.grpc.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Key: key, Value: value, Text: []byte(encryptText), AccessToken: accessToken})
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	return nil
}

func (c Client) GetNodeText(ctx context.Context, key, value, password, accessToken string) (string, error) {
	var plaintext string
	secretKey := encryption.AesKeySecureRandom([]byte(password))
	getNodeText, err := c.grpc.HandleGetNodeText(context.Background(), &gophkeeper.GetNodeTextRequest{Key: key, Value: value, AccessToken: accessToken})
	if err != nil {
		c.Logger.Error(err)
		return plaintext, err
	}
	plaintext = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
	if err != nil {
		c.Logger.Error(err)
		return plaintext, err
	}
	return plaintext, nil
}

func (c Client) GetListText(ctx context.Context, accessToken string) (*gophkeeper.GetListTextResponse, error) {
	getListText, err := c.grpc.HandleGetListText(context.Background(), &gophkeeper.GetListTextRequest{AccessToken: accessToken})
	if err != nil {
		c.Logger.Error(err)
		return getListText, err
	}
	return getListText, nil
}
