package loginPassword

import (
	"database/sql"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
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
	err := lp.db.Pool.QueryRow("SELECT login_password.data FROM metadata "+
		"inner join login_password on metadata.entity_id = login_password.login_password_id "+
		"inner join users on login_password.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4",
		string(variables.Name), loginPasswordRequest.Value, loginPasswordRequest.UserID, string(variables.LoginPassword)).
		Scan(&loginPassword.Data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return loginPassword, nil
}

func (lp *LoginPassword) GetListLoginPassword(userId int64) ([]model.LoginPassword, error) {
	ListLoginPassword := []model.LoginPassword{}

	rows, err := lp.db.Pool.Query("SELECT metadata.entity_id, metadata.key, login_password.data, metadata.value, login_password.created_at, "+
		"login_password.updated_at FROM metadata "+
		"inner join login_password on metadata.entity_id = login_password.login_password_id "+
		"inner join users on login_password.user_id  = users.user_id "+
		"where users.user_id = $1 and metadata.type = $2",
		userId, string(variables.Card))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		loginPassword := model.LoginPassword{}
		err = rows.Scan(&loginPassword.ID, &loginPassword.Key, &loginPassword.Data, &loginPassword.Value, &loginPassword.CreatedAt, &loginPassword.UpdatedAt)
		if err != nil {
			return nil, err
		}
		ListLoginPassword = append(ListLoginPassword, loginPassword)
	}

	return ListLoginPassword, nil
}
