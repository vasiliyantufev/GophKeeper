package model

import "time"

type Card struct {
	Name          string    `json:"Name"`
	PaymentSystem string    `json:"PaymentSystem"`
	Number        string    `json:"Number"`
	Holder        string    `json:"Holder"`
	EndData       time.Time `json:"EndData"`
	CVC           int       `json:"CVC"`
}
