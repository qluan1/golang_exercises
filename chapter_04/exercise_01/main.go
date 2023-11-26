package main

import (
	"crypto/sha256"
	"fmt"
)

func main(){
	s1 := "x"
	s2 := "X"
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Printf(
		"Number of different bits of \"%s\" and \"%s\" is %d\n",
		s1,
		s2,
		CountNumBitDiffHash(&h1, &h2),
	)
}

func CountNumBitDiffHash(h1, h2 *[32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += CountNumBitDiffByte(h1[i], h2[i])
	}
	return count
}

func CountNumBitDiffByte(b1, b2 byte) int {
	diff := b1 ^ b2
	count := 0
	for i := 0; i < 8; i++ {
		if (diff & 1) == 1 { count++ }
		diff >>= 1
	}
	return count
}