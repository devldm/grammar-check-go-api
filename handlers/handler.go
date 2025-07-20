package handlers

import "github.com/devldm/grammar-check-go/internal/database"

type APIConfig struct {
	DB *database.Queries
}

func NewAPIConfig(db *database.Queries) *APIConfig {
	return &APIConfig{
		DB: db,
	}
}
