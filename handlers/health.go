package handlers

import (
	"net/http"

	"github.com/devldm/grammar-check-go/helpers"
)

func HandlerCheckHealth(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}
