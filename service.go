package main

import (
	"encoding/json"
	"net/http"
)

func getNewDeck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createDeck())
}
