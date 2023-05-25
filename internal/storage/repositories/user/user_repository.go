package user

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/database"
	"github.com/vasiliyantufev/gophkeeper/internal/model"
)

type Constructor interface {
	Registration(data *model.RegistrationRequest) error
}

type User struct {
	db *database.DB
}

func New(db *database.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Registration(user *model.RegistrationRequest) (int, error) {
	var id int
	if err := u.db.Pool.QueryRow(
		"INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING user_id",
		user.Username,
		user.Password,
		time.Now(),
	).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (u *User) Login(user *model.LoginRequest) (bool, error) {
	var exists bool
	row := u.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM users where username = $1 AND password = $2)", user.Username, user.Password)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
	//authorizedUser := &model.User{}
	//if err := u.db.Pool.QueryRow("SELECT * FROM users where username = $1 AND password = $2", user.Username, user.Password).Scan(
	//	&authorizedUser.ID,
	//	&authorizedUser.Username,
	//	&authorizedUser.Password,
	//	&authorizedUser.CreatedAt,
	//	&authorizedUser.DeletedAt,
	//); err != nil {
	//	return nil, err
	//}
	//return authorizedUser, nil
}
