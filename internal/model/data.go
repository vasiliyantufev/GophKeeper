package model

import "time"

type Data struct {
	ID       int       `json:"data_id"`
	UserId   int       `json:"user_id"`
	Data     string    `json:"data"`
	Type     string    `json:"type"`
	CreateAt time.Time `json:"created_at"`
}

type DataRequest struct {
	ID     int    `json:"data_id"`
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
	Type   string `json:"type"`
}
