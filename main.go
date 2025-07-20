package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devldm/grammar-check-go/config"
	"github.com/devldm/grammar-check-go/db"
	"github.com/devldm/grammar-check-go/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	portString := config.Config("PORT")

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router := router.SetupRouter(dbConn)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}
