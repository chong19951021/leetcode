package leetcode718_zuichangchongfuzishuzu

// 动态规划
func findLength(nums1 []int, nums2 []int) int {
	var length int
	m := len(nums1)
	n := len(nums2)
	// 二维数组res[i][j]表示以索引i-1，j-1的下标所对应元素为子数组尾元素时的长度。
	// 我们要确定边界值。如果不-1的话，res[0][0]的值无法表示
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
