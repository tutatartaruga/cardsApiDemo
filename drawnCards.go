package main

type drawnCards struct {
	DeckID   string `json:"deck_id"`
	DeckSize int    `json:"deck_size"`
	Cards    []card `json:"cards"`
}
