package main

import "fmt"

func main() {
	const (
		january  = iota + 1
		february = iota + 1
		march    = iota + 1
		april    = iota + 1
		may      = iota + 1
	)
	fmt.Println(january, february, march, april, may)
}
