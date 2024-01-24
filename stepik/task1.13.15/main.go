package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int64
	fmt.Scan(&a)
	b := strconv.FormatInt(a, 2)
	fmt.Println(b)
}
