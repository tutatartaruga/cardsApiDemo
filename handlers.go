package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func newDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ResponseDeck{Status: http.StatusOK, Data: getNewDeck()})
}

func findDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var deck = findDeckByID(params["id"])
	if len(deck.ID) > 0 {
		json.NewEncoder(w).Encode(&ResponseDeck{Status: http.StatusOK, Data: deck})
	} else {
		json.NewEncoder(w).Encode(&ResponseDeckNotFound{Status: http.StatusNotFound, Message: deckNotFound})
	}
}

func shuffleDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var index = getDeckIndex(params["id"])
	if index >= 0 {
		json.NewEncoder(w).Encode(&ResponseDeck{Status: http.StatusOK, Data: shuffleDeck(&issuedDecks[index])})
		return
	}
	json.NewEncoder(w).Encode(&ResponseDeckNotFound{Status: http.StatusNotFound, Message: deckNotFound})
}

func newShuffleDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var index = getNewDeckIndex()
	json.NewEncoder(w).Encode(&ResponseDeck{Status: http.StatusOK, Data: shuffleDeck(&issuedDecks[index])})
}
