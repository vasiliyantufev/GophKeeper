package api

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

type Client struct {
	grpc    grpc.GophkeeperClient
	logger  *logrus.Logger
	context context.Context
	grpc.UnimplementedGophkeeperServer
}

// NewClient - creates a new grpc client instance
func NewClient(ctx context.Context, log *logrus.Logger, client grpc.GophkeeperClient) *Client {
	return &Client{context: ctx, logger: log, grpc: client}
}

func (c Client) Ping() (string, error) {
	c.logger.Info("ping")
	msg, err := c.grpc.HandlePing(c.context, &grpc.PingRequest{})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	return msg.Message, nil
}

func (c Client) UserExist(username string) (bool, error) {
	c.logger.Info("user exist")
	user, err := c.grpc.HandleUserExist(c.context, &grpc.UserExistRequest{Username: username})
	if err != nil {
		c.logger.Error(err)
		return user.Exist, err
	}
	return user.Exist, nil
}

func (c Client) Registration(username, password string) (model.Token, error) {
	c.logger.Info("registration")
	token := model.Token{}
	password, err := encryption.HashPassword(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	registeredUser, err := c.grpc.HandleRegistration(c.context, &grpc.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	created, _ := service.ConvertTimestampToTime(registeredUser.AccessToken.CreatedAt)
	endDate, _ := service.ConvertTimestampToTime(registeredUser.AccessToken.EndDateAt)
	token = model.Token{AccessToken: registeredUser.AccessToken.Token, UserID: registeredUser.AccessToken.UserId,
		CreatedAt: created, EndDateAt: endDate}
	return token, nil
}

func (c Client) Authentication(username, password string) (model.Token, error) {
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

func (c Client) CreateText(key, value, password, plaintext string, token model.Token) error {
	c.logger.Info("create text")
	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText, err := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	createdText, err := c.grpc.HandleCreateText(context.Background(),
		&grpc.CreateTextRequest{Key: key, Value: value, Text: []byte(encryptText),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	c.logger.Debug(createdText.Text)
	return nil
}

//func (c Client) GetNodeText(key, value, password, accessToken string) (string, error) {
//	var plaintext string
//	secretKey := encryption.AesKeySecureRandom([]byte(password))
//	getNodeText, err := c.grpc.HandleGetNodeText(c.context, &gophkeeper.GetNodeTextRequest{Key: key, Value: value, AccessToken: accessToken})
//	if err != nil {
//		c.logger.Error(err)
//		return plaintext, err
//	}
//	plaintext = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
//	if err != nil {
//		c.logger.Error(err)
//		return plaintext, err
//	}
//	return plaintext, nil
//}
//
//func (c Client) GetListText(accessToken string) (*gophkeeper.GetListTextResponse, error) {
//	getListText, err := c.grpc.HandleGetListText(c.context, &gophkeeper.GetListTextRequest{AccessToken: accessToken})
//	if err != nil {
//		c.logger.Error(err)
//		return nil, err
//	}
//	return getListText, nil
//}

func (c Client) Sync(userId int64) ([][]string, [][]string) {
	dataTblText := [][]string{{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"},
		{"NAME_1", "DATA", "DESCRIPTION", "01/01/2000", "01/01/2000"}}
	dataTblCart := [][]string{{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED_AT", "UPDATED_AT"},
		{"NAME_1", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "01/01/2000", "01/01/2000", "01/01/2000"}}

	return dataTblText, dataTblCart
}
