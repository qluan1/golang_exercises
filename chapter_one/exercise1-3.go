package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s := ""
	sep := ""
	start := time.Now().UnixMilli()
	for _, val := range os.Args[1:] {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)
	end := time.Now().UnixMilli()
	fmt.Println(start, end)
	fmt.Println("Inefficient method took", end-start)

	start = time.Now().UnixMilli()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end = time.Now().UnixMilli()
	fmt.Println(start, end)
	fmt.Println("Strings.Join method took", end-start)
}
