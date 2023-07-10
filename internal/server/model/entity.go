package model

import (
	"time"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

type Entity struct {
	ID        int64
	UserID    int64
	Data      []byte
	Metadata  MetadataEntity
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateEntityRequest struct {
	UserID      int64
	Data        []byte
	Metadata    MetadataEntity
	AccessToken string
}

type CreateEntityResponse struct {
	Entity Entity
}

type MetadataEntity struct {
	Name        string
	Description string
	Type        string
}

func GetEntity(data *Entity) *grpc.Entity {
	//var metadata MetadataEntity
	created, _ := service.ConvertTimeToTimestamp(data.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(data.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(data.DeletedAt)
	//err := json.Unmarshal([]byte(data.Metadata), &metadata)
	//if err != nil {
	//	return &grpc.Entity{}, err
	//}
	return &grpc.Entity{
		UserId: data.UserID,
		Data:   data.Data,
		//Metadata: data.Metadata,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}
