package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/config"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/randomizer"
	"github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log := logrus.New()
	config := config.NewConfig(log)
	log.SetLevel(config.DebugLevel)

	conn, err := grpc.Dial(config.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := gophkeeper.NewGophkeeperClient(conn)
	resp, err := client.HandlePing(context.Background(), &gophkeeper.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(resp.Message)

	username := randomizer.RandStringRunes(10)
	password := "Пароль-1"

	password, err = encryption.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	registeredUser, err := client.HandleRegistration(context.Background(), &gophkeeper.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	authenticatedUser, err := client.HandleAuthentication(context.Background(), &gophkeeper.AuthenticationRequest{Username: registeredUser.User.Username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{ID: authenticatedUser.User.UserId, Username: authenticatedUser.User.Username}
	log.Info(user)

	randName := randomizer.RandStringRunes(10)
	plaintext := "Hi my sweetly friends!!!!!!!TeST ВСЕМПРИВЕТ!"

	secretKey := encryption.AesKeySecureRandom([]byte(password))

	encryptText := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdText, err := client.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Key: "Username", Value: randName, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText.Text)

	getNodeText, err := client.HandleGetNodeText(context.Background(), &gophkeeper.GetNodeTextRequest{Key: "Username", Value: randName, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)

	createdText2, err := client.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Key: "Name2", Value: randName, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText2.Text)

	getListText, err := client.HandleGetListText(context.Background(), &gophkeeper.GetListTextRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListText)
}
