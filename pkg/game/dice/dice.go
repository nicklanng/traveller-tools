package dice

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func D6() int {
	return rand.Intn(6) + 1
}

func D66() int {
	return D6()*10 + D6()
}
