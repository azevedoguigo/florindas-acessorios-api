package handler

import (
	"encoding/json"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
)

type CartProductHandler struct {
	service service.CartProductService
}

func NewCartProductHandler(service service.CartProductService) *CartProductHandler {
	return &CartProductHandler{service: service}
}

func (h *CartProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newCartProduct contract.NewCartProductDTO

	if err := json.NewDecoder(r.Body).Decode(&newCartProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&newCartProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
