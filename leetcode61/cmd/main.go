package main

import (
	"coding/leetcode61"
	"fmt"
)

func main() {

	head := leetcode61.ListNode{

		Val: 1,
		Next: &leetcode61.ListNode{
			Val: 2,
			Next: &leetcode61.ListNode{
				Val: 3,
				Next: &leetcode61.ListNode{
					Val: 4,
					Next: &leetcode61.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(leetcode61.RotateRight(&head, 2).Val)
}
