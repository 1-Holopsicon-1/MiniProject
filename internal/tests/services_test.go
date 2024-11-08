package tests

import (
	"MiniProject/internal/app/db"
	"MiniProject/internal/app/entity"
	"MiniProject/internal/app/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"testing"
)

func TestGetById(t *testing.T) {
	session := db.TestConnect()
	testEntity := entity.Wallet{Id: uuid.New(), Amount: float64(rand.IntN(100))}
	session.Migrator().DropTable(&entity.Wallet{})
	session.Migrator().CreateTable(&entity.Wallet{})
	session.Create(&testEntity)
	walletService := new(service.WalletService)
	walletById := walletService.GetByUUID(session, testEntity.Id.String())
	if walletById.Id != testEntity.Id || walletById.Amount != testEntity.Amount {
		t.Errorf("Entities not compare, got: %v, want: %v", walletById, testEntity)
	}
}

func TestDeposit(t *testing.T) {
	session := db.TestConnect()
	testEntity := entity.Wallet{Id: uuid.New(), Amount: float64(rand.IntN(100))}
	session.Migrator().DropTable(&entity.Wallet{})
	session.Migrator().CreateTable(&entity.Wallet{})
	session.Create(&testEntity)
	oldAmount := testEntity.Amount
	walletService := new(service.WalletService)
	operation := struct {
		WalletId      uuid.UUID `json:"walletId"`
		OperationType string    `json:"operationType"`
		Amount        float64   `json:"amount"`
	}{WalletId: testEntity.Id,
		OperationType: "deposit",
		Amount:        30,
	}
	req, err := json.Marshal(operation)
	if err != nil {
		log.Fatalln("error of json")
	}

	r, err := http.NewRequest("POST", "http://localhost", bytes.NewBuffer(req))
	if err != nil {
		log.Fatalln("Fail to create request", err)
	}
	fmt.Println(r.Body)

	walletService.Transfer(session, r)
	session.Find(&testEntity)
	if testEntity.Amount-oldAmount != operation.Amount {
		t.Errorf("Error of deposit, not done want: %f, got: %f", operation.Amount, testEntity.Amount-oldAmount)
	}
}

func TestWithDraw(t *testing.T) {
	session := db.TestConnect()
	testEntity := entity.Wallet{Id: uuid.New(), Amount: float64(rand.IntN(100))}
	session.Migrator().DropTable(&entity.Wallet{})
	session.Migrator().CreateTable(&entity.Wallet{})
	session.Create(&testEntity)
	oldAmount := testEntity.Amount
	walletService := new(service.WalletService)
	operation := struct {
		WalletId      uuid.UUID `json:"walletId"`
		OperationType string    `json:"operationType"`
		Amount        float64   `json:"amount"`
	}{WalletId: testEntity.Id,
		OperationType: "withdraw",
		Amount:        -30,
	}
	req, err := json.Marshal(operation)
	if err != nil {
		log.Fatalln("error of json")
	}

	r, err := http.NewRequest("POST", "http://localhost", bytes.NewBuffer(req))
	if err != nil {
		log.Fatalln("Fail to create request", err)
	}
	status := walletService.Transfer(session, r)
	session.Find(&testEntity)
	if operation.Amount < 0 {
		operation.Amount *= -1
	}
	if status != strconv.Itoa(http.StatusBadRequest) {
		if testEntity.Amount+operation.Amount != oldAmount {
			t.Errorf("Error of withdraw, not done want: %f, got: %f", oldAmount, testEntity.Amount+operation.Amount)
		}
	}

}
