package main

import (
	"fmt"
)

func hashint64(val int64) uint64 {
	return uint64(val % 1000) // code here
}

func hashstr(s string) uint64 {

	return uint64(i % 1000)
}

func main() {
	fmt.Println(hashint64(543))
	fmt.Println(hashstr("green"))
}
