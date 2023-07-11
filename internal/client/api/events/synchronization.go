package events

import (
	"encoding/json"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/table"
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

	titleText := []string{"NAME", "DESCRIPTION", "DATA", "CREATED AT", "UPDATED AT"}
	titleCard := []string{"NAME", "DESCRIPTION", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED AT", "UPDATED AT"}
	titleLoginPassword := []string{"NAME", "DESCRIPTION", "LOGIN", "PASSWORD", "CREATED AT", "UPDATED AT"}
	titleBinary := []string{"NAME", "CREATED AT"}

	dataTblText = append(dataTblText, titleText)
	dataTblCard = append(dataTblCard, titleCard)
	dataTblLoginPassword = append(dataTblLoginPassword, titleLoginPassword)
	dataTblBinary = append(dataTblBinary, titleBinary)

	dataTblTextPointer := &dataTblText
	dataTblCardPointer := &dataTblCard
	dataTblLoginPasswordPointer := &dataTblLoginPassword
	dataTblBinaryPointer := &dataTblBinary

	//-----------------------------------------------
	nodesTextEntity, err := c.grpc.HandleGetListEntity(c.context,
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
	nodesCardEntity, err := c.grpc.HandleGetListEntity(c.context,
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
	nodesLoginPasswordEntity, err := c.grpc.HandleGetListEntity(c.context,
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
	nodesBinary, err := c.grpc.HandleGetListBinary(c.context,
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
