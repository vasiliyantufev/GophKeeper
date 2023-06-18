package table

import (
	"strconv"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

const ColId = 0
const ColName = 1
const ColText = 2
const ColDescription = 3
const ColTblText = 5
const ColTblCart = 9

func SearchByColumn(slice [][]string, targetColumn int, targetValue string) bool {
	for i := 1; i < len(slice) && len(slice) > 1; i++ {
		if slice[i][targetColumn] == targetValue {
			return true
		}
	}
	return false
}

func GetIndexText(slice [][]string, targetColumn int, targetValue string) (index int) {
	for index = 1; index < len(slice) && len(slice) > 1; index++ {
		if slice[index][targetColumn] == targetValue {
			return index
		}
	}
	return 0
}

func AppendText(node *grpc.Text, dataTblText *[][]string, plaintext string) {
	layout := "01/02/2006 15:04:05"
	created, _ := service.ConvertTimestampToTime(node.CreatedAt)
	updated, _ := service.ConvertTimestampToTime(node.UpdatedAt)
	if node.Key == string(variables.Name) {
		row := []string{strconv.Itoa(int(node.Id)), node.Value, plaintext, "", created.Format(layout), updated.Format(layout)}
		*dataTblText = append(*dataTblText, row)
	} else if node.Key == string(variables.Description) {
		row := []string{strconv.Itoa(int(node.Id)), "", plaintext, node.Value, created.Format(layout), updated.Format(layout)}
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

func DeleteTextColId(dataTblText *[][]string) {
	for index := range *dataTblText {
		(*dataTblText)[index] = (*dataTblText)[index][1:]
	}
}
