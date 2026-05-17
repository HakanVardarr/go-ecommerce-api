package repository

import (
	"database/sql"
	"fmt"

	"github.com/HakanVardarr/go-ecommerce-api/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{}
}

func (s *UserStore) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password, is_seller) VALUES (?, ?, ?);`

	isSellerInt := 0
	if user.IsSeller {
		isSellerInt = 1
	}

	_, err := s.db.Exec(query, user.Email, user.Password, isSellerInt)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password, is_seller, created_at FROM users WHERE email = ?;`

	var user models.User
	var isSellerInt int

	err := s.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &isSellerInt, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	user.IsSeller = isSellerInt == 1
	return &user, nil
}
