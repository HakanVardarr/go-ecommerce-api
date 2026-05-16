package handlers

import (
	"encoding/json"
	"net/http"

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
	products := h.Store.GetAllProduct()
	json.NewEncoder(w).Encode(products)
}
