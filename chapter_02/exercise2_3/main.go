package main

import (
	"fmt"
	"strconv"
	"os"
	"exercise2_3/popcount"
)

func main() {
	input := os.Args[1]
	x, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", popcount.PopCount(x))
}