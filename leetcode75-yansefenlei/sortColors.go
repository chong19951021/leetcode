package leetcode75_yansefenlei

// 解法1 ： 利用额外数组进行计数
func SortColors(nums []int) {
	//resSlice := make([]int, len(nums))
	//resSlice := nums
	if len(nums) <= 1 {
		return
	}
	tmpSlice := make([]int, 3)
	for _, num := range nums {
		tmpSlice[num]++
	}
	for i := 0; i < tmpSlice[0]; i++ {
		nums[i] = 0
	}
	for i := tmpSlice[0]; i < tmpSlice[0]+tmpSlice[1]; i++ {

		nums[i] = 1
	}
	for i := tmpSlice[0] + tmpSlice[1]; i < len(nums); i++ {
		nums[i] = 2
	}
}

func SortColors2(nums []int) {
	left := 0
	right := len(nums) - 1
	for i := 0; i < right; i++ {
		if left == right {
			return
		}
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			if nums[i] == 2 {
				i--
			}
		}
	}
}
