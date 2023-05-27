package model

import "time"

type Metadata struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type CreateMetadataRequest struct {
	Name        string
	Description string
	SessionKey  string
}

type GetMetadataRequest struct {
	MetadataId int64
	SessionKey string
}
