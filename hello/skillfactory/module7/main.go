package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	ar := []int{a, b, c, d}
	sort.Ints(ar)
	fmt.Println(ar[0])
}
