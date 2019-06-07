package main

import "fmt"

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

func drawFromDeck(count, deckIndex int) drawnCards {
	var index int
	if deckIndex == -1 {
		index = getNewDeckIndex()
	} else {
		index = deckIndex
	}
	var dCards []card
	if count <= issuedDecks[index].Size {
		dCards = append(dCards, issuedDecks[index].Cards[:count]...)
		issuedDecks[index].Cards = append(issuedDecks[index].Cards[:0], issuedDecks[index].Cards[0+count:]...)
		issuedDecks[index].Size = issuedDecks[index].Size - count
	}
	fmt.Println(dCards)
	return drawnCards{DeckID: issuedDecks[index].ID, DeckSize: issuedDecks[index].Size, Cards: dCards}
}
