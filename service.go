package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

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

func shuffleDeckByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var index int
	for i := 0; i < len(issuedDecks); i++ {
		if issuedDecks[i].ID == params["id"] {
			index = i
			break
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(issuedDecks[index].Cards), func(i, j int) {
		issuedDecks[index].Cards[i], issuedDecks[index].Cards[j] = issuedDecks[index].Cards[j], issuedDecks[index].Cards[i]
	})
	json.NewEncoder(w).Encode(issuedDecks[index])
}
