package main

import "fmt"

func main() {
	var num1, num2, num3, num4 int
	fmt.Scan(&num1)
	num2 = num1 % 10
	num3 = (num1 / 10) % 10
	num4 = num1 / 100
	fmt.Println(num2 + num3 + num4)
}
