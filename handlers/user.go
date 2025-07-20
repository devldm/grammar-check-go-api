package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devldm/grammar-check-go/helpers"
	"github.com/devldm/grammar-check-go/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func UserCreatedWebhook(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

func (c *APIConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ClerkID       string `json:"id"`
		ClerkUsername string `json:"username"`
		ClerkEmail    string `json:"email"`
		ClerkImage    string `json:"image"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	user, err := c.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		ClerkID:       params.ClerkID,
		ClerkUsername: params.ClerkUsername,
		ClerkEmail:    params.ClerkEmail,
		ClerkImage:    params.ClerkImage,
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusCreated, user)
}

func (c *APIConfig) GetUserByClerkID(w http.ResponseWriter, r *http.Request) {
	clerkUserIDParam := chi.URLParam(r, "clerkUserId")

	user, err := c.DB.GetUserByClerkId(r.Context(), clerkUserIDParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding user by clerk_id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, user)
}
