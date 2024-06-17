package rubikhash

import (
	"fmt"
	"math/rand"

	"github.com/unixpickle/gocube"
)

// A HashFunc turns a sequence of bits into a cube state.
type HashFunc struct {
	Bits   int
	Moves0 []gocube.Move
	Moves1 []gocube.Move
}

// RandomHashFunc creates a new hash function from a random
// source given the number of bits to encode and the number
// of passes to do over the bits.
//
// The passes argument should be at least 1, and the more
// passes which are done, the more random the hash is.
func RandomHashFunc(source rand.Source, bits, passes int) *HashFunc {
	rng := rand.New(source)
	res := &HashFunc{
		Bits:   bits,
		Moves0: make([]gocube.Move, 0, bits*passes),
		Moves1: make([]gocube.Move, 0, bits*passes),
	}
	for i := 0; i < bits*passes; i++ {
		move0 := gocube.Move(rng.Intn(18))
		move1 := gocube.Move(rng.Intn(17))
		if move1 >= move0 {
			move1 += 1
		}
		res.Moves0 = append(res.Moves0, move0)
		res.Moves1 = append(res.Moves0, move1)
	}
	return res
}

// HashExact hashes data of the exact size that is expected
// by this hash function.
//
// The resulting set of moves are applied to the initial
// state.
//
// If the data is incorrectly sized, this will panic().
func (h *HashFunc) HashExact(state *gocube.CubieCube, data []byte) {
	if len(data)*8 != h.Bits {
		panic(fmt.Sprintf("incorrect number of bits: %d (expected %d)", len(data)*8, h.Bits))
	}
	for i, m0 := range h.Moves0 {
		m1 := h.Moves1[i]
		if data[i/8]&(1<<uint(i%8)) == 0 {
			state.Move(m0)
		} else {
			state.Move(m1)
		}
	}
}
