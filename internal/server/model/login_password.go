package model

import "time"

type LoginPassword struct {
	ID        int64
	UserID    int64
	Key       string
	Value     string
	CardData  []byte
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
