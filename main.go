package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	InitRouter()
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, World!")
}

func InitRouter() {
	router := httprouter.New()
	router.GET("/", Hello)

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

	fmt.Println("CORS works!")

	return handler
}
