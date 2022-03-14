package customprng

import (
	"testing"
)

func BenchmarkSlice(b *testing.B) {
	rng := New(20)
	for n := 0; n < b.N; n++ {
		rng.slice(6, 20)
	}
}

func BenchmarkString(b *testing.B) {
	rng := New(20)
	for n := 0; n < b.N; n++ {
		rng.String(6, 20)
	}
}

func BenchmarkStringSuffix(b *testing.B) {
	rng := New(24)
	for n := 0; n < b.N; n++ {
		rng.StringSuffix(6, 20, ".ru.")
	}
}
