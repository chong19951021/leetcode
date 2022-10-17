package paixusuanfa

// 快速排序算法
func quickSort(unSortList []int) {
	// 分裂准备
	quickSort1(unSortList, 0, len(unSortList)-1)
}

func quickSort1(arr []int, left, right int) {
	if left < right {
		pivot := arr[left]
		j := left
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		quickSort1(arr, left, j)
		quickSort1(arr, j+1, right)
	}
}

// 冒泡排序算法
func bubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
