package main

import "fmt"

func main() {
	var a, b string
	var s []string
	fmt.Scan(&a, &b)
	for _, r := range a {
		if string(r) != b {
			s = append(s, string(r))
		}
	}
	for _, r := range s {
		fmt.Print(r)
	}
	//fmt.Println(s)
}
