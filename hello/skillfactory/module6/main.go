package main

import "fmt"

func main() {
	switch s := 15; {
	default:
		fmt.Println("this is the end")
	case s%4 == 1:
		fmt.Println("endorama")
	case s == int(60.0/4):
		fmt.Println("kreator")
	case s == 15:
		fmt.Println(1999)
	}
}
