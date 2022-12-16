package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/internals/controllers"
	"github.com/morelmiles/go-events/internals/helpers"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes() {

	port := os.Getenv("PORT")

	helpers.InitLogger()

	router := mux.NewRouter().StrictSlash(true)
	corsHandler := cors.Default().Handler(router)
	router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.WrapHandler)

	// Home
	router.HandleFunc("/", controllers.Home).Methods("GET")

	// Users
	router.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/v1/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", controllers.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/api/v1/user/{id}", controllers.UpdateUserById).Methods("PATCH")
	router.HandleFunc("/api/v1/user/events/{id}", controllers.GetAllEventsByUser).Methods("GET")

	// Events
	router.HandleFunc("/api/v1/events", controllers.GetEvents).Methods("GET")
	router.HandleFunc("/api/v1/events/{id}", controllers.GetEventById).Methods("GET")
	router.HandleFunc("/api/v1/create", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/api/v1/events/{id}", controllers.DeleteEventById).Methods("DELETE")
	router.HandleFunc("/api/v1/events/{id}", controllers.UpdateEventById).Methods("PUT")

	http.ListenAndServe(":"+port, corsHandler)
}
