package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wordCount := make(map[string]int)
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path, such as ./example")
		return
	}

	filePath := filepath.Clean(os.Args[1])
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exists", filePath)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		wordCount[word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	fmt.Printf("\nWord\tcount\n")
	for c, n := range wordCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
}