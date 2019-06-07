package main

import (
	"math/rand"
	"time"
)

// Deck struct
type Deck struct {
	ID    string `json:"id"`
	Size  int    `json:"size"`
	Cards []card `json:"cards"`
}

func createDeck() Deck {
	var id = getRandomString()
	var cards []card
	index := 0
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(suits); j++ {
			cards = append(cards, card{Code: getCode(values[i], suits[j]), Value: values[i], Suit: suits[j]})
			index++
		}
	}
	return Deck{id, len(cards), cards}
}

func getCode(value, suit string) string {
	if value == "10" {
		return value[0:2] + suit[0:1]
	}
	return value[0:1] + suit[0:1]
}

func shuffleDeck(deck *Deck) Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return *deck
}
