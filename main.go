package main

import (
	"Microservice_REST_UDEMY/items-api_bookstore/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)
func main(){
	router.HandleFunc("/items", controllers.Create).Methods(http.MethodPost)
}