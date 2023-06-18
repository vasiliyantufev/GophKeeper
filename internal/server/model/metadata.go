package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

type Metadata struct {
	ID        int64
	EntityId  int64
	Key       string
	Value     string
	Type      string
	CreatedAt timestamp.Timestamp
	UpdatedAt timestamp.Timestamp
	DeletedAt timestamp.Timestamp
}

type CreateMetadataRequest struct {
	EntityId    int64
	Key         string
	Value       string
	Type        string
	AccessToken string
}

func GetMetadata(data *Metadata) *grpc.Metadata {
	return &grpc.Metadata{
		EntityId:  data.EntityId,
		Key:       data.Key,
		Value:     data.Value,
		Type:      data.Type,
		CreatedAt: &data.CreatedAt,
		UpdatedAt: &data.UpdatedAt,
		DeletedAt: &data.DeletedAt,
	}
}
