package api

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
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

func (c Client) Registration(username, password string) (model.Token, error) {
	token := model.Token{}
	password, err := encryption.HashPassword2(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	registeredUser, err := c.grpc.HandleRegistration(c.context, &gophkeeper.RegistrationRequest{Username: username, Password: password})
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
	token := model.Token{}
	password, err := encryption.HashPassword2(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	authenticatedUser, err := c.grpc.HandleAuthentication(c.context, &gophkeeper.AuthenticationRequest{Username: username, Password: password})
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

//func (c Client) CreateText(key, value, password, text string, token model.Token) error {
//	secretKey := encryption.AesKeySecureRandom([]byte(password))
//	encryptText := encryption.Encrypt(text, secretKey)
//	_, err := c.grpc.HandleCreateText(c.context, &gophkeeper.CreateTextRequest{Key: key, Value: value, Text: []byte(encryptText), AccessToken: token})
//	if err != nil {
//		c.logger.Error(err)
//		return err
//	}
//	return nil
//}
//
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
