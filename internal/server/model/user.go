package model

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt timestamp.Timestamp
	UpdatedAt timestamp.Timestamp
	DeletedAt timestamp.Timestamp
}

type GetAllUsers struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserRequest struct {
	Username string
	Password string
}

func GetUserData(data *User) *grpc.User {
	return &grpc.User{
		UserId:    data.ID,
		Username:  data.Username,
		CreatedAt: &data.CreatedAt,
		UpdatedAt: &data.UpdatedAt,
		DeletedAt: &data.DeletedAt,
	}
}
