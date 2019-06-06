package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	subrouter := router.PathPrefix("/api/deck").Subrouter()
	subrouter.HandleFunc("/new", getNewDeck).Methods("GET")
	subrouter.HandleFunc("/{id}/get", findDeckByID).Methods("GET")
	subrouter.HandleFunc("/{id}/shuffle", shuffleDeckByID).Methods("GET")
	subrouter.HandleFunc("/shuffle/new", getNewShuffledDeck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
