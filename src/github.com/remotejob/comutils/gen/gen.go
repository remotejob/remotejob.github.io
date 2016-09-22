package gen

import (
	"math/rand"
)

func Random(min, max int) int {
	//	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
