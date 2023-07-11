package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) FileUpload(name string, password string, file []byte, token model.Token) (string, error) {
	c.logger.Info("file upload")

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptFile, err := encryption.Encrypt(string(file), secretKey)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}

	createdToken, err := service.ConvertTimeToTimestamp(token.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	endDateToken, err := service.ConvertTimeToTimestamp(token.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	uploadFile, err := c.grpc.HandleUploadBinary(context.Background(),
		&grpc.UploadBinaryRequest{Name: name, Data: []byte(encryptFile),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	c.logger.Debug(uploadFile.Name)
	return uploadFile.Name, nil
}
