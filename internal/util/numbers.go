package util

import (
	"math/rand"
)

// RandomBetween generates a random integer between a min int and a max int
func RandomBetween(min int, max int) int {
	return rand.Intn(max - min) + min
}
