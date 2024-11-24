package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth = jwtauth.New("HS256", []byte("secretkey"), nil)

type AuthHandler struct {
	userService   service.UserService
	adminService  service.AdminService
	clientService service.ClientService
}

func NewAuthHandler(
	userService service.UserService,
	adminService service.AdminService,
	clientService service.ClientService,
) *AuthHandler {
	return &AuthHandler{
		userService:   userService,
		adminService:  adminService,
		clientService: clientService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO contract.LoginDTO
	json.NewDecoder(r.Body).Decode(&loginDTO)

	token, err := h.userService.Login(&loginDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims, _ := tokenAuth.Decode(tokenString)

	role, _ := claims.Get("role")
	userID, _ := claims.Get("user_id")

	if role == "admin" {
		client, err := h.adminService.GetAdminByUserID(fmt.Sprintf("%v", userID))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(client); err != nil {
			http.Error(w, "Error to enconde resonse", http.StatusInternalServerError)
			return
		}
	} else if role == "client" {
		client, err := h.clientService.GetClientByUserID(fmt.Sprintf("%v", userID))

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
}
