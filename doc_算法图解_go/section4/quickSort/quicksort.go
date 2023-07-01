package quicksort

func QuickSort(arr []int) (res []int) {
	if len(arr) < 2 {
		return arr //基线条件，为空或只包含一个元素的数组是有序的
	} else {
		pivot := arr[0] //基准值，递归条件
		var lsArr []int //小于基准值的元素组成的左子数组
		var gtArr []int //大于基准值的元素组成的右子数组
		lens := len(arr)
		for i := 1; i < lens; i++ {
			if arr[i] < pivot {
				lsArr = append(lsArr, arr[i])
			} else {
				gtArr = append(gtArr, arr[i])
			}
		}
		res = append(res, QuickSort(lsArr)...)
		res = append(res, pivot)
		res = append(res, QuickSort(gtArr)...)
	}
	return
}
