package main

import (
	"math/rand"
	"time"
)

var issuedDecks []Deck

func getNewDeck() Deck {
	var deck = createDeck()
	issuedDecks = append(issuedDecks, deck)
	return deck
}

func findDeckByID(id string) Deck {
	for i := 0; i < len(issuedDecks); i++ {
		if issuedDecks[i].ID == id {
			return issuedDecks[i]
		}
	}
	return Deck{}
}

func getDeckIndex(id string) int {
	index := -1
	for i := 0; i < len(issuedDecks); i++ {
		if issuedDecks[i].ID == id {
			return i
		}
	}
	return index

}

func getNewDeckIndex() int {
	var deck = createDeck()
	issuedDecks = append(issuedDecks, deck)
	return len(issuedDecks) - 1
}

func drawFromNewDeck(count int) drawnCards {
	var index = getNewDeckIndex()

	var dCards []card
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		var cardInd = rand.Intn(issuedDecks[index].Size)

		dCards = append(dCards, issuedDecks[index].Cards[cardInd])
		issuedDecks[index].Cards = append(issuedDecks[index].Cards[:cardInd], issuedDecks[index].Cards[cardInd+1:]...)
		issuedDecks[index].Size--

	}
	return drawnCards{DeckID: issuedDecks[index].ID, DeckSize: issuedDecks[index].Size, Cards: dCards}
}
