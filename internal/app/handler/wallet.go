package handler

import (
	"MiniProject/internal/app/service"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var walletService service.WalletService

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	err := json.NewEncoder(w).Encode(walletService.GetByUUID(h.DB, uuid))
	if err != nil {
		log.Fatalln("Error of Encoding")
	}
}

func (h *Handler) transfer(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(walletService.Transfer(h.DB, r))
	if err != nil {
		log.Fatalln(err)
	}
}
