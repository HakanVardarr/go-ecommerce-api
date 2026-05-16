package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/HakanVardarr/go-ecommerce-api/models"
)

type ProductStore struct {
	mu       sync.RWMutex
	products map[int]models.Product
	lastId   int
}

func NewProductStore(filePath string) (*ProductStore, error) {
	store := &ProductStore{
		products: make(map[int]models.Product),
		lastId:   0,
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return store, nil
	}

	var initalProducts []models.Product
	err = json.Unmarshal(fileData, &initalProducts)
	if err != nil {
		return store, fmt.Errorf("Failed to unmarshal: %w", err)
	}

	for _, p := range initalProducts {
		store.lastId++
		p.Id = store.lastId
		store.products[p.Id] = p
	}

	return store, nil
}

func (s *ProductStore) AddProduct(p models.Product) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastId++
	p.Id = s.lastId
	s.products[p.Id] = p

	return p.Id
}

func (s *ProductStore) GetAllProduct() []models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]models.Product, 0, len(s.products))
	for _, p := range s.products {
		list = append(list, p)
	}
	return list
}
