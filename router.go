package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/deck/new").Subrouter()
	subrouter.HandleFunc("", newDeckHandler).Methods("GET")
	subrouter.HandleFunc("/shuffle", newShuffleDeckHandler).Methods("GET")
	router.HandleFunc("/api/deck/{id}/get", findDeckHandler).Methods("GET")
	router.HandleFunc("/api/deck/{id}/shuffle", shuffleDeckHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
