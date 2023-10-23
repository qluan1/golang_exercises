package main

import (
	"fmt"
	"os"
)

func main() {
	for i, val := range os.Args[1:] {
		fmt.Println(i, val)
	}
}
