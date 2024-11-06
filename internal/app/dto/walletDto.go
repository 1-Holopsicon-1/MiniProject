package dto

import "github.com/google/uuid"

type WalletDto struct {
	Id     uuid.UUID `json:"walletId"`
	Amount float64   `json:"amount"`
}
