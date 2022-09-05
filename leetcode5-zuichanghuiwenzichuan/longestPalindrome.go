package leetcode5_zuichanghuiwenzichuan

// 动态规划
func LongestPalindrome(s string) string {
	return ""
}

// 中心扩展法
func LongestPalindrome2(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		var max string
		s1 := k(s, i, i)
		s2 := k(s, i, i+1)
		if len(s1) > len(s2) {
			max = s1
		} else {
			max = s2
		}
		if len(max) > len(res) {
			res = max
		}

	}
	return res
}

func k(src string, left, right int) string {

	for left >= 0 && right < len(src) && src[left] == src[right] {
		left--
		right++
	}
	return src[left+1 : right]

}
