package model

import "time"

type Cart struct {
	Name          string
	PaymentSystem string
	Number        string
	Holder        string
	EndData       time.Time
	CVC           int
}
