package leetcode1143_zuichanggonggongzixulie

import (
	"coding/common"
)

func K(str1 string, str2 string) int {

	group := make([][]int, len(str1)+1)

	for i, _ := range group {
		group[i] = make([]int, len(str2)+1)
	}
	group[0][0] = 0
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				group[i][j] = group[i-1][j-1] + 1
			} else {
				group[i][j] = common.Max(group[i][j-1], group[i-1][j])
			}
		}
	}
	return group[len(str1)][len(str2)]

}
