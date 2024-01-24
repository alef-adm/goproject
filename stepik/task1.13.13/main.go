package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	for i := 0; float64(i) < n; i++ {
		a := math.Pow(2, float64(i))
		if a < n {
			fmt.Print(a, " ")
		} else {
			break
		}
	}
}
