package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devldm/grammar-check-go/helpers"
	"github.com/devldm/grammar-check-go/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *APIConfig) HandlerGetAllGrammars(w http.ResponseWriter, r *http.Request) {
	grammars, err := c.DB.GetGrammars(r.Context(), 50)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching grammars: %v", err))
		return
	}
	// models.DatabaseFeedsToFeeds(feeds)

	helpers.RespondWithJSON(w, http.StatusOK, grammars)
}

func (c *APIConfig) GetGrammarByID(w http.ResponseWriter, r *http.Request) {
	grammarIDParam := chi.URLParam(r, "grammarId")
	grammarID, err := uuid.Parse(grammarIDParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing grammar id: %v", err))
	}

	grammar, err := c.DB.GetGrammarById(r.Context(), grammarID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding grammar by grammar id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, grammar)
}

func (c *APIConfig) CreateGrammarChallenge(w http.ResponseWriter, r *http.Request) {
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

	grammar, err := c.DB.CreateGrammar(r.Context(), database.CreateGrammarParams{
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
