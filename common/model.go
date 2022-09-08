package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func Max(num1, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}
