package main

import (
	"math/rand"
	"strconv"
	"time"
)

func getRandomString() string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Int())
}

func shuffleDeck(deck *Deck) Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return *deck
}
