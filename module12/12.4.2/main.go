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
	fmt.Println("до сортировки", ar)
	selectionSort(ar)

	fmt.Println("после сортировки", ar)
}

func selectionSort(ar []int) {
	for i := 0; i < len(ar)-i; i++ {
		minK := i
		maxK := len(ar) - 1 - i
		var tmpmin, tmpmax int
		for j := i; j < len(ar)-i; j++ {
			fmt.Println("i=", i, "j=", j, "ar[j]=", ar[j], "ar[i]=", ar[i])
			if ar[j] < ar[minK] {
				minK = j
				tmpmin = ar[minK]
				fmt.Println("minK=", minK)
			}
			if ar[j] > ar[maxK] {
				maxK = j
				tmpmax = ar[maxK]
				fmt.Println("maxK=", maxK)

			}
		}
		if maxK == i && minK != len(ar)-1-i {
			ar[len(ar)-1-i], ar[maxK] = ar[maxK], ar[len(ar)-1-i]
			ar[i], ar[minK] = ar[minK], ar[i]
		} else if minK == len(ar)-1-i && maxK != i {
			ar[i], ar[minK] = ar[minK], ar[i]
			ar[len(ar)-1-i], ar[maxK] = ar[maxK], ar[len(ar)-1-i]
		} else if maxK == i && minK == len(ar)-1-i {
			mTemp := ar[maxK]
			ar[len(ar)-1-i], ar[maxK] = ar[maxK], ar[len(ar)-1-i]
			ar[len(ar)-1-i] = mTemp
		} else {
			if minK != 0 && ar[minK] == tmpmin {
				ar[i], ar[minK] = ar[minK], ar[i]

			}
			if maxK != 0 && ar[maxK] == tmpmax {
				ar[len(ar)-1-i], ar[maxK] = ar[maxK], ar[len(ar)-1-i]

			}
		}
	}

}
