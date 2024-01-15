package main

import "fmt"

func main() {
	array := [100]int{}
	var n int
	fmt.Scan(&n)
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")

	}
}
