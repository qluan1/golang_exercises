package main

import "fmt"

func main() {
	s := []string{
		"a",
		"b",
		"b",
		"b",
		"c",
		"c",
		"b",
		"c",
		"c",
		"d",
	}
	fmt.Printf("slice: %v\n", s)
	fmt.Println("remove adjacent duplicate")
	removeAdjacentDuplicate(&s)
	fmt.Printf("slice: %v\n", s)
}

func removeAdjacentDuplicate(s *[]string) {
	n := len(*s)
	offset := 0
	pre := ""
	for i := 0; i < n; i++ {
		str := (*s)[i]
		if str == pre {
			offset++
			continue
		}
		(*s)[i-offset] = str
		pre = str
	}
	(*s) = (*s)[:n-offset]
}