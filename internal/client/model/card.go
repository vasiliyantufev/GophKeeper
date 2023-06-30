package model

import "time"

type Card struct {
	Name          string    `json:"Name"`
	Description   string    `json:"Description"`
	PaymentSystem string    `json:"PaymentSystem"`
	Number        string    `json:"Number"`
	Holder        string    `json:"Holder"`
	CVC           int       `json:"CVC"`
	EndData       time.Time `json:"EndData"`
}
