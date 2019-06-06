package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var issuedDecks []Deck

func getNewDeck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newDeck = createDeck()
	issuedDecks = append(issuedDecks, newDeck)
	json.NewEncoder(w).Encode(newDeck)
}

func findDeckByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, deck := range issuedDecks {
		if deck.ID == params["id"] {
			json.NewEncoder(w).Encode(deck)
		}
	}
}
