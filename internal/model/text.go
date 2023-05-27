package model

import "time"

type Text struct {
	ID         int64
	UserID     int64
	MetadataID int64
	Text       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type CreateTextRequest struct {
	UserID int64
	Text   string
}

type GetListTextRequest struct {
	UserID int64
}

type GetNodeTextRequest struct {
	Name string
}
