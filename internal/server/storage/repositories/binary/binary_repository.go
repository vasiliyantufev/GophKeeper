package binary

import (
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

func (t *Binary) UploadBinary(binaryRequest *model.UploadBinaryRequest) (*model.Binary, error) {
	binary := &model.Binary{}
	return binary, nil
}

func (t *Binary) GetNodeBinary(binaryRequest *model.GetNodeBinaryRequest) (*model.Binary, error) {
	binary := &model.Binary{}
	return binary, nil
}

func (t *Binary) GetListBinary(userId int64) ([]model.Binary, error) {
	listBinary := []model.Binary{}
	return listBinary, nil
}

func (t *Binary) GetIdBinary(value string, userID int64) (int64, error) {
	var BinaryID int64
	return BinaryID, nil
}

func (t *Binary) KeyExists(BinaryRequest *model.UploadBinaryRequest) (bool, error) {
	var exists bool
	return exists, nil
}

func (t *Binary) DeleteBinary(binaryID int64) error {
	return nil
}

func (t *Binary) DownloadBinary(binaryID int64, data []byte) error {
	return nil
}
