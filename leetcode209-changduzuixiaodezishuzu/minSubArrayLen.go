package leetcode209_changduzuixiaodezishuzu

//大于某个target的最短子数组
func k(nums []int, target int) int {
	var left, right, res, cur int
	for right < len(nums) {
		cur += nums[right]
		for cur >= target {
			if res == 0 || right-left+1 < res {
				res = right - left + 1
			}
			cur -= nums[left]
			left++
		}
		right++
	}
	return res

}
