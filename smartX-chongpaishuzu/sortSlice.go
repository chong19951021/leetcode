package smartX_chongpaishuzu

func SortSlice(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			res[i] = nums[right]
			right--
		} else {
			res[i] = nums[left]
			left++
		}
	}
	return res

}
