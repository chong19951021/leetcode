package leetcode121_maimaigupiaodezuijiashiji

import "math"

// 买卖股票的最佳时机
func BuyTime(nums []int) int {
	minPrice := math.MaxInt
	var maxGold int
	for i := 0; i < len(nums); i++ {
		if nums[i]-minPrice > maxGold {
			maxGold = nums[i] - minPrice
		} else if nums[i] < minPrice {
			minPrice = nums[i]
		}
	}
	return maxGold
}
