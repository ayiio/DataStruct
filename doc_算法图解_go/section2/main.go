package main

import (
	"fmt"
	selectsort "section2/selectSort"
)

func main() {
	arr := [...]int{5, 3, 6, 2, 10}
	selectSortT(arr[:])
}

func selectSortT(arr []int) {
	res := selectsort.SelectSort(arr)
	fmt.Println(res)
}
