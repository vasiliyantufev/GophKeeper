package events

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/layouts"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) EventUpdateCard(name, passwordSecure, paymentSystem, number, holder, cvc, endDateCard string, token model.Token) error {
	c.logger.Info("Update Card")

	intCvc, err := strconv.Atoi(cvc)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	timeEndDate, err := time.Parse(layouts.LayoutDate.ToString(), endDateCard)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	card := model.Card{PaymentSystem: paymentSystem, Number: number, Holder: holder, CVC: intCvc, EndDate: timeEndDate}
	jsonCard, err := json.Marshal(card)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptCard, err := encryption.Encrypt(string(jsonCard), secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	createdToken, err := service.ConvertTimeToTimestamp(token.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	endDateToken, err := service.ConvertTimeToTimestamp(token.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	//updateCard, err := c.grpc.HandleUpdateCard(context.Background(), &grpc.UpdateCardRequest{Name: name, Data: []byte(encryptCard),
	//	AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	//if err != nil {
	//	c.logger.Error(err)
	//	return err
	//}

	updatedCardEntityID, err := c.grpc.HandleUpdateEntity(context.Background(),
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptCard), Type: variables.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	//c.logger.Debug(updateCard)
	c.logger.Debug(updatedCardEntityID)
	return nil
}
