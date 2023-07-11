package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

	log.Fatal(http.ListenAndServe(":3000", router))
}
