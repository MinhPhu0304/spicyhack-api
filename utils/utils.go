package utils

import (
	"math/rand"
	"time"
)

func GetRandomIndex(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
