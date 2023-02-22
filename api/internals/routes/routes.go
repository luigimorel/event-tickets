package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/internals/controllers"
	"github.com/morelmiles/go-events/internals/helpers"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes() {

	helpers.InitLogger()

	router := mux.NewRouter().StrictSlash(true)
	corsHandler := cors.Default().Handler(router)

	// Swggers
	router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.WrapHandler)

	// Home
	router.HandleFunc("/", helpers.SetMiddlewareJSON(controllers.Home)).Methods("GET")

	// Auth
	router.HandleFunc("/api/v1/login", helpers.SetMiddlewareJSON(controllers.Login)).Methods("POST")

	// Users
	router.HandleFunc("/api/v1/users", helpers.SetMiddlewareAuthentication(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", helpers.SetMiddlewareJSON(controllers.GetUserById)).Methods("GET")
	router.HandleFunc("/api/v1/register", helpers.SetMiddlewareJSON(controllers.CreateUser)).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", helpers.SetMiddlewareAuthentication(controllers.DeleteUserById)).Methods("DELETE")
	router.HandleFunc("/api/v1/user/{id}", helpers.SetMiddlewareAuthentication(controllers.UpdateUserById)).Methods("PATCH")
	router.HandleFunc("/api/v1/user/events/{id}", helpers.SetMiddlewareJSON(controllers.GetAllEventsByUser)).Methods("GET")

	// Events
	router.HandleFunc("/api/v1/events", helpers.SetMiddlewareJSON(controllers.GetEvents)).Methods("GET")
	router.HandleFunc("/api/v1/homepage_events", helpers.SetMiddlewareJSON(controllers.GetHomePageEvents)).Methods("GET")
	router.HandleFunc("/api/v1/events/{id}", helpers.SetMiddlewareJSON(controllers.GetEventById)).Methods("GET")
	router.HandleFunc("/api/v1/create", helpers.SetMiddlewareAuthentication(controllers.CreateEvent)).Methods("POST")
	router.HandleFunc("/api/v1/events/{id}", helpers.SetMiddlewareAuthentication(controllers.DeleteEventById)).Methods("DELETE")
	router.HandleFunc("/api/v1/events/{id}", helpers.SetMiddlewareAuthentication(controllers.UpdateEventById)).Methods("PUT")

	// Server port
	http.ListenAndServe(":8080", corsHandler)
}
