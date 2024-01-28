package handlers

import (
	"database/sql"
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

func HandlerGetAllGrammars(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	grammars, err := apiConfig.DB.GetGrammars(r.Context(), 10)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching grammars: %v", err))
		return
	}
	//models.DatabaseFeedsToFeeds(feeds)

	helpers.RespondWithJSON(w, http.StatusOK, grammars)
}

func GetGrammarById(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	grammarIdParam := chi.URLParam(r, "grammarId")
	grammarId, err := uuid.Parse(grammarIdParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing grammar id: %v", err))
	}

	grammar, err := apiConfig.DB.GetGrammarById(r.Context(), grammarId)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding grammar by grammar id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, grammar)

}

func CreateGrammarChallenge(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	type parameters struct {
		Grammar     string `json:"grammar"`
		Difficulty  string `json:"difficulty"`
		Description string `json:"description"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	description := sql.NullString{}
	if params.Description != "" {
		description.String = params.Description
		description.Valid = true
	}

	difficulty := sql.NullString{}
	if params.Difficulty != "" {
		difficulty.String = params.Difficulty
		difficulty.Valid = true
	}

	grammar, err := apiConfig.DB.CreateGrammar(r.Context(), database.CreateGrammarParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Grammar:     params.Grammar,
		Description: description,
		Difficulty:  difficulty,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating grammar: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusCreated, grammar)
}
