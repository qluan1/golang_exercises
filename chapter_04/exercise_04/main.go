package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "could not parse argument as integer\n")
		return
	}
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Original slice: %v\n", s)
	rotate(s, k)
	fmt.Printf("After rotate by %d: %v\n", k, s)
}

func reverse(s []int, start, end int) {
    for start < end {
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}

func rotate(s []int, k int) {
    n := len(s)
	if k < 0 {
		k = n + k % n
	}
	if k >=  n {
		k = k % n
	}    
    reverse(s, 0, n-1)
    reverse(s, 0, k-1)
    reverse(s, k, n-1)
}