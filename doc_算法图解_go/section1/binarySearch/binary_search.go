package binarysearch

//有序列表
func BinarySearch(list []int, item int) int {
  //左右边界
	L := 0
	R := len(list) - 1

  //直到范围缩小到只包含一个元素
	for L < R {
		mid := (L + R) >> 1  //中间元素
		guess := list[mid]
		if guess == item {   //找到
			return mid
		} else if guess > item { //剩余左区间
			R = mid
		} else {  //剩余右区间
			L = mid
		}
	}
	return -1  //未找到返回-1
}
