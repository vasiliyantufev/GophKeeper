package user

import (
	"database/sql"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
)

type Constructor interface {
	Registration(user *model.UserRequest) (*model.User, error)
	Authentication(userRequest *model.UserRequest) (*model.User, error)
	UserExists(username string) (bool, error)
}

type User struct {
	db *database.DB
}

func New(db *database.DB) *User {
	return &User{
		db: db,
	}
}
func (u *User) Registration(user *model.UserRequest) (*model.User, error) {
	registeredUser := &model.User{}
	if err := u.db.Pool.QueryRow(
		"INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING user_id, username",
		user.Username,
		user.Password,
		time.Now(),
	).Scan(&registeredUser.ID, &registeredUser.Username); err != nil {
		return nil, err
	}
	return registeredUser, nil
}

func (u *User) Authentication(userRequest *model.UserRequest) (*model.User, error) {
	authenticatedUser := &model.User{}
	err := u.db.Pool.QueryRow("SELECT user_id, username FROM users WHERE username=$1 and password=$2",
		userRequest.Username, userRequest.Password).Scan(
		&authenticatedUser.ID,
		&authenticatedUser.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrWrongUsernameOrPassword
		} else {
			return nil, err
		}
	}
	return authenticatedUser, nil
}

func (u *User) UserExists(username string) (bool, error) {
	var exists bool
	row := u.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM users where username = $1)", username)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (u *User) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	rows, err := u.db.Pool.Query("SELECT user_id, username FROM users where deleted_at IS NULL")
	//rows, err := u.db.Pool.Query("SELECT * FROM users where deleted_at IS NULL")
	if err != nil {
		if err == sql.ErrNoRows {
			return users, errors.ErrRecordNotFound
		} else {
			return users, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		user := model.User{}

		err = rows.Scan(&user.ID, &user.Username)
		//err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) Block(userID int64) (int64, error) {
	var id int64
	if err := u.db.Pool.QueryRow("UPDATE users SET deleted_at = $1 "+
		"where user.user_id = $2 RETURNING user_id",
		time.Now(),
		userID,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *User) Unblock(userID int64) (int64, error) {
	var id int64
	if err := u.db.Pool.QueryRow("UPDATE users SET deleted_at = null "+
		"where user.user_id = $1 RETURNING user_id",
		userID,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
