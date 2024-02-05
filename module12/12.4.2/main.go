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
	ar := make([]int, 5)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	selectionSort(ar)

	fmt.Println(ar)
}

func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		start := i
		finish := len(ar) - 1 - i
		nMin := ar[start]
		nMax := ar[finish]
		fmt.Println(nMin, nMax)
		var min_key, max_key int
		//fmt.Println("start", start, "finish", finish)
		for j := start; j == finish; j++ {
			if ar[j] <= nMin {
				nMin = ar[j]
				min_key = j
			}
			if ar[j] >= nMax {
				nMax = ar[j]
				max_key = j
			}
			//fmt.Println(min_key, max_key)
		}
		ar[start], ar[min_key] = ar[min_key], ar[start]
		ar[finish], ar[max_key] = ar[max_key], ar[finish]
		fmt.Println(ar)
	}

}
