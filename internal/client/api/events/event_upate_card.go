package events

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/variables"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventUpdateCard(name, passwordSecure, paymentSystem, number, holder, cvc, endDateCard string, token model.Token) error {
	c.logger.Info("Update Card")

	intCvc, err := strconv.Atoi(cvc)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	timeEndDate, err := time.Parse(string(variables.LayoutDate), endDateCard)
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

	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	updateCard, err := c.grpc.HandleUpdateCard(context.Background(), &grpc.UpdateCardRequest{Name: name, Data: []byte(encryptCard),
		AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(updateCard)
	return nil
}
