package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var s string
	fmt.Scan(&a)
	fmt.Scan(&s)
	fmt.Scan(&b)
	switch s {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	}
}
