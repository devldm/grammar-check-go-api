package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devldm/grammar-check-go/config"
	"github.com/devldm/grammar-check-go/helpers"
	"github.com/devldm/grammar-check-go/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetAllSolutions(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)
	limit := r.URL.Query().Get("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing limit to integer: %v", err))

	}

	solutions, err := apiConfig.DB.GetSolutions(r.Context(), int32(limitInt))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding grammar by grammar id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, solutions)
}

func GetSolutionsByGrammarIdWithUserData(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)
	limit := r.URL.Query().Get("limit")
	grammarIdParam := chi.URLParam(r, "grammarId")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing limit to integer: %v", err))

	}

	uuidGrammarId, err := uuid.Parse(grammarIdParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing grammar id: %v", err))
	}

	solutions, err := apiConfig.DB.GetSolutionsByGrammarIdWithUserData(r.Context(), database.GetSolutionsByGrammarIdWithUserDataParams{
		GrammarID: uuidGrammarId,
		Limit:     int32(limitInt),
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding grammar by grammar id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, solutions)
}

func GetSolutionsByUser(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)
	clerkUserIdParam := chi.URLParam(r, "clerkUserId")

	user, err := apiConfig.DB.GetUserByClerkId(r.Context(), clerkUserIdParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching user by clerkId: %v", err))
	}

	solutions, err := apiConfig.DB.GetSolutionsByUserId(r.Context(), user.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error finding grammar by grammar id: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusOK, solutions)
}

func GetHasUserSolvedGrammar(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)
	clerkUserIdParam := chi.URLParam(r, "clerkUserId")
	grammar_id := chi.URLParam(r, "grammarId")

	uuidGrammarId, err := uuid.Parse(grammar_id)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing grammar id: %v", err))
	}

	user, err := apiConfig.DB.GetUserByClerkId(r.Context(), clerkUserIdParam)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching user by clerkId: %v", err))
	}

	solutions, err := apiConfig.DB.GetHasUserSolved(r.Context(), database.GetHasUserSolvedParams{
		UserID:    user.ID,
		GrammarID: uuidGrammarId,
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

func CreateSolution(w http.ResponseWriter, r *http.Request) {
	apiConfig := r.Context().Value("api_config").(*config.APIConfig)

	type parameters struct {
		GrammarId uuid.UUID `json:"grammar_id"`
		UserId    string    `json:"user_id"`
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

	grammar, err := apiConfig.DB.GetGrammarById(r.Context(), params.GrammarId)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Grammar could not be found.")
	}

	user, err := apiConfig.DB.GetUserByClerkId(r.Context(), params.UserId)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "User could not be found.")
	}

	solution, err := apiConfig.DB.CreateSolution(r.Context(), database.CreateSolutionParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Solution:  params.Solution,
		UserID:    user.ID,
		GrammarID: params.GrammarId,
		Grammar:   grammar.Grammar,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating grammar: %v", err))
	}

	helpers.RespondWithJSON(w, http.StatusCreated, solution)
}
