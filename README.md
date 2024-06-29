# rubik-hash

This is an experiment in creating a cryptographic hash function using a digital representation of the Rubik's cube. I have run this through the [NIST Statistical Test Suite](https://csrc.nist.gov/projects/random-bit-generation/documentation-and-software) successfully, suggesting that it's at least acceptable for implementing a (somewhat inefficient) pseudo-random number generator. However, I have not attempted to do more advanced cryptographic tests on the hash algorithm.

# Idea

*Note:* I'll adopt the standard [Singmaster notation](https://en.wikipedia.org/wiki/Rubik%27s_Cube#Singmaster_notation) to denote moves on a Rubik's cube in this section.

Our input is a string of bits `b[0], ..., b[N-1]` where `b[i]` is a boolean. The output of our hash algorithm will be the state of a Rubik's cube, which can be represented as an integer in `[0, 43,252,003,274,489,856,000)`. The basic idea is to define two alternative random sequences of moves, and choose which move to make at each step based on the corresponding bit in the input.

In particular, we will assume we have two sequences `s0` and `s1`, each of length `N*K`, where `s0[i]` and `s1[i]` are both face turns on the Rubik's cube (e.g. F, D2, L'). For example, sequence 0 might be `R, L, D2, F', ...` and sequence 1 might be `D, R2, B2, L', ...`. To hash the input bits, we simply apply the move sequence `s` where `s[i] = (b[i%N] ? s0[i] : s1[i])`.

Note the modulus in the above formula, which is necessary because we reuse each bit `b[i]` of the input sequence `K` times. In particular, our sequences of moves are `K` times as long as our input bit sequence, so we use bit `b[i]` at positions `i`, `i+N`, ..., `i+N*(K-1)`. This is intended to cause a waterfall effect, where changing a single bit can completely alter ever aspect of the final cube state. Without reusing a bit multiple times at different parts of the sequence, a given bit would only change four corners and four edges of the final cube state.
