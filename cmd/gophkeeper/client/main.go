package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/config/configagent"
	"github.com/vasiliyantufev/gophkeeper/internal/model"
	"github.com/vasiliyantufev/gophkeeper/internal/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/service/randomizer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcClient "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

func main() {
	log := logrus.New()
	config := configagent.NewConfigClient(log)
	log.SetLevel(config.DebugLevel)

	conn, err := grpc.Dial(config.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := grpcClient.NewGophkeeperClient(conn)
	resp, err := client.HandlePing(context.Background(), &grpcClient.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(resp.Message)

	username := randomizer.RandStringRunes(10)
	password := "Пароль-1"
	registeredUser, err := client.HandleRegistration(context.Background(), &grpcClient.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	authenticatedUser, err := client.HandleAuthentication(context.Background(), &grpcClient.AuthenticationRequest{Username: registeredUser.Username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{ID: authenticatedUser.UserId, Username: authenticatedUser.Username}
	log.Info(user)

	randName := randomizer.RandStringRunes(10)
	plaintext := "Hi my sweetly friends!!!"

	secretKey := encryption.AesKeySecureRandom(password)

	encryptText := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdText, err := client.HandleCreateText(context.Background(), &grpcClient.CreateTextRequest{UserId: user.ID, Name: randName, Text: []byte(encryptText)})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText.Text)

	getNodeText, err := client.HandleGetNodeText(context.Background(), &grpcClient.GetNodeTextRequest{Name: createdText.Name})
	if err != nil {
		log.Fatal(err)
	}
	plaintext = encryption.Decrypt(string(getNodeText.Text), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)
}
