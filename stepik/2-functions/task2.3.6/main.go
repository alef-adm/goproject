package main

import "fmt"

//Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
//ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.
//
//Sample Input:
//
//2 4
//Sample Output:
//
//4 2

func main() {
	var a, b int
	a = 1
	b = 2
	test(&a, &b)
}
func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1 // здесь ваш код
	fmt.Println(*x1, *x2)
}
