package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "hello, \t\n世界\n\nwolrd"
	arr := []byte(s)
	fmt.Println(string(arr))
	fmt.Println("squash!")
	squash(&arr)
	fmt.Println(string(arr))	
}

func squash(p *[]byte) {
	head := *p
	writeIndex := 0
	preIsSpace := false
	for len(head) > 0 {
		r, size := utf8.DecodeRune(head)
		isSpace := unicode.IsSpace(r)
		if isSpace && !preIsSpace {
			(*p)[writeIndex] = ' '
			writeIndex++
		} else if !isSpace {
			for i := 0; i < size; i++ {
				(*p)[writeIndex] = head[i]
				writeIndex++
			}
		}
		preIsSpace = isSpace
		head = head[size:]
	}
	*p = (*p)[:writeIndex]
}