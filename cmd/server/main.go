package main

import (
	"log"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/config"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/handler"
	userMiddlwere "github.com/azevedoguigo/florindas-acessorios-api/internal/middleware"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(userRepo, adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(userRepo, clientRepo)
	clientHadler := handler.NewClientHandler(clientService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productImageRepo := repository.NewProductImageRepository(db)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, productImageRepo)
	productHandler := handler.NewProductHandler(productService)

	authHanlder := handler.NewAuthHandler(userService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/admins", func(r chi.Router) {
		r.Use(userMiddlwere.AdminMiddleware)

		r.Post("/", adminHandler.CreateAdmin)
	})

	router.Route("/clients", func(r chi.Router) {
		r.Post("/", clientHadler.CreateClient)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/", authHanlder.Login)
	})

	router.Route("/categories", func(r chi.Router) {
		r.Use(userMiddlwere.AdminMiddleware)

		r.Post("/", categoryHandler.CreateCategory)
	})

	router.Route("/products-admin", func(r chi.Router) {
		r.Use(userMiddlwere.AdminMiddleware)

		r.Post("/", productHandler.CreateProduct)
	})

	router.Route("/products", func(r chi.Router) {
		r.Use(userMiddlwere.AuthMiddleware)

		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProductByID)
	})

	log.Println("Server running in port: 3000")
	http.ListenAndServe(":3000", router)
}
