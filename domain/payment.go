package domain

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID          uint64 `gorm:"primary_ky"`
	AccountId   uint64
	Type        string
	Amount      float64
	Channel     string
	Description string
	Destination string
}

type PaymentResponse struct {
	ID          uint64  `json:"id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Channel     string  `json:"channel"`
	Description string  `json:"description"`
	Destination string  `json:"destination"`
}
