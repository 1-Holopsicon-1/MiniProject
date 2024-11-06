package utils

import (
	"MiniProject/internal/app/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/rand"
)

func Generate(db *gorm.DB) {
	for i := 0; i < 5; i++ {
		newObject := entity.Wallet{Id: uuid.New(), Amount: float64(rand.Intn(1000))}
		db.Create(&newObject)
	}
}
