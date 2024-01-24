package main

import "fmt"

func main() {
	var a float64
	var first, second, third float64
	first = 0
	second = 1
	fmt.Scan(&a)
	for i := 0; float64(i) < second+first; i++ {
		third = second + first
		if a == first {
			fmt.Println(i)
			break
		}
		if first > a {
			fmt.Println("-1")
			break
		}
		//	fmt.Println(first)
		first = second
		second = third
	}

}
