package model

import "time"

type User struct {
	ID        int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
