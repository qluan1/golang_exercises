package main

import (
	"bytes"
	"fmt"
)

func breakString(s string) (string, string, string) {
	// assume s is valid string of floating point number with possible sign
	var sign string
	var intPart string
	var decPart string

	cur := 0

	if (s[cur] == '+' || s[cur] == '-') {
		sign = s[cur:cur+1]
		cur++
	}

	for cur < len(s) && s[cur] != '.' {
		cur++
	}

	if (len(sign) == 0) {
		intPart = s[0: cur]
	} else {
		intPart = s[1: cur]
	}

	decPart = s[cur:]

	return sign, intPart, decPart
}

func commaInt(s string) string {
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

func comma(s string) string {
	sign, intPart, decPart := breakString(s)
	return sign + commaInt(intPart) + decPart
}

func main () {
	fmt.Println(comma("1"))
	fmt.Println(comma("10"))
	fmt.Println(comma("100"))
	fmt.Println(comma("1000"))
	fmt.Println(comma("100000000"))
	fmt.Println(comma("10000000000000"))

	fmt.Println(comma("+1"))
	fmt.Println(comma("-10"))
	fmt.Println(comma("+100"))
	fmt.Println(comma("-1000"))
	fmt.Println(comma("+100000000"))
	fmt.Println(comma("-10000000000000"))

	fmt.Println(comma("1.001"))
	fmt.Println(comma("10.0001"))
	fmt.Println(comma("100.001"))
	fmt.Println(comma("1000.1"))
	fmt.Println(comma("100000000.0"))
	fmt.Println(comma("0.1"))

	fmt.Println(comma("-1.001"))
	fmt.Println(comma("+10.0001"))
	fmt.Println(comma("-100.001"))
	fmt.Println(comma("+1000.1"))
	fmt.Println(comma("-100000000.0"))
	fmt.Println(comma("+0.1"))
}