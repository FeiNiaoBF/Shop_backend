package utility

import (
	"math/rand"
	"time"
)

func RandInt(max int) int {
	seed := time.Now().UnixMilli()
	randomNum := rand.New(rand.NewSource(seed))
	return randomNum.Intn(max)
}
