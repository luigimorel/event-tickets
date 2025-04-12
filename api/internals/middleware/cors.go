package middleware

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

func SetupCORS() func(http.Handler) http.Handler {
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Accept",
			"Origin",
		},
		AllowCredentials: true,
		Debug:            false,
	})

	return corsOpts.Handler
}
