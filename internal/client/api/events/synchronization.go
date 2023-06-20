package events

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/table"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) Synchronization(password string, token model.Token) ([][]string, [][]string, error) {
	dataTblText := [][]string{}
	dataTblCard := [][]string{}
	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	nodes, err := c.grpc.HandleGetListText(c.context,
		&grpc.GetListTextRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, err
	}

	var plaintext string

	titleText := []string{"ID", "NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"}
	titleCart := []string{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC",
		"END DATE", "CREATED_AT", "UPDATED_AT"}
	dataTblText = append(dataTblText, titleText)
	dataTblCard = append(dataTblCard, titleCart)
	dataTblTextPointer := &dataTblText

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	for _, node := range nodes.Node {
		plaintext, err = encryption.Decrypt(string(node.Text), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, err
		}
		index := table.GetIndexText(dataTblText, table.ColId, strconv.Itoa(int(node.Id)))
		if index == 0 { // entity_id does not exist, add record
			table.AppendText(node, dataTblTextPointer, plaintext)
		} else { // entity_id exists, update tags
			table.UpdateText(node, dataTblTextPointer, index)
		}
	}
	table.DeleteTextColId(dataTblTextPointer)
	logrus.Debug(dataTblText)

	return dataTblText, dataTblCard, nil
}
