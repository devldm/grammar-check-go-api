package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devldm/grammar-check-go/config"
	"github.com/devldm/grammar-check-go/helpers"
	"github.com/devldm/grammar-check-go/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func UserCreatedWebhook(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	type parameters struct {
		CLERK_ID       string `json:"id"`
		CLERK_USERNAME string `json:"username"`
		CLERK_EMAIL    string `json:"email"`
		CLERK_IMAGE    string `json:"image"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		ClerkID:       params.CLERK_ID,
		ClerkUsername: params.CLERK_USERNAME,
		ClerkEmail:    params.CLERK_EMAIL,
		ClerkImage:    params.CLERK_IMAGE,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusCreated, user)
}

func GetUserByClerkId(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	clerkUserIdParam := chi.URLParam(r, "clerkUserId")

	user, err := apiConfig.DB.GetUserByClerkId(r.Context(), clerkUserIdParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding user by clerk_id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, user)

}
