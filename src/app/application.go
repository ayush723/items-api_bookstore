package app

import (
	"net/http"
	"time"

	"github.com/ayush723/items-api_bookstore/src/clients/elasticsearch"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8082",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	if err := (srv.ListenAndServe()); err != nil {
		panic(err)
	}
}
