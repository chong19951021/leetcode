package main

import (
	smartX_chongpaishuzu "coding/smartX-chongpaishuzu"
	"fmt"
)

// 提供一个排序好的数组，使其重新排列。排列格式为最小，最大，次小，次大，次次小，次次大
// eg . [7,6,5,4,3,2,1] -> [1,7,2,6,3,5,4]
func main() {
	nums := []int{

		8, 7, 6, 5, 4, 3, 2, 1,
	}

	slice := smartX_chongpaishuzu.SortSlice(nums)
	for _, i := range slice {
		fmt.Println(i)
	}
}
