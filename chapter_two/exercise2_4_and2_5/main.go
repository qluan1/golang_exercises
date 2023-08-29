package main

import (
	"fmt"
	"strconv"
	"os"
	"exercise2_4_and2_5/popcount"
)

func main() {
	input := os.Args[1]
	x, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Exercise 2.4: %d\nExercise 2.5: %d\n", popcount.PopCount1(x), popcount.PopCount2(x))
}