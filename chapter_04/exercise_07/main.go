package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界!"
	arr := []byte(s)
	fmt.Printf("%v\n", string(arr))
	fmt.Println("===> reverse!!!")
	reverseChar(arr)
	fmt.Printf("%v\n", string(arr))
}

func reverseChar(a []byte) {
	head := a
	for len(head) > 0 {
		_, size := utf8.DecodeRune(head)
		reverseSlice(head, 0, size-1)
		head = head[size:]
	}
	reverseSlice(a, 0, len(a)-1)
}

func reverseSlice(a []byte, start, end int) {
	if end < start || start < 0 || end >= len(a) {
		fmt.Fprintln(os.Stderr, "wrong indices")
		return
	}
	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
}