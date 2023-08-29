package main

import (
	"fmt"
	"os"
	"strconv"
	"mft/lenconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		len, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprint(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		ft := lenconv.Foot(len)
		m := lenconv.Meter(len)
		fmt.Printf("%s = %s, %s = %s\n",
			ft, lenconv.FtToM(ft), m, lenconv.MToFt(m))
	}
}