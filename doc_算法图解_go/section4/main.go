package main

import (
	"fmt"
	quicksort "section4/quick_sort"
)

func main() {
	var arr = []int{19, 5, 2, 3, 1, 40, 6, 2, 1, 4, 8, 90, 0}
	res := quicksort.QuickSort(arr)
	fmt.Println(res)
}
