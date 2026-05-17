package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsSeller  bool      `json:"is_seller"` // Satıcı mı alıcı mı ayrımı
	CreatedAt time.Time `json:"created_at"`
}

var CreateUserTableQuery = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_seller INTEGER DEFAULT 0, -- SQLite'ta bool olmadığı için 0 (false) veya 1 (true) tutulur
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`
