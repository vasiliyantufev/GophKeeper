package model

import (
	"time"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

type Card struct {
	ID        int64
	UserID    int64
	CardData  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateCardRequest struct {
	UserID      int64
	Name        string
	Type        string
	CardData    []byte
	AccessToken string
}

type CreateCardResponse struct {
	Card Card
}

func GetCardData(data *Card) *grpc.Card {
	created, _ := service.ConvertTimeToTimestamp(data.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(data.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(data.DeletedAt)
	return &grpc.Card{
		UserId:    data.UserID,
		Data:      data.CardData,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}
