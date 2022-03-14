package customprng

import "github.com/valyala/fastrand"

/*
THIS CODE IS NOT THREAD-SAFE
*/

type RNG struct {
	rng    fastrand.RNG
	buff   []byte
	number uint32
}

func New(maxStringLength int) *RNG {
	return &RNG{
		rng:  fastrand.RNG{},
		buff: make([]byte, maxStringLength),
	}
}

var randomRunesList = []byte("abcdefghijklmnopqrstuvwxyz")
var randomRunesListLength = uint32(len(randomRunesList))

func (r *RNG) slice(min, max uint32) []byte {
	r.number = r.rng.Uint32n(max-min) + min
	r.buff = r.buff[:r.number]
	for r.number > 0 {
		r.number--
		idx := r.rng.Uint32n(randomRunesListLength)
		r.buff[r.number] = randomRunesList[idx]
	}
	return r.buff
}

func (r *RNG) String(min, max uint32) string {
	return string(r.slice(min, max))
}

func (r *RNG) StringSuffix(min, max uint32, suffix string) string {
	return string(append(r.slice(min, max), []byte(suffix)...))
}
