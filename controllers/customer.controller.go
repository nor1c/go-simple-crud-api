package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	customerEntity "github.com/nor1c/go-simple-crud-api/entities"
)

type Controller interface {
	GetAll()
}

var validate *validator.Validate

func Register() {
	validate = validator.New()

	customer := &customerEntity.CustomerRegister{
		FullName: "A. Fauzi",
		Username: "nor1c",
		Password: "12345678",
	}

	err := validate.Struct(customer)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting all customers..")
}
