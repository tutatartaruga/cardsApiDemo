package main

type card struct {
	Code     string `json:"code"`
	Value    string `json:"value"`
	Suit     string `json:"suit"`
	ImageURL string `json:"image_url"`
}
