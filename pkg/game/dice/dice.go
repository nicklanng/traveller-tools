package dice

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func D6(number int) (sum int) {
	for i := 0; i < 10; i++ {
		sum += rand.Intn(6) + 1
	}
	return
}

func D66() int {
	return D6(1)*10 + D6(1)
}
