package table

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

func InitTblText(row int) [][]string {
	matrix := make([][]string, row)
	for i := range matrix {
		matrix[i] = make([]string, ColTblText)
	}
	return matrix
}

func InitTblCart(row int) [][]string {
	matrix := make([][]string, row)
	for i := range matrix {
		matrix[i] = make([]string, ColTblCart)
	}
	return matrix
}
