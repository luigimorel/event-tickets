package routes

import (
	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/internals/controllers"
	"github.com/morelmiles/go-events/internals/middleware"
	"github.com/morelmiles/go-events/internals/server"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes() error {
	middleware.InitLogger()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggerMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/", middleware.SetMiddlewareJSON(controllers.Home)).Methods("GET")

	auth := api.PathPrefix("").Subrouter()
	auth.HandleFunc("/login", middleware.SetMiddlewareJSON(controllers.Login)).Methods("POST")
	auth.HandleFunc("/register", middleware.SetMiddlewareJSON(controllers.CreateUser)).Methods("POST")

	users := api.PathPrefix("/user").Subrouter()
	users.HandleFunc("s", middleware.SetMiddlewareAuthentication(controllers.GetUsers)).Methods("GET")
	users.HandleFunc("/{id}", middleware.SetMiddlewareJSON(controllers.GetUserById)).Methods("GET")
	users.HandleFunc("/{id}", middleware.SetMiddlewareAuthentication(controllers.DeleteUserById)).Methods("DELETE")
	users.HandleFunc("/{id}", middleware.SetMiddlewareAuthentication(controllers.UpdateUserById)).Methods("PATCH")
	users.HandleFunc("/events/{id}", middleware.SetMiddlewareJSON(controllers.GetAllEventsByUser)).Methods("GET")

	events := api.PathPrefix("/events").Subrouter()
	events.HandleFunc("", middleware.SetMiddlewareJSON(controllers.GetEvents)).Methods("GET")
	events.HandleFunc("/homepage", middleware.SetMiddlewareJSON(controllers.GetHomePageEvents)).Methods("GET")
	events.HandleFunc("/{id}", middleware.SetMiddlewareJSON(controllers.GetEventById)).Methods("GET")
	events.HandleFunc("", middleware.SetMiddlewareAuthentication(controllers.CreateEvent)).Methods("POST")
	events.HandleFunc("/{id}", middleware.SetMiddlewareAuthentication(controllers.DeleteEventById)).Methods("DELETE")
	events.HandleFunc("/{id}", middleware.SetMiddlewareAuthentication(controllers.UpdateEventById)).Methods("PUT")

	corsHandler := middleware.SetupCORS()(router)
	srv := server.NewServer(corsHandler, nil)
	return srv.Start()
}
