package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	customerController "github.com/nor1c/go-simple-crud-api/controllers"
)

func CustomerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", customerController.GetAll)

	return r
}
