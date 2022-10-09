package leetcode80_26_shanchushuzuzhongchongfuyuansu

// leetcode26 删除有序数组中的重复元素，重复元素不超过一个
func removeDuplicates(nums []int) int {
	length := len(nums)
	if len(nums) <= 1 {
		return length
	}
	// slow 指针表示当前结果数组中的要被填充的索引，此索引之前的数字都是可以被保留的
	// fast 指针表示当前要被校验的索引
	slow, fast := 1, 1
	for fast < length {
		// 当前要校验的元素是否等于被保留的最后一位，因为重复次数只允许等于1.如果允许等于2，就slow-2
		if nums[fast] != nums[slow-1] {
			// 此时将当前元素加入可被保留数组中，更新slow索引
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// leetcode80 删除有序数组中的重复元素，重复元素不超过2个
func remove2(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	slow, fast := 2, 2
	for fast < length {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow

}
