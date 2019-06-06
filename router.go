package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/deck/new", getNewDeck).Methods("GET")
	r.HandleFunc("/api/deck/{id}/get", findDeckByID).Methods("GET")
	r.HandleFunc("/api/deck/{id}/shuffle", shuffleDeckByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
