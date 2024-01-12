package main

import (
	"fmt"

	"task7.7.1/calc"
)

func main() {
	var num1, num2 float64
	var oper string
	fmt.Println("введите первое число")
	fmt.Scanln(&num1)
	fmt.Println("введите второе число")
	fmt.Scanln(&num2)
	fmt.Println("введите желаемое действие")
	fmt.Scanln(&oper)
	calculator := calc.NewCalculator()
	result := calculator.Calculate(num1, num2, oper)
	fmt.Println(result)
}
