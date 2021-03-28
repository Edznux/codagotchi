package rand

import (
	"math/rand"
)

/*
These random MUST BE fully deterministic so we create a new seed for EACH random.
I guess we could optimize it with some xorshift but I couldn't find anything that produced
convincing random numbers.
*/
func Int64(seed int64) int64 {
	r := rand.New(rand.NewSource(seed))
	return r.Int63()
}

//Float64 returns [0.0, 1.0)
func Float64(seed int64) float64 {
	r := rand.New(rand.NewSource(seed))
	return r.Float64()
}
