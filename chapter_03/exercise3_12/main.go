package main

import (
	"fmt"
	"unicode"
)

func areAnagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	runeCount := make(map[rune]int)

	for _, runeValue := range s1 {
		runeCount[unicode.ToLower(runeValue)]++
	}

	for _, runeValue := range s2 {
		runeCount[unicode.ToLower(runeValue)]--
	}

	for _, count := range runeCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	s1 := "Listen"
	s2 := "Silent"
	fmt.Println(areAnagrams(s1, s2))  // Output: true

	s1 = "hello"
	s2 = "billion"
	fmt.Println(areAnagrams(s1, s2))  // Output: false
}
