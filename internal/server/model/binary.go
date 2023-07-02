package model

import (
	"time"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

type Binary struct {
	ID        int64
	UserID    int64
	Key       string
	Value     string
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UploadBinaryRequest struct {
	UserID      int64
	Name        string
	Description string
	Type        string
	Data        []byte
	AccessToken string
}

type UploadBinaryResponse struct {
	Binary Binary
}

// ----------------------------------------
type GetNodeBinaryRequest struct {
	UserID      int64
	Key         string
	Value       string
	AccessToken string
}

type GetNodeBinaryResponse struct {
	Key   string
	Value string
	Data  Binary
}

// ----------------------------------------
type GetListBinaryRequest struct {
	UserID      int64
	AccessToken string
}

type GetListBinaryResponse struct {
	Binary []Binary
}

func GetBinary(binary *Binary) *grpc.Binary {
	created, _ := service.ConvertTimeToTimestamp(binary.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(binary.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(binary.DeletedAt)
	return &grpc.Binary{
		UserId:    binary.UserID,
		Data:      binary.Data,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}

func GetListBinary(binary []Binary) []*grpc.Binary {
	items := make([]*grpc.Binary, len(binary))
	for i := range binary {
		created, _ := service.ConvertTimeToTimestamp(binary[i].CreatedAt)
		updated, _ := service.ConvertTimeToTimestamp(binary[i].UpdatedAt)
		items[i] = &grpc.Binary{Id: binary[i].ID, Key: binary[i].Key, Data: binary[i].Data, Value: binary[i].Value, CreatedAt: created, UpdatedAt: updated}
	}
	return items
}
