package c111_huichanganpai

import (
	"sort"
)

// N 场演唱会, 以 [{startTime, endTime}…] 的形式给出, 计算出最多能听几场演唱会
type Meeting struct {
	Start int
	End   int
}

func OrderMeeting(meetings []Meeting) int {

	// 按照结束时间升序排序
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].End < meetings[j].End
	})
	max := 1
	endTime := meetings[0].End
	for i := 1; i < len(meetings); i++ {
		if meetings[i].Start > endTime {
			max++
			endTime = meetings[i].End
		}
	}
	return max
}
