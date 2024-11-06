package entity

import "github.com/google/uuid"

type Wallet struct {
	Id     uuid.UUID `gorm: "primaryKey;"`
	Amount float64
}
