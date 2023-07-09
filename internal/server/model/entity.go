package model

import (
	"time"
)

type Entity struct {
	ID        int64
	UserID    int64
	Data      []byte
	Metadata  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateEntityRequest struct {
	UserID      int64
	Data        []byte
	Metadata    string
	Type        string
	AccessToken string
}

type CreateEntityResponse struct {
	Entity Entity
}
