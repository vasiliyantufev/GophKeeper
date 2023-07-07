package binary

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
)

type Binary struct {
	db *database.DB
}

func New(db *database.DB) *Binary {
	return &Binary{
		db: db,
	}
}

func (b *Binary) UploadBinary(binaryRequest *model.BinaryRequest) (*model.Binary, error) {
	binary := &model.Binary{}
	if err := b.db.Pool.QueryRow(
		"INSERT INTO binary_data (user_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING binary_id, name",
		binaryRequest.UserID,
		binaryRequest.Name,
		time.Now(),
		time.Now(),
	).Scan(&binary.ID, &binary.Name); err != nil {
		return nil, err
	}
	return binary, nil
}

func (b *Binary) GetNodeBinary(binaryRequest *model.GetNodeBinaryRequest) (*model.Binary, error) {
	binary := &model.Binary{}
	return binary, nil
}

func (b *Binary) GetListBinary(userId int64) ([]model.Binary, error) {
	listBinary := []model.Binary{}
	return listBinary, nil
}

func (b *Binary) GetIdBinary(value string, userID int64) (int64, error) {
	var BinaryID int64
	return BinaryID, nil
}

func (b *Binary) FileExists(binaryRequest *model.BinaryRequest) (bool, error) {
	var exists bool
	row := b.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM binary_data "+
		"where binary_data.user_id = $1 and binary_data.name = $2 and binary_data.deleted_at IS NULL)",
		binaryRequest.UserID,
		binaryRequest.Name)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (b *Binary) DeleteBinary(binaryRequest *model.BinaryRequest) (int64, error) {
	var id int64
	if err := b.db.Pool.QueryRow("UPDATE binary_data SET deleted_at = $1 "+
		"where binary_data.user_id = $2 and binary_data.name = $3 RETURNING binary_id",
		time.Now(),
		binaryRequest.UserID,
		binaryRequest.Name,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
