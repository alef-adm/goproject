package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r := []rune(text)
	f := 0
	//fmt.Printf("\ntext - %v \n r - %v", text, r)
	//fmt.Println("длина", len(r)-2)
	for i := 0; i < (len(r)-2)/2; i++ {
		//fmt.Println(r[i], r[len(r)-3-i])
		if r[i] == r[len(r)-3-i] {
			//fmt.Println("r", r[i], "len", r[len(r)-3-i])
			f = 1
			//fmt.Println(f)
		} else {
			f = 0
			//fmt.Println(f)
			break
		}
	}
	if f == 1 {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
