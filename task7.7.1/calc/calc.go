package calc

import "fmt"

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}
func (c *calculator) Calculate(num1, num2 float64, oper string) float64 {
	switch oper {
	case "+":
		return c.addition(num1, num2)
	case "-":
		return c.subtraction(num1, num2)
	case "*":
		return c.multiplication(num1, num2)
	case "/":
		return c.division(num1, num2)
	default:
		fmt.Printf("несуществующий оператор: %s\n", oper)
		return 0
	}
}

func (c *calculator) addition(num1, num2 float64) float64 {
	return num1 + num2
}
func (c *calculator) subtraction(num1, num2 float64) float64 {
	return num1 - num2
}
func (c *calculator) multiplication(num1, num2 float64) float64 {
	return num1 * num2
}
func (c *calculator) division(num1, num2 float64) float64 {
	if num2 == 0 {
		fmt.Println("Ошибка, деление на ноль")
		return 0
	}
	return num1 / num2

}
