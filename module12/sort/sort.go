package sort

import (
	"math/rand"
)

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

func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[j] > ar[i] {
				ar[j], ar[i] = ar[i], ar[j]
			}
		}
	}
}

func dualSelectionSort(ar []int) {
	for i := 0; i < len(ar)-i; i++ {
		minK := i
		maxK := len(ar) - 1 - i
		var tmpmin, tmpmax int
		for j := i; j < len(ar)-i; j++ {
			if ar[j] < ar[minK] {
				minK = j
				tmpmin = ar[minK]
			}
			if ar[j] > ar[maxK] {
				maxK = j
				tmpmax = ar[maxK]

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

func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		for j := i; j > 0 && ar[j-1] > ar[j]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}
}

func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	middle := len(ar) / 2

	sortedAr := make([]int, 0, len(ar))
	left, right := mergeSort(ar[:middle]), mergeSort(ar[middle:])

	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			sortedAr = append(sortedAr, right[j])
			j++
		} else {
			sortedAr = append(sortedAr, left[i])
			i++
		}
	}

	sortedAr = append(sortedAr, left[i:]...)
	sortedAr = append(sortedAr, right[j:]...)

	return sortedAr
}
func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar)-1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left+1:])

	return
}
