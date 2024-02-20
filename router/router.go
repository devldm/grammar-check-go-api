package router

import (
	"github.com/devldm/grammar-check-go/config"
	"github.com/devldm/grammar-check-go/handlers"
	"github.com/devldm/grammar-check-go/internal/database"
	"github.com/devldm/grammar-check-go/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRouter(dbq *database.Queries) *chi.Mux {
	router := chi.NewRouter()
	config := config.NewAPIConfig(dbq)
	router.Use(middleware.ConfigMiddleware(config))

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

	v1Router.Get("/grammars", handlers.HandlerGetAllGrammars)
	v1Router.Post("/grammars", handlers.CreateGrammarChallenge)
	v1Router.Get("/grammars/{grammarId}", handlers.GetGrammarById)

	v1Router.Post("/user", handlers.CreateUser)
	v1Router.Get("/user/{clerkUserId}", handlers.GetUserByClerkId)

	v1Router.Post("/solutions", handlers.CreateSolution)
	v1Router.Get("/solutions", handlers.GetAllSolutions)
	v1Router.Post("/solutions/delete", handlers.DeleteSolution)

	v1Router.Get("/solutions/user/{clerkUserId}", handlers.GetSolutionsByUser)
	v1Router.Get("/solutions/{clerkUserId}/{grammarId}", handlers.GetHasUserSolvedGrammar)
	v1Router.Get("/solutions/{grammarId}", handlers.GetSolutionsByGrammarIdWithUserData)

	router.Mount("/v1", v1Router)

	return router
}
