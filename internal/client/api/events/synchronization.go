package events

import (
	"encoding/json"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/table"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/labels"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) Synchronization(password string, token model.Token) ([][]string, [][]string, [][]string, [][]string, error) {
	c.logger.Info("synchronization")

	dataTblText := [][]string{}
	dataTblCard := [][]string{}
	dataTblLoginPassword := [][]string{}
	dataTblBinary := [][]string{}

	created, err := service.ConvertTimeToTimestamp(token.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}
	endDate, err := service.ConvertTimeToTimestamp(token.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}
	//-----------------------------------------------
	var plaintext string
	secretKey := encryption.AesKeySecureRandom([]byte(password))

	titleText := []string{labels.NameItem, labels.DescriptionItem, labels.DataItem, labels.CreatedAtItem, labels.UpdatedAtItem}
	titleCard := []string{labels.NameItem, labels.DescriptionItem, labels.PaymentSystemItem, labels.NumberItem, labels.HolderItem,
		labels.CVCItem, labels.EndDateItem, labels.CreatedAtItem, labels.UpdatedAtItem}
	titleLoginPassword := []string{labels.NameItem, labels.DescriptionItem, labels.LoginItem, labels.PasswordItem,
		labels.CreatedAtItem, labels.UpdatedAtItem}
	titleBinary := []string{labels.NameItem, labels.CreatedAtItem}

	dataTblText = append(dataTblText, titleText)
	dataTblCard = append(dataTblCard, titleCard)
	dataTblLoginPassword = append(dataTblLoginPassword, titleLoginPassword)
	dataTblBinary = append(dataTblBinary, titleBinary)

	dataTblTextPointer := &dataTblText
	dataTblCardPointer := &dataTblCard
	dataTblLoginPasswordPointer := &dataTblLoginPassword
	dataTblBinaryPointer := &dataTblBinary

	//-----------------------------------------------
	nodesTextEntity, err := c.grpc.EntityGetList(c.context,
		&grpc.GetListEntityRequest{Type: variables.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesTextEntity.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = table.AppendTextEntity(node, dataTblTextPointer, plaintext)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}

	//-----------------------------------------------
	nodesCardEntity, err := c.grpc.EntityGetList(c.context,
		&grpc.GetListEntityRequest{Type: variables.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesCardEntity.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}

		var card model.Card
		err = json.Unmarshal([]byte(plaintext), &card)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = table.AppendCardEntity(node, dataTblCardPointer, card)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	nodesLoginPasswordEntity, err := c.grpc.EntityGetList(c.context,
		&grpc.GetListEntityRequest{Type: variables.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesLoginPasswordEntity.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}

		var loginPassword model.LoginPassword
		err = json.Unmarshal([]byte(plaintext), &loginPassword)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = table.AppendLoginPasswordEntity(node, dataTblLoginPasswordPointer, loginPassword)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	nodesBinary, err := c.grpc.FileGetList(c.context,
		&grpc.GetListBinaryRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesBinary.Node {
		err = table.AppendBinary(node, dataTblBinaryPointer)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, nil
}
