package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
	"gorm.io/gorm"
)

type CartHandler struct {
	service service.CartService
}

func NewCartHandler(service service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) GetCartByUserID(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims, _ := tokenAuth.Decode(tokenString)

	userID, _ := claims.Get("user_id")

	cart, err := h.service.GetCartByUserID(fmt.Sprintf("%v", userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User does not exists!", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cart); err != nil {
		http.Error(w, "Error to encode respnse", http.StatusInternalServerError)
		return
	}
}
