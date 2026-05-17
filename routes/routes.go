package routes

import (
	"net/http"

	"github.com/HakanVardarr/go-ecommerce-api/handlers"
	"github.com/HakanVardarr/go-ecommerce-api/middleware"
	"github.com/HakanVardarr/go-ecommerce-api/repository"
)

func NewRouter(productStore *repository.ProductStore, userStore *repository.UserStore) http.Handler {
	mux := http.NewServeMux()

	productHandler := handlers.NewProductHandler(productStore)
	healthHandler := handlers.NewHealthHandler()

	mux.HandleFunc("GET /health", healthHandler.Healthcheck)
	mux.HandleFunc("GET /products", productHandler.GetProducts)

	sellerOnly := middleware.AllowOnly("seller")
	mux.Handle("POST /products", middleware.AuthMiddleware(sellerOnly(http.HandlerFunc(productHandler.CreateProduct))))

	return mux
}
