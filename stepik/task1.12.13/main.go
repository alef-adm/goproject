package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var d int
	var s []int
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		s = append(s, d)

	}
	fmt.Println(s[3])
}
