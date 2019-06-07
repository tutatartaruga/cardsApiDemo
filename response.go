package main

// ResponseDeck struct
type ResponseDeck struct {
	Status int  `json:"status"`
	Data   Deck `json:"data"`
}

// ResponseDeckNotFound struct
type ResponseDeckNotFound struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//  ResponseDrawnCards struct
type ResponseDrawnCards struct {
	Status int        `json:"status"`
	Data   drawnCards `json:"data"`
}
