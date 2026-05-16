package routes

import (
	"net/http"

	"github.com/HakanVardarr/go-ecommerce-api/handlers"
	"github.com/HakanVardarr/go-ecommerce-api/repository"
)

func NewRouter(store *repository.ProductStore) http.Handler {
	mux := http.NewServeMux()

	productHandler := handlers.NewProductHandler(store)
	healthHandler := handlers.NewHealthHandler()

	mux.HandleFunc("GET /health", healthHandler.Healthcheck)
	mux.HandleFunc("GET /products", productHandler.GetProducts)

	return mux
}
