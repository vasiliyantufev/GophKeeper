package table

import (
	"strconv"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

const ColId = 0
const ColName = 1
const ColDescription = 2
const ColText = 3
const ColTblText = 5
const ColTblCard = 9

func SearchByColumn(slice [][]string, targetColumn int, targetValue string) bool {
	for i := 1; i < len(slice) && len(slice) > 1; i++ {
		if slice[i][targetColumn] == targetValue {
			return true
		}
	}
	return false
}

func GetIndex(slice [][]string, targetColumn int, targetValue string) (index int) {
	for index = 1; index < len(slice) && len(slice) > 1; index++ {
		if slice[index][targetColumn] == targetValue {
			return index
		}
	}
	return 0
}

func RemoveRow(slice [][]string, indexRow int) [][]string {
	return append(slice[:indexRow], slice[indexRow+1:]...)
}

func UpdateRowLoginPassword(login, password string, slice [][]string, indexRow int) [][]string {
	layout := "01/02/2006 15:04:05"
	indexColLogin := 2
	indexColPassword := 3
	indexColUpdateAt := 5
	slice[indexRow][indexColLogin] = login
	slice[indexRow][indexColPassword] = password
	slice[indexRow][indexColUpdateAt] = time.Now().Format(layout)
	return slice
}

func UpdateRowText(text string, slice [][]string, indexRow int) [][]string {
	layout := "01/02/2006 15:04:05"
	indexColText := 2
	indexColUpdateAt := 4
	slice[indexRow][indexColText] = text
	slice[indexRow][indexColUpdateAt] = time.Now().Format(layout)
	return slice
}

func AppendText(node *grpc.Text, dataTblText *[][]string, plaintext string) {
	layout := "01/02/2006 15:04:05"
	created, _ := service.ConvertTimestampToTime(node.CreatedAt)
	updated, _ := service.ConvertTimestampToTime(node.UpdatedAt)
	if node.Key == string(variables.Name) {
		row := []string{strconv.Itoa(int(node.Id)), node.Value, "", plaintext, created.Format(layout), updated.Format(layout)}
		*dataTblText = append(*dataTblText, row)
	} else if node.Key == string(variables.Description) {
		row := []string{strconv.Itoa(int(node.Id)), "", node.Value, plaintext, created.Format(layout), updated.Format(layout)}
		*dataTblText = append(*dataTblText, row)
	}
}

func UpdateText(node *grpc.Text, dataTblText *[][]string, index int) {
	if node.Key == string(variables.Name) {
		(*dataTblText)[index][ColName] = node.Value
	} else if node.Key == string(variables.Description) {
		(*dataTblText)[index][ColDescription] = node.Value
	}
}

func AppendCard(node *grpc.Card, dataTblCard *[][]string, jsonCard model.Card) {
	layoutEndData := "01/02/2006"
	layout := "01/02/2006 15:04:05"
	created, _ := service.ConvertTimestampToTime(node.CreatedAt)
	updated, _ := service.ConvertTimestampToTime(node.UpdatedAt)
	if node.Key == string(variables.Name) {
		row := []string{strconv.Itoa(int(node.Id)), node.Value, "", jsonCard.PaymentSystem, jsonCard.Number,
			jsonCard.Holder, strconv.Itoa(jsonCard.CVC), jsonCard.EndData.Format(layoutEndData), created.Format(layout), updated.Format(layout)}
		*dataTblCard = append(*dataTblCard, row)
	} else if node.Key == string(variables.Description) {
		row := []string{strconv.Itoa(int(node.Id)), "", node.Value, jsonCard.PaymentSystem, jsonCard.Number,
			jsonCard.Holder, strconv.Itoa(jsonCard.CVC), jsonCard.EndData.Format(layoutEndData), created.Format(layout), updated.Format(layout)}
		*dataTblCard = append(*dataTblCard, row)
	}
}

func UpdateCard(node *grpc.Card, dataTblCard *[][]string, index int) {
	if node.Key == string(variables.Name) {
		(*dataTblCard)[index][ColName] = node.Value
	} else if node.Key == string(variables.Description) {
		(*dataTblCard)[index][ColDescription] = node.Value
	}
}

func AppendLoginPassword(node *grpc.LoginPassword, dataTblLoginPassword *[][]string, jsonLoginPassword model.LoginPassword) {
	layout := "01/02/2006 15:04:05"
	created, _ := service.ConvertTimestampToTime(node.CreatedAt)
	updated, _ := service.ConvertTimestampToTime(node.UpdatedAt)
	if node.Key == string(variables.Name) {
		row := []string{strconv.Itoa(int(node.Id)), node.Value, "", jsonLoginPassword.Login, jsonLoginPassword.Password,
			created.Format(layout), updated.Format(layout)}
		*dataTblLoginPassword = append(*dataTblLoginPassword, row)
	} else if node.Key == string(variables.Description) {
		row := []string{strconv.Itoa(int(node.Id)), "", node.Value, jsonLoginPassword.Login, jsonLoginPassword.Password,
			created.Format(layout), updated.Format(layout)}
		*dataTblLoginPassword = append(*dataTblLoginPassword, row)
	}
}

func UpdateLoginPassword(node *grpc.LoginPassword, dataTblLoginPassword *[][]string, index int) {
	if node.Key == string(variables.Name) {
		(*dataTblLoginPassword)[index][ColName] = node.Value
	} else if node.Key == string(variables.Description) {
		(*dataTblLoginPassword)[index][ColDescription] = node.Value
	}
}

func DeleteColId(dataTblText *[][]string) {
	for index := range *dataTblText {
		(*dataTblText)[index] = (*dataTblText)[index][1:]
	}
}
