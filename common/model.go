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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
