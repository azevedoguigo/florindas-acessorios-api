package handler

import (
	"encoding/json"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory contract.NewCategoryDTO

	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCategory(&newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
