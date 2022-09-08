package leetcode448_shuzuzhongxiaoshideshuzi

import "math"

// 解法1 暴力破解,使用map存储。浪费空间
func findDisappearedNumbers(nums []int) []int {

	var res []int
	curMap := make(map[int]struct{}, len(nums))
	for i := 1; i <= len(nums); i++ {

		curMap[i] = struct{}{}

	}
	for _, num := range nums {
		if _, ok := curMap[num]; ok {

			delete(curMap, num)
		}
	}
	for k, _ := range curMap {
		res = append(res, k)
	}

	return res

}

// 原地修改 使用一个数组来解决
// 如果这个数出现过，就在对应的数组索引下打个标记。因为是[1,n],所以正好可以用数组索引表示。
func FindDisappearedNumbers2(nums []int) []int {
	var res []int
	for i := range nums {

		nums[int(math.Abs(float64(nums[i])))-1] = -int(math.Abs(float64(nums[nums[i]-1])))

	}
	for i := range nums {
		if nums[i] >= 0 {
			res = append(res, i+1)
		}
	}
	return res
}
