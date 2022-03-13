package stoppropaganda

import "github.com/valyala/fastrand"

// Returns random string, from 6 to 20 characters length
func getRandomString(rng *fastrand.RNG) string {
	randomLength := rng.Uint32n(20-6) + 6 // from 6 to 20 characters length
	runes := uint32(len(randomDomainRunes))
	b := make([]rune, randomLength)
	for i := range b {
		idx := int(rng.Uint32n(runes))
		b[i] = randomDomainRunes[idx]
	}
	return string(b)
}
