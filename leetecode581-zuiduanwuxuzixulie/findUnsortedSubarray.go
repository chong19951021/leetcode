package leetecode581_zuiduanwuxuzixulie

// 解题的关键就是找到两个边界：左边界和右边界
func findUnsortedSubarray(nums []int) int {

	if len(nums) == 1 {
		return 0
	}
	min := nums[len(nums)-1]
	max := nums[0]
	left, right := 0, -1 // 这样赋值的意义在于：如果数组遵循单调性，结果能为0
	// 找到右边界:此索引值右边的值必须为单调递增
	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			// 继续往下走，看是不是单调递增
			max = nums[i]
		} else {
			right = i
		}
	}
	// 找到左边界：此索引左边的值必须为单调递减
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			left = i
		}
	}
	return right - left + 1
}
