package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/deck/new").Subrouter()
	subrouter.HandleFunc("", getNewDeck).Methods("GET")
	subrouter.HandleFunc("/shuffle", getNewShuffledDeck).Methods("GET")
	router.HandleFunc("/api/deck/{id}/get", findDeckByID).Methods("GET")
	router.HandleFunc("/api/deck/{id}/shuffle", shuffleDeckByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
