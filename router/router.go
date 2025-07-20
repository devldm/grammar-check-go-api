package router

import (
	"database/sql"

	"github.com/devldm/grammar-check-go/handlers"
	"github.com/devldm/grammar-check-go/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRouter(dbq *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	handler := handlers.NewAPIConfig(database.New(dbq))
	// router.Use(middleware.ConfigMiddleware(config))

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlers.HandlerCheckHealth)

	v1Router.Get("/grammars", handler.HandlerGetAllGrammars)
	v1Router.Post("/grammars", handler.CreateGrammarChallenge)
	v1Router.Get("/grammars/{grammarId}", handler.GetGrammarByID)

	v1Router.Post("/user", handler.CreateUser)
	v1Router.Get("/user/{clerkUserId}", handler.GetUserByClerkID)

	v1Router.Post("/solutions", handler.CreateSolution)
	v1Router.Get("/solutions", handler.GetAllSolutions)
	v1Router.Post("/solutions/delete", handler.DeleteSolution)

	v1Router.Get("/solutions/user/{clerkUserId}", handler.GetSolutionsByUser)
	v1Router.Get("/solutions/{clerkUserId}/{grammarId}", handler.GetHasUserSolvedGrammar)
	v1Router.Get("/solutions/{grammarId}", handler.GetSolutionsByGrammarIDWithUserData)

	router.Mount("/v1", v1Router)

	return router
}
