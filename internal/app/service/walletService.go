package service

import (
	"MiniProject/internal/app/dto"
	"MiniProject/internal/app/entity"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type WalletService struct{}

func (walletService *WalletService) GetByUUID(db *gorm.DB, uuid string) dto.WalletDto {
	var (
		walletEntity entity.Wallet
		walletDto    dto.WalletDto
	)

	db.First(&walletEntity, "id = ?", uuid)
	walletDto.Id, walletDto.Amount = walletEntity.Id, walletEntity.Amount

	return walletDto
}

func (walletService *WalletService) Transfer(db *gorm.DB, r *http.Request) string {
	var (
		walletDto dto.WalletDto
	)
	jsonDecode := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&jsonDecode)
	if err != nil {
		log.Fatalln("Fail to decode", err)
	}
	walletDto.Id, err = uuid.Parse(jsonDecode["walletId"].(string))
	if err != nil {
		log.Println(err, "Error to uuid")
		return strconv.Itoa(http.StatusBadRequest) + " Wrong uuid"
	}
	walletDto.Amount, err = strconv.ParseFloat(jsonDecode["amount"].(string), 64)
	if err != nil {
		log.Println(err, "Fail to convert str to float")
		return strconv.Itoa(http.StatusBadRequest) + " Wrong number to amount"
	}
	if strings.ToLower(jsonDecode["operationType"].(string)) == "deposit" {
		return walletService.deposit(db, walletDto)
	} else if strings.ToLower(jsonDecode["operationType"].(string)) == "withdraw" {
		return walletService.withdraw(db, walletDto)
	} else {
		return strconv.Itoa(http.StatusBadRequest) + " No such command"
	}

}

func (walletService *WalletService) deposit(db *gorm.DB, walletDto dto.WalletDto) string {
	var walletEntity entity.Wallet
	walletEntity.Id = walletDto.Id
	if db.Where("id = ?", walletEntity.Id).Select("amount").First(&walletEntity).Error != nil {
		log.Println("No wallets by this id")
		return strconv.Itoa(http.StatusBadRequest) + " No wallets by this id"
	}
	walletEntity.Amount += walletDto.Amount
	db.Save(&walletEntity)
	return strconv.Itoa(http.StatusOK) + " Success deposit"
}

func (walletService *WalletService) withdraw(db *gorm.DB, walletDto dto.WalletDto) string {
	var walletEntity entity.Wallet
	walletEntity.Id = walletDto.Id
	if db.Where("id = ?", walletEntity.Id).Select("amount").First(&walletEntity).Error != nil {
		log.Println("No wallets by this id")
		return strconv.Itoa(http.StatusBadRequest) + " No wallets by this id"
	}
	if walletDto.Amount < 0 {
		walletDto.Amount *= -1
	}
	if walletEntity.Amount-walletDto.Amount <= 0 {
		return strconv.Itoa(http.StatusBadRequest) + " Not enough money"
	}
	walletEntity.Amount -= walletDto.Amount
	db.Save(&walletEntity)
	return strconv.Itoa(http.StatusOK) + " Success withdraw"
}
