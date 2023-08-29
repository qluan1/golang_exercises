package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(filenames[line], ", "))
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, filenames map[string][]string) {
	tempMap := make(map[string]int)
	input := bufio.NewScanner((f))
	for input.Scan() {
		counts[input.Text()]++
		tempMap[input.Text()]++
		if tempMap[input.Text()] == 1 {
			filenames[input.Text()] = append(filenames[input.Text()], filename)
		}
	}
}
