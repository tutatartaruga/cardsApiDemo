package main

import (
	"math/rand"
	"time"
)

func getRandomUint32() uint32 {
	rand.Seed(time.Now().Unix())
	return rand.Uint32()
}
