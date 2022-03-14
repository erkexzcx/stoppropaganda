package stoppropaganda

import (
	"github.com/valyala/fastrand"
)

var randomRunesList = []byte("abcdefghijklmnopqrstuvwxyz")
var randomRunesListLength = uint32(len(randomRunesList))

// Returns random string, from 6 to 20 characters length
func getRandomString(rng *fastrand.RNG) string {
	randomLength := rng.Uint32n(20-6) + 6
	b := make([]byte, randomLength)
	for i := uint32(0); i < randomLength; i++ {
		idx := int(rng.Uint32n(randomRunesListLength))
		b[i] = randomRunesList[idx]
	}
	return string(b)
}
