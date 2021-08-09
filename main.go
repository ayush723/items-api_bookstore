package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)
func main(){
	router.HandleFunc("/items", controllers.Create).Methods(http.MethodPost)
}