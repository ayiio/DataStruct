package main

import (
	"fmt"
	binarysearch "section1/binarySearch"
)

func main() {
	BinarySearchT()
}

//有序数组二分查找测试
func BinarySearchT() {
	testList := [5]int{1, 3, 5, 7, 9}
	idx1 := binarysearch.BinarySearch(testList[:], 3)
	fmt.Println(idx1)
	idx2 := binarysearch.BinarySearch(testList[:], -1)
	fmt.Println(idx2)
}
