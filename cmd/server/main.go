package main

import (
	"log"
	"net/http"

	"github.com/azevedoguigo/florindas-acessorios-api/config"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/handler"
	userMiddlwere "github.com/azevedoguigo/florindas-acessorios-api/internal/middleware"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/service"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	pkg.LoadEnv()
	db := config.InitDB()
	s3Client := config.InitAWS()
	mercadoPagoClient := config.InitMercadoPago()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	cartRepo := repository.NewCartRepository(db)

	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(userRepo, adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(userRepo, clientRepo, cartRepo)
	clientHadler := handler.NewClientHandler(clientService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productImageRepo := repository.NewProductImageRepository(db)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, productImageRepo, s3Client)
	productHandler := handler.NewProductHandler(productService)

	authHanlder := handler.NewAuthHandler(userService, adminService, clientService)

	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	cartProductRepo := repository.NewCartProductRepository(db)
	cartProductService := service.NewCartProductService(cartProductRepo, productRepo)
	cartProductHandler := handler.NewCartProductHandler(cartProductService)

	paymentService := service.NewPaymentService(mercadoPagoClient)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/admins", func(r chi.Router) {
		r.Use(userMiddlwere.AdminMiddleware)

		r.Post("/", adminHandler.CreateAdmin)
	})

	router.Route("/clients", func(r chi.Router) {
		r.Post("/", clientHadler.CreateClient)

		r.With(userMiddlwere.AuthMiddleware).Get("/{id}", clientHadler.GetClientByID)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/", authHanlder.Login)
		r.With(userMiddlwere.AuthMiddleware).Get("/me", authHanlder.Me)
	})

	router.Route("/categories", func(r chi.Router) {
		r.Use(userMiddlwere.AuthMiddleware)

		r.With(userMiddlwere.AdminMiddleware).Post("/", categoryHandler.CreateCategory)
		r.Get("/", categoryHandler.GetCategories)
		r.Get("/{id}", categoryHandler.GetCategoryByID)
	})

	router.Route("/products-admin", func(r chi.Router) {
		r.Use(userMiddlwere.AdminMiddleware)

		r.Post("/", productHandler.CreateProduct)
	})

	router.Route("/products", func(r chi.Router) {
		r.Use(userMiddlwere.AuthMiddleware)

		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProductByID)
		r.Put("/{id}", productHandler.UpdateProduct)
	})

	router.Route("/cart", func(r chi.Router) {
		r.Use(userMiddlwere.AuthMiddleware)

		r.Get("/", cartHandler.GetCartByUserID)
	})

	router.Route("/cart-product", func(r chi.Router) {
		r.Use(userMiddlwere.AuthMiddleware)

		r.Post("/", cartProductHandler.CreateCartProduct)
		r.Put("/", cartProductHandler.UpdateCartProductQuantity)
	})

	router.Route("/payment", func(r chi.Router) {
		r.Use(userMiddlwere.AuthMiddleware)

		r.Post("/", paymentHandler.Pay)
	})

	log.Println("Server running in port: 3000")
	http.ListenAndServe(":3000", router)
}
