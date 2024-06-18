// Command hash_bitstring creates a string of bits to be
// fed to an RNG test suite.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/gocube"
	"github.com/unixpickle/rubik-hash/rubikhash"
)

func main() {
	var seed int64
	var rounds int
	var numBits int
	var outputFile string
	flag.Int64Var(&seed, "seed", 0, "seed for building hash")
	flag.IntVar(&rounds, "rounds", 2, "rounds for hash")
	flag.IntVar(&numBits, "bits", 1000000, "bits to generate")
	flag.StringVar(&outputFile, "output-file", "", "dump binary to file")
	flag.Parse()

	// Build a hash function deterministically based on a seed.
	// The amount of data we read from the source is small.
	source := rand.NewSource(seed)
	hash := rubikhash.RandomHashFunc(source, 32, rounds)

	var outputCount int
	var outputData []byte
	for i := 0; outputCount < numBits; i++ {
		inputData := []byte{
			byte(i & 0xff),
			byte((i >> 8) & 0xff),
			byte((i >> 16) & 0xff),
			byte((i >> 24) & 0xff),
		}
		state := gocube.SolvedCubieCube()
		hash.HashExact(&state, inputData)
		val := rubikhash.StateToBits(&state)
		if outputFile == "" {
			for j := uint(0); j < 32; j++ {
				if outputCount >= numBits {
					break
				}
				if val&(1<<j) == 0 {
					fmt.Print("0")
				} else {
					fmt.Print("1")
				}
				outputCount++
			}
		} else {
			outputData = append(
				outputData,
				byte(val&0xff),
				byte((val>>8)&0xff),
				byte((val>>16)&0xff),
				byte((val>>24)&0xff),
			)
			outputCount += 32
		}
	}
	if outputFile == "" {
		fmt.Println("")
	} else {
		essentials.Must(os.WriteFile(outputFile, outputData[:numBits/8], 0644))
	}
}
