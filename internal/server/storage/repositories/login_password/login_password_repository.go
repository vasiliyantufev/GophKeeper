package login_password

import (
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
)

type LoginPassword struct {
	db *database.DB
}

func New(db *database.DB) *LoginPassword {
	return &LoginPassword{
		db: db,
	}
}

func (lp *LoginPassword) CreateCard(cardRequest *model.CreateLoginPasswordRequest) (*model.LoginPassword, error) {
	loginPassword := &model.LoginPassword{}
	return loginPassword, nil
}

func (lp *LoginPassword) KeyExists(cardRequest *model.CreateLoginPasswordRequest) (bool, error) {
	var exists bool
	return exists, nil
}

func (lp *LoginPassword) GetNodeCard(cardRequest *model.GetNodeLoginPasswordRequest) (*model.LoginPassword, error) {
	loginPassword := &model.LoginPassword{}
	return loginPassword, nil
}

func (lp *LoginPassword) GetListCard(userId int64) ([]model.LoginPassword, error) {
	ListLoginPassword := []model.LoginPassword{}
	return ListLoginPassword, nil
}
