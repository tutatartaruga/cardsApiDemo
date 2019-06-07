package main

// ResponseDeck struct
type ResponseDeck struct {
	Status int  `json:"status"`
	Data   Deck `json:"data"`
}

// ResponseDrawnCards struct
type ResponseDrawnCards struct {
	Status int        `json:"status"`
	Data   drawnCards `json:"data"`
}

// ResponseError struct
type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
