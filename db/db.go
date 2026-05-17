package db

import (
	"database/sql"
	"fmt"

	"github.com/HakanVardarr/go-ecommerce-api/models"
	_ "modernc.org/sqlite"
)

func InitDB(dbPath string) (*sql.DB, error) {
	database, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = database.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	_, err = database.Exec(models.CreateUserTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	_, err = database.Exec(models.CreateProductTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create products table: %w", err)
	}

	return database, nil
}
