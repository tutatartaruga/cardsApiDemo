package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		json.NewEncoder(w).Encode(&ResponseError{Status: http.StatusNotFound, Message: deckNotFound})
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
	json.NewEncoder(w).Encode(&ResponseError{Status: http.StatusNotFound, Message: deckNotFound})
}

func newShuffleDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var index = getNewDeckIndex()
	json.NewEncoder(w).Encode(&ResponseDeck{Status: http.StatusOK, Data: shuffleDeck(&issuedDecks[index])})
}

func newDeckDrawHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var count, err = strconv.Atoi(params["count"])
	if err == nil && count > 0 {
		var drawnCards = drawFromDeck(count, -1)
		if len(drawnCards.Cards) > 0 {
			json.NewEncoder(w).Encode(&ResponseDrawnCards{Status: http.StatusOK, Data: drawnCards})
			return
		}
		json.NewEncoder(w).Encode(&ResponseError{Status: http.StatusNotFound, Message: notEnoughCards})
		return
	}
	json.NewEncoder(w).Encode(&ResponseError{Status: http.StatusNotFound, Message: countZero})
}

func oldDeckDrawHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var index = getDeckIndex(params["id"])
	if index < 0 {
		json.NewEncoder(w).Encode(&ResponseError{Status: http.StatusNotFound, Message: deckNotFound})
		return
	}
	var count, err = strconv.Atoi(params["count"])
	if err == nil && count > 0 {
		var drawnCards = drawFromDeck(count, index)
		if len(drawnCards.Cards) > 0 {
			json.NewEncoder(w).Encode(&ResponseDrawnCards{Status: http.StatusOK, Data: drawnCards})
			return
		}
		json.NewEncoder(w).Encode(&ResponseError{Status: http.StatusNotFound, Message: notEnoughCards})
		return
	}
	json.NewEncoder(w).Encode(&ResponseError{Status: http.StatusNotFound, Message: countZero})
}
