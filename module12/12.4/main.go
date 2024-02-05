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
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	selectionSort(ar)

	fmt.Println(ar)
}

func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[j] > ar[i] {
				ar[j], ar[i] = ar[i], ar[j]
			}
		}
	}
}
