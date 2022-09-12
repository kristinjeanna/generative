package random

import (
	"crypto/rand"
	"encoding/binary"
	mrnd "math/rand"

	"github.com/kbjorklu/xoshiro"
)

// NewXoshiro512StarStar returns a new xoshiro512** instance. Can
// be passed to Initialize to initialize the PRNG used by this package.
func NewXoshiro512StarStar() mrnd.Source64 {
	// read 8 cryptographically-secure bytes & convert to int64
	b := make([]byte, 8)
	_, err := rand.Read(b) //#nosec
	if err != nil {
		panic(err)
	}
	seed1 := int64(binary.BigEndian.Uint64(b))

	// next, create a SplitMix64 PRNG with seed1; SplitMix64 is the
	// recommended way to seed a xoshiro512** PRNG, per the authors
	sm64 := xoshiro.NewSplitMix64(seed1)
	seed2 := int64(sm64.Uint64())

	// finally create the xoshiro512** PRNG
	return xoshiro.NewXoshiro512StarStar(seed2)
}
