package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%10 == 1 && a != 11 {
		fmt.Println(a, "korova")
	} else if (a%10 > 1 && a%10 < 5) && (a < 10 || a > 20) {
		fmt.Println(a, "korovy")
	} else {
		fmt.Println(a, "korov")
	}
}
