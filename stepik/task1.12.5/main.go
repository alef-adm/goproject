package main

import "fmt"

func main() {
	//var workarray [16]uint8 = [16]uint8{99, 151, 137, 71, 117, 187, 20, 93, 187, 67, 1, 2, 3, 5, 7, 8}
	var workarray [16]uint8
	for i := 0; i < len(workarray); i++ {

		fmt.Scan(&workarray[i])

	}
	workarray[workarray[10]], workarray[workarray[11]] = workarray[workarray[11]], workarray[workarray[10]]
	workarray[workarray[12]], workarray[workarray[13]] = workarray[workarray[13]], workarray[workarray[12]]
	workarray[workarray[14]], workarray[workarray[15]] = workarray[workarray[15]], workarray[workarray[14]]

	fmt.Println(workarray[workarray[10]], workarray[workarray[11]])
	//workarray[12], workarray[13] = workarray[13], workarray[12]
	//workarray[14], workarray[15] = workarray[15], workarray[14]
	for i := 0; i < len(workarray)-6; i++ {
		//
		fmt.Print(workarray[i])
		fmt.Print(" ")
		//fmt.Println(inp[4]-1, inp[5]-1)
	}
	fmt.Print("type ok")
}
