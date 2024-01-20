package main

import "fmt"

func main() {
	var a [100]int
	var sl int
	var num int
	fmt.Scan(&num)
	count := 0
	for i := 0; i < num; i++ {
		fmt.Scan(&sl)
		a[i] = sl
	}
	for i := 0; i < num; i++ {
		if a[i] == 0 {
			count++
		}
	}
	fmt.Println(count)
}
