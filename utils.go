package stoppropaganda

import "math/rand"

var randomDomainRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func getRandomDomain() string {
	randomLength := rand.Intn(20-6) + 6 // from 6 to 20 characters length + ".ru"
	b := make([]rune, randomLength)
	for i := range b {
		b[i] = randomDomainRunes[rand.Intn(len(randomDomainRunes))]
	}
	return string(b) + ".ru"
}
