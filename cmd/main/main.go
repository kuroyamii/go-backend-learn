package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuroyamii/go-backend-learn/pkg/controller"
)

func main() {
	router := mux.NewRouter()
	controller.InitializeControllers(router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err.Error())
	}
}
