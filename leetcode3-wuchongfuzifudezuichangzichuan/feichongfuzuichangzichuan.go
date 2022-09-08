package leetcode3_wuchongfuzifudezuichangzichuan

func maxLenNoDouble(str string) int {
	var left, right int
	if len(str) == 1 || len(str) == 0 {
		return len(str)
	}
	length := 1
	curMap := make(map[uint8]*struct{})
	for right < len(str) {
		for curMap[str[right]] != nil {
			delete(curMap, str[left])
			left++
		}
		curMap[str[right]] = &struct{}{}
		if len(curMap) > length {
			length = len(curMap)
		}
		right++
	}
	return length
}
