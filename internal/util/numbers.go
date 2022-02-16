package util

import (
  "math/rand"
)

func RandomBetween(min int, max int) int {
  return rand.Intn(max - min) + min
}
