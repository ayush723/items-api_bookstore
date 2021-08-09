package app

import (
	"net/http"

	"github.com/ayush723/items-api_bookstore/controllers"
)

func mapUrls() {
	

	// router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
