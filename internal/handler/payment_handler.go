package handler

import (
	"encoding/json"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) Pay(w http.ResponseWriter, r *http.Request) {
	var createPreferenceDTO contract.CreatePreferenceDTO

	if err := json.NewDecoder(r.Body).Decode(&createPreferenceDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.service.Pay(createPreferenceDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error to enconde response", http.StatusInternalServerError)
		return
	}
}
