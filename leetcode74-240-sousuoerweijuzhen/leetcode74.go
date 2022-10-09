package leetcode74_240_sousuoerweijuzhen

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
// 此解法也适用于240题
func searchMatrix(matrix [][]int, target int) bool {
	hight, width := 0, len(matrix[0])
	for hight < len(matrix) && width >= 0 {
		if matrix[hight][width] > target {
			width--
		} else if matrix[hight][width] < target {
			hight++
		} else {
			return true
		}

	}
	return false
}
