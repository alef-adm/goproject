package main

import "fmt"

func main() {
	var num1, m, h int
	fmt.Scan(&num1)
	h = num1 / 3600
	m = (num1 % 3600) / 60
	fmt.Print("It is ", h, " hours ", m, " minutes.")
}
