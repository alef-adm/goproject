package main

import "fmt"

func main() {
	fmt.Print(vote(1, 0, 0))
}
func vote(x int, y int, z int) int {
	//print your code here
	if x+y+z > 1 {
		return 1
	} else {
		return 0
	}
}
