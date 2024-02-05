package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 500)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	bubbleSort(ar)
	fmt.Println(ar)
	bubbleSortReversed(ar)
	fmt.Println(ar)
}

func bubbleSort(ar []int) {

	for i := 0; i < len(ar); i++ {
		c := 0
		for j := 0; j < len(ar); j++ {

			if j < len(ar)-1 {
				if ar[j] > ar[j+1] {
					ar[j], ar[j+1] = ar[j+1], ar[j]
					c++
				}
			}
			//fmt.Println(ar, c)

		}
		if c == 0 {
			break
		}
	}
}

func bubbleSortReversed(ar []int) {
	for i := 0; i < len(ar); i++ {
		c := 0
		for j := 0; j < len(ar); j++ {

			if j < len(ar)-1 {
				if ar[j] < ar[j+1] {
					ar[j], ar[j+1] = ar[j+1], ar[j]
					c++
				}
			}
			//fmt.Println(ar, c)

		}
		if c == 0 {
			break
		}
	}
}
