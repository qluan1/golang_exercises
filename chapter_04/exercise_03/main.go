package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Original array %v\n", arr)
	reverse(&arr)
	fmt.Printf("After reverse %v\n", arr)
}

/*
	array must have fixed length so this would not
	work for arbitrary array
*/
func reverse(pt *[6]int) { 
	for i, j := 0, len(*pt) - 1; i < j; {
		pt[i], pt[j] = pt[j], pt[i]
		i++
		j--
	}
}