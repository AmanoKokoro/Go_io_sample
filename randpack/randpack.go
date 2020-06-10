package randpack

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

//Random aaa
type Random struct {
	Param int
}

// Roulette returns a randomly
// generated number of type Int
func (r Random) Roulette() int {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return int(rand.Int63()) % r.Param
}
