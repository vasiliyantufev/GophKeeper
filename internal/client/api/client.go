package api

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	gophkeeper "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

type Client struct {
	logger *logrus.Logger
	gophkeeper.UnimplementedGophkeeperServer
}

// NewClient - creates a new grpc client instance
func NewClient(log *logrus.Logger) *Client {
	return &Client{logger: log}
}

func Registration(ctx context.Context, username, password string) {
	registeredUser, err := client.HandleRegistration(context.Background(), &gophkeeper.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	return registeredUser
}

func Authentication(ctx context.Context, username, password string) {
	authenticatedUser, err := client.HandleAuthentication(context.Background(),
		&gophkeeper.AuthenticationRequest{Username: registeredUser.User.Username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{ID: authenticatedUser.User.UserId, Username: authenticatedUser.User.Username}
	return user
}

func CreateText(ctx context.Context, key, value, text, accessToken string) {
	createdText, err := client.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Key: "Username", Value: randName, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	return createdText
}

func GetNodeText(ctx context.Context, key, value, accessToken string) {
	getNodeText, err := client.HandleGetNodeText(context.Background(), &gophkeeper.GetNodeTextRequest{Key: "Username", Value: randName, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	return plaintext
}

func GetListText(ctx context.Context, accessToken string) {
	getListText, err := client.HandleGetListText(context.Background(), &gophkeeper.GetListTextRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	return getListText
}
