package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventDownload(name string, password string, token model.Token) error {
	c.logger.Info("Download binary data")

	//secretKey := encryption.AesKeySecureRandom([]byte(password))
	//encryptFile, err := encryption.Encrypt(string(file), secretKey)
	//if err != nil {
	//	c.logger.Error(err)
	//	return err, nil
	//}
	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)

	downloadFile, err := c.grpc.HandleDownloadBinary(context.Background(),
		&grpc.DownloadBinaryRequest{Name: name, AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
			CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	err = service.UploadFile(c.config.FileFolder, token.UserID, name, downloadFile.Data)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(name)
	return nil
}
