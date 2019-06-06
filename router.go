package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/getNewDeck", getNewDeck).Methods("GET")
	r.HandleFunc("/api/findDeckById/{id}", findDeckByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
