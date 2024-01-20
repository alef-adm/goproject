package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var m int
	count := 0
	var s []int
	for i := 0; i < a; i++ {
		var t int
		fmt.Scan(&t)
		s = append(s, t)
	}
	for i, r := range s {
		if i == 0 {
			m = r
			count++
		}
		if i > 0 && m == r {
			count++
		}
		if i > 0 && r < m {
			m = r
			count = 1
		}

	}
	fmt.Println(count)
}
