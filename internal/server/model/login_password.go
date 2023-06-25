package model

import (
	"time"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

type LoginPassword struct {
	ID        int64
	UserID    int64
	Key       string
	Value     string
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateLoginPasswordRequest struct {
	UserID      int64
	Name        string
	Description string
	Type        string
	Data        []byte
	AccessToken string
}

type GetNodeLoginPasswordRequest struct {
	UserID      int64
	Key         string
	Value       string
	AccessToken string
}

func GetLoginPasswordData(data *LoginPassword) *grpc.LoginPassword {
	created, _ := service.ConvertTimeToTimestamp(data.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(data.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(data.DeletedAt)
	return &grpc.LoginPassword{
		UserId:    data.UserID,
		Data:      data.Data,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}
