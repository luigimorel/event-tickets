package controllers

import (
	"net/http"

	"github.com/morelmiles/go-events/internals/helpers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	helpers.JSON(w, http.StatusOK, "Welcome to the Ticket API")
}
