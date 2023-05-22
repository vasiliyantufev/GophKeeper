package model

import "time"

type Metadata struct {
	ID       int       `json:"metadata_id"`
	DataId   int       `json:"data_id"`
	Metadata string    `json:"metadata"`
	CreateAt time.Time `json:"created_at"`
}

type MetaRequest struct {
	ID int `json:"metadata_id"`
}
