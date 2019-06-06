package main

// Deck struct
type Deck struct {
	ID    string   `json:"id"`
	Size  int      `json:"size"`
	Cards [52]card `json:"cards"`
}

func createDeck() Deck {
	var id = getRandomString()
	var cards [52]card
	index := 0
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(suits); j++ {
			cards[index] = card{Code: getCode(values[i], suits[j]), Value: values[i], Suit: suits[j]}
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
