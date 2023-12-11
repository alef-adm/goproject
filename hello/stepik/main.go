package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	for {
		fmt.Scan(&a, &b)
		if (a >= 0 && a <= 10000) && (b >= 0 && b <= 10000) && (a != b) {
			break
		}
	}
	astr := strconv.Itoa(a)
	bstr := strconv.Itoa(b)
	for i := 0; i < len(astr); i++ {
		for j := 0; j < len(bstr); j++ {
			if astr[i] == bstr[j] {
				fmt.Print(string(astr[i]) + " ")
			}
		}
	}
}
