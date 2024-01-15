package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var d int
	c := 0
	var s []int
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		s = append(s, d)

	}
	for i := range s {
		if s[i] >= 0 {
			c++
		}
	}
	fmt.Print(c)

}
