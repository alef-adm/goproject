package main

import "fmt"

func main() {
	var inp [16]uint8
	//var workarray [10]uint8
	for i := 0; i < 16; i++ {

		fmt.Scan(&inp[i])

	}
	inp[10], inp[11] = inp[11], inp[10]
	inp[12], inp[13] = inp[13], inp[12]
	inp[14], inp[15] = inp[15], inp[14]
	for i := 0; i < 16; i++ {

		fmt.Print(inp[i])
		fmt.Print(" ")
		//fmt.Println(inp[4]-1, inp[5]-1)
	}
}
