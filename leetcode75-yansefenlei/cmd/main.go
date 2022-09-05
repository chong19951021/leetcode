package main

import (
	leetcode75_yansefenlei "coding/leetcode75-yansefenlei"
	"fmt"
)

// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//必须在不使用库的sort函数的情况下解决这个问题。
//
//
//
//示例 1：
//
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
//
//示例 2：
//
//输入：nums = [2,0,1]
//输出：[0,1,2]

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	// 为什么在方法里对数组重新赋值，不生效？
	leetcode75_yansefenlei.SortColors2(nums)
	for _, num := range nums {
		fmt.Println(num)
	}

}
