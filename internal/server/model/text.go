package model

import (
	"time"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

type Text struct {
	ID        int64
	UserID    int64
	Key       string
	Value     string
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateTextRequest struct {
	UserID      int64
	Name        string
	Description string
	Type        string
	Data        []byte
	AccessToken string
}

type CreateTextResponse struct {
	Text Text
}

// ----------------------------------------
type GetNodeTextRequest struct {
	UserID      int64
	Key         string
	Value       string
	AccessToken string
}

type GetNodeTextResponse struct {
	Key   string
	Value string
	Data  Text
}

// ----------------------------------------
type GetListTextRequest struct {
	UserID      int64
	AccessToken string
}

type GetListTextResponse struct {
	Text []Text
}

func GetText(text *Text) *grpc.Text {
	created, _ := service.ConvertTimeToTimestamp(text.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(text.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(text.DeletedAt)
	return &grpc.Text{
		UserId:    text.UserID,
		Text:      text.Data,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}

func GetListText(texts []Text) []*grpc.Text {
	items := make([]*grpc.Text, len(texts))
	for i := range texts {
		created, _ := service.ConvertTimeToTimestamp(texts[i].CreatedAt)
		updated, _ := service.ConvertTimeToTimestamp(texts[i].UpdatedAt)
		items[i] = &grpc.Text{Id: texts[i].ID, Key: texts[i].Key, Text: texts[i].Data, Value: texts[i].Value, CreatedAt: created, UpdatedAt: updated}
	}
	return items
}
