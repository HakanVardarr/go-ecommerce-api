package repository

import (
	"database/sql"
	"fmt"

	"github.com/HakanVardarr/go-ecommerce-api/models"
)

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

func (s *ProductStore) CreateProduct(product *models.Product) error {
	query := `INSERT INTO products (seller_id, name, price, stock) VALUES (?, ?, ?, ?);`

	result, err := s.db.Exec(query, product.SellerID, product.Name, product.Price, product.Stock)
	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	id, err := result.LastInsertId()
	if err == nil {
		product.ID = int(id)
	}

	return nil
}

func (s *ProductStore) GetAllProducts() ([]models.Product, error) {
	query := `SELECT id, seller_id, name, price, stock, created_at FROM products;`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.SellerID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %w", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return products, nil
}
