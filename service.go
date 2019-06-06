package main

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
	var newDeck = createDeck()
	issuedDecks = append(issuedDecks, newDeck)
	return len(issuedDecks) - 1
}
