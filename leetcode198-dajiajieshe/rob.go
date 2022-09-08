package leetcode198_dajiajieshe

// 解法1
// 使用二维数组记录状态
func Rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make([][]int, len(nums))
	for i, _ := range tmp {
		tmp[i] = make([]int, 2)
	}
	tmp[0][0] = 0
	tmp[0][1] = nums[0]
	tmp[1][0] = nums[0]
	tmp[1][1] = nums[1]
	for i := 2; i < len(nums); i++ {
		tmp[i][0] = max(tmp[i-1][0], tmp[i-1][1])
		tmp[i][1] = tmp[i-1][0] + nums[i]
	}
	return max(tmp[len(nums)-1][0], tmp[len(nums)-1][1])

}

// 解法2
// 不使用二维数组记录状态
func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make([]int, len(nums))
	tmp[0] = nums[0]
	tmp[1] = max(tmp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp[i] = max(tmp[i-1], tmp[i-2]+nums[i])

	}
	return tmp[len(nums)-1]
}

// 解法3
// 滚动数组
func Rob3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prevPrev := nums[0]
	prev := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		cur := max(prev, prevPrev+nums[i])
		prevPrev = prev
		prev = cur
	}
	return prev

}
func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	}
	return num2
}
