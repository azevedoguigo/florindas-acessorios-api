package handler

import (
	"encoding/json"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
	"github.com/go-chi/chi/v5"
)

type CartProductHandler struct {
	service service.CartProductService
}

func NewCartProductHandler(service service.CartProductService) *CartProductHandler {
	return &CartProductHandler{service: service}
}

func (h *CartProductHandler) CreateCartProduct(w http.ResponseWriter, r *http.Request) {
	var newCartProductDTO contract.NewCartProductDTO

	if err := json.NewDecoder(r.Body).Decode(&newCartProductDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCartProduct(&newCartProductDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CartProductHandler) UpdateCartProductQuantity(w http.ResponseWriter, r *http.Request) {
	var updateCartProductQuantityDTO contract.UpdateCartProductQuantityDTO

	if err := json.NewDecoder(r.Body).Decode(&updateCartProductQuantityDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCartProductQuantity(&updateCartProductQuantityDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CartProductHandler) DeleteCartProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.DeleteCartProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
