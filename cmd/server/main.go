package main

import (
	"log"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/config"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/handler"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
	})

	log.Println("Server running in port: 3000")
	http.ListenAndServe(":3000", router)
}
