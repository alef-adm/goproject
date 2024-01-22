package main

import (
	"fmt"
)

func main() {
	var number int64
	fmt.Scan(&number)
	d := number
	var n1, total int64
	total = 0
	//	var sum []int64
	if total == 0 || total > 9 {
		for d > 0 {
			n1 = d % 10
			//			sum = append(sum, n1)
			d = d / 10
			total += n1
			if d == 0 && total > 9 {
				d = total
				total = 0
			}
		}
		fmt.Println(total)
	}
}
