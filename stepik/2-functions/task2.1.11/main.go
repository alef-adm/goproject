package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 2, 3, 5, 6))
}
func sumInt(num ...int) (int, int) {
	var l, s int
	var rez []int
	l = len(num)
	rez = append(rez, l)
	for _, n := range num {
		s += n
	}
	rez = append(rez, s)
	return l, s
}
