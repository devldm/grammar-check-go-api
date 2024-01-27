package config

import "github.com/devldm/grammar-check-go/internal/database"

type APIConfig struct {
	DB *database.Queries
}

// Initialize the config
func NewAPIConfig(db *database.Queries) *APIConfig {
	return &APIConfig{
		DB: db,
	}
}
