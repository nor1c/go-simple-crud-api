package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/unrolled/secure"

	route "github.com/nor1c/go-simple-crud-api/routes"
)

func main() {
	InitRouter()
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func InitRouter() {
	router := chi.NewRouter()
	router.Get("/", Hello)
	router.Mount("/customers", route.CustomerRoutes())

	handler := middlewares(router)

	// serve and automatically print log if error occured
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func middlewares(handler http.Handler) http.Handler {
	handler = cors.New(
		cors.Options{
			AllowedOrigins: []string{"http://localhost:5173"},
			AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			Debug:          false,
		},
	).Handler(handler)

	handler = secure.New(secure.Options{
		FrameDeny: true,
	}).Handler(handler)

	return handler
}
