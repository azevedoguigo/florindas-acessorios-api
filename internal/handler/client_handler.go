package handler

import (
	"encoding/json"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
	"github.com/go-chi/chi/v5"
)

type ClientHandler struct {
	service service.ClientService
}

func NewClientHandler(service service.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var newClient contract.NewClientDTO

	if err := json.NewDecoder(r.Body).Decode(&newClient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateClient(&newClient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ClientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	client, err := h.service.GetClientByID(id)
	if err != nil && err.Error() == "client not found" {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, "Error to enconde resonse", http.StatusInternalServerError)
		return
	}
}
