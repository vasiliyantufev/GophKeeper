package model

import "time"

type Token struct {
	ID       int       `json:"token_id"`
	UserId   int       `json:"user_id"`
	Token    string    `json:"token"`
	CreateAt time.Time `json:"created_at"`
}
