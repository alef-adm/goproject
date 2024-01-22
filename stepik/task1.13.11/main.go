package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	//var s []int
	fmt.Scan(&a, &b)
	c = a - 1
	if a <= b {
		for i := a; i <= b; i++ {
			if i%7 == 0 {
				c = i
			}
		}
		if c >= a {
			fmt.Println(c)
		} else {
			fmt.Println("NO")
		}
	}

}
