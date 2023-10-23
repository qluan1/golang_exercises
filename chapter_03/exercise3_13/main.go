package main

import (
	"fmt"
)

const (
    KB = 1000
    MB = KB * KB
    GB = KB * MB
    TB = KB * GB
    PB = KB * TB
    EB = KB * PB
    ZB = KB * EB
    YB = KB * ZB
)

func main() {
    fmt.Println(KB, MB, GB, TB, PB, EB)
	fmt.Printf("1 ZB is %d EB\n", ZB/EB)
	fmt.Printf("1 YB is %d ZB\n", YB/ZB)
}
