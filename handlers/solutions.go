package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devldm/grammar-check-go/helpers"
	"github.com/devldm/grammar-check-go/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *APIConfig) GetAllSolutions(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing limit to integer: %v", err))
	}

	solutions, err := c.DB.GetSolutions(r.Context(), int32(limitInt))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding solutions: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, solutions)
}

func (c *APIConfig) GetSolutionsByGrammarIDWithUserData(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	grammarIDParam := chi.URLParam(r, "grammarId")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing limit to integer: %v", err))
	}

	uuidGrammarID, err := uuid.Parse(grammarIDParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing grammar id: %v", err))
	}

	solutions, err := c.DB.GetSolutionsByGrammarIdWithUserData(r.Context(), database.GetSolutionsByGrammarIdWithUserDataParams{
		GrammarID: uuidGrammarID,
		Limit:     int32(limitInt),
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding solutions by grammar id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, solutions)
}

func (c *APIConfig) GetSolutionsByUser(w http.ResponseWriter, r *http.Request) {
	clerkUserIDParam := chi.URLParam(r, "clerkUserId")

	user, err := c.DB.GetUserByClerkId(r.Context(), clerkUserIDParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching user by clerkId: %v", err))
	}

	solutions, err := c.DB.GetSolutionsByUserId(r.Context(), user.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding solutions by user: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, solutions)
}

func (c *APIConfig) GetHasUserSolvedGrammar(w http.ResponseWriter, r *http.Request) {
	clerkUserIDParam := chi.URLParam(r, "clerkUserId")
	grammarID := chi.URLParam(r, "grammarId")

	uuidGrammarID, err := uuid.Parse(grammarID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing grammar id: %v", err))
	}

	user, err := c.DB.GetUserByClerkId(r.Context(), clerkUserIDParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching user by clerkId: %v", err))
	}

	solutions, err := c.DB.GetHasUserSolved(r.Context(), database.GetHasUserSolvedParams{
		UserID:    user.ID,
		GrammarID: uuidGrammarID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			helpers.RespondWithJSON(w, http.StatusOK, nil)
			return
		} else {
			helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding solutions: %v", err))
		}
	}

	helpers.RespondWithJSON(w, http.StatusOK, solutions)
}

func (c *APIConfig) CreateSolution(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		GrammarID uuid.UUID `json:"grammar_id"`
		UserID    string    `json:"user_id"`
		Solution  string    `json:"solution"`
		Grammar   string    `json:"grammar"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	grammar, err := c.DB.GetGrammarById(r.Context(), params.GrammarID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Grammar could not be found.")
	}

	user, err := c.DB.GetUserByClerkId(r.Context(), params.UserID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "User could not be found.")
	}

	solution, err := c.DB.CreateSolution(r.Context(), database.CreateSolutionParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Solution:  params.Solution,
		UserID:    user.ID,
		GrammarID: params.GrammarID,
		Grammar:   grammar.Grammar,
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating grammar: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusCreated, solution)
}

func (c *APIConfig) DeleteSolution(w http.ResponseWriter, r *http.Request) {
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

	uuidSolutionID, err := uuid.Parse(params.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing solution id: %v", err))
	}

	err = c.DB.DeleteSolutionBySolutionId(r.Context(), uuidSolutionID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting solution with solution id: %v", err))
	}
}
