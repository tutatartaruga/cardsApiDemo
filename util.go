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
