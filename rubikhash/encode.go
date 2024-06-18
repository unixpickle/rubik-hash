package rubikhash

import "github.com/unixpickle/gocube"

// StateToBits creates a "mostly uniform" bit string from
// the state of a cube.
//
// Even if the state is sampled uniformly at random, some
// bit strings will have slightly higher probability by a
// delta of about 2^(-34), since the number of unique
// states is not an exact multiple of 2^32.
func StateToBits(state *gocube.CubieCube) uint32 {
	return uint32(uint64(state.Corners.EncodeIndex()) +
		state.Edges.EncodeIndex(false)*(3*3*3*3*3*3*3*1*2*3*4*5*6*7*8))
}
