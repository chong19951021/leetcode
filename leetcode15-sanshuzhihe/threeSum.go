package leetcode15_sanshuzhihe

import "sort"

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
//你返回所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
func threeSum(nums []int) [][]int {
	res := [][]int{}
	curMap := make(map[int]int)
	for _, num := range nums {
		curMap[num]++
	}
	singleNums := []int{}
	for key, _ := range curMap {
		singleNums = append(singleNums, key)
	}
	// 进行排序
	sort.Ints(singleNums)
	for i := 0; i < len(singleNums); i++ {
		if singleNums[i]*3 == 0 && curMap[singleNums[i]] >= 3 {
			res = append(res, []int{singleNums[i], singleNums[i], singleNums[i]})
		}
		for j := i + 1; j < len(singleNums); i++ {
			if singleNums[i]*2+singleNums[j] == 0 && curMap[singleNums[i]] >= 2 {
				res = append(res, []int{singleNums[i], singleNums[i], singleNums[j]})
			}
			if singleNums[j]*2+singleNums[i] == 0 && curMap[singleNums[j]] >= 2 {
				res = append(res, []int{singleNums[i], singleNums[j], singleNums[j]})
			}
			c := 0 - singleNums[i] - singleNums[j]
			if curMap[c] > 0 && c > singleNums[j] {
				res = append(res, []int{singleNums[i], singleNums[j], c})
			}
		}
	}
	return res
}
