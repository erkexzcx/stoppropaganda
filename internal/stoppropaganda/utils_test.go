package stoppropaganda

import (
	"testing"

	"github.com/valyala/fastrand"
)

func BenchmarkGetRandomString(b *testing.B) {
	rng := &fastrand.RNG{}
	for n := 0; n < b.N; n++ {
		getRandomString(rng)
	}
}
