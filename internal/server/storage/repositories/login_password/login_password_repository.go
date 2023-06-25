package loginPassword

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

type LoginPassword struct {
	db *database.DB
}

func New(db *database.DB) *LoginPassword {
	return &LoginPassword{
		db: db,
	}
}

func (lp *LoginPassword) CreateLoginPassword(loginPasswordRequest *model.CreateLoginPasswordRequest) (*model.LoginPassword, error) {
	loginPassword := &model.LoginPassword{}
	if err := lp.db.Pool.QueryRow(
		"INSERT INTO login_password (user_id, data, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING login_password_id, data",
		loginPasswordRequest.UserID,
		loginPasswordRequest.Data,
		time.Now(),
		time.Now(),
	).Scan(&loginPassword.ID, &loginPassword.Data); err != nil {
		return nil, err
	}
	return loginPassword, nil
}

func (lp *LoginPassword) KeyExists(loginPasswordRequest *model.CreateLoginPasswordRequest) (bool, error) {
	var exists bool
	row := lp.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM metadata "+
		"inner join login_password on metadata.entity_id = login_password.login_password_id "+
		"inner join users on login_password.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4)",
		string(variables.Name), loginPasswordRequest.Name, loginPasswordRequest.UserID, string(variables.LoginPassword))
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (lp *LoginPassword) GetNodeLoginPassword(loginPasswordRequest *model.GetNodeLoginPasswordRequest) (*model.LoginPassword, error) {
	loginPassword := &model.LoginPassword{}
	return loginPassword, nil
}

func (lp *LoginPassword) GetListLoginPassword(userId int64) ([]model.LoginPassword, error) {
	ListLoginPassword := []model.LoginPassword{}
	return ListLoginPassword, nil
}
