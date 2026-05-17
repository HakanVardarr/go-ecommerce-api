package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HakanVardarr/go-ecommerce-api/middleware"
	"github.com/HakanVardarr/go-ecommerce-api/models"
	"github.com/HakanVardarr/go-ecommerce-api/repository"
)

type ProductHandler struct {
	Store *repository.ProductStore
}

func NewProductHandler(store *repository.ProductStore) *ProductHandler {
	return &ProductHandler{Store: store}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.Store.GetAllProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized: User ID missing from context"})
		return
	}

	var prod models.Product
	err := json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	prod.SellerID = userID

	err = h.Store.CreateProduct(&prod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prod)
}
