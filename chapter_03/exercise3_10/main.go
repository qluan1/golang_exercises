package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	// assuming all char of s is digit
	// first char is not 0
	var buf bytes.Buffer
	n := len(s)
	if n == 0 {
		return ""
	}
	r := 3 - n%3
	for i := 0; i < n; i++ {
		if r == 0 {
			buf.WriteByte(',')
		}
		r = (r + 1)%3
		buf.WriteString(s[i:i+1])
	}
	return buf.String()
}

func main () {
	fmt.Println(comma("1")) // "1"
	fmt.Println(comma("10")) // "10"
	fmt.Println(comma("100")) // "100"
	fmt.Println(comma("1000")) // "1,000"
	fmt.Println(comma("100000000")) // "100,000,000"
	fmt.Println(comma("10000000000000")) // "10,000,000,000,000"
}