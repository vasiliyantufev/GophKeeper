package model

import (
	"time"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

type Card struct {
	ID        int64
	UserID    int64
	Key       string
	Value     string
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateCardRequest struct {
	UserID      int64
	Name        string
	Description string
	Type        string
	Data        []byte
	AccessToken string
}

type CreateCardResponse struct {
	Card Card
}

// ----------------------------------------
type GetNodeCardRequest struct {
	UserID      int64
	Key         string
	Value       string
	AccessToken string
}

type GetNodeCardResponse struct {
	Key   string
	Value string
	Card  Card
}

// ----------------------------------------

func GetCard(data *Card) *grpc.Card {
	created, _ := service.ConvertTimeToTimestamp(data.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(data.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(data.DeletedAt)
	return &grpc.Card{
		UserId:    data.UserID,
		Data:      data.Data,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}

func GetListCard(data []Card) []*grpc.Card {
	items := make([]*grpc.Card, len(data))
	for i := range data {
		created, _ := service.ConvertTimeToTimestamp(data[i].CreatedAt)
		updated, _ := service.ConvertTimeToTimestamp(data[i].UpdatedAt)
		items[i] = &grpc.Card{Id: data[i].ID, Key: data[i].Key, Data: data[i].Data, Value: data[i].Value, CreatedAt: created, UpdatedAt: updated}
	}
	return items
}
