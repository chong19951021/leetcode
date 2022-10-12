package leetcode53_zuidazishuzuhe

// 动态规划
// nums[i] 表示直到当前索引下，所能达到的最大值
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果当前的元素加上之前的最大值是大于当前元素的，表示递增趋势还在继续
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max

}
