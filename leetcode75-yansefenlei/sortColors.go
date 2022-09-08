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

func findLength(nums1 []int, nums2 []int) int {
	var length int
	m := len(nums1)
	n := len(nums2)
	res := make([][]int, m+1)
	for i, _ := range res {
		res[i] = make([]int, n+1)
	}
	res[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = 0
			}
			if res[i][j] > length {
				length = res[i][j]
			}

		}

	}
	return length
}
