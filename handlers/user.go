package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devldm/grammar-check-go/config"
	"github.com/devldm/grammar-check-go/helpers"
	"github.com/devldm/grammar-check-go/internal/database"
	"github.com/google/uuid"
)

func UserCreatedWebhook(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	type parameters struct {
		CLERK_ID string `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		ClerkID:   params.CLERK_ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusCreated, user)
}

func GetUserByClerkId(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	type parameters struct {
		ID string `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	user, err := apiConfig.DB.GetUserByClerkId(r.Context(), params.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding user by clerk_id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, user)

}
