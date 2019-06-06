package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/getNewDeck", getNewDeck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
