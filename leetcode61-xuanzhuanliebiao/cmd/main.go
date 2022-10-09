package main

import (
	"coding/leetcode61-xuanzhuanliebiao"
	"fmt"
)

func main() {

	head := leetcode61_xuanzhuanliebiao.ListNode{

		Val: 1,
		Next: &leetcode61_xuanzhuanliebiao.ListNode{
			Val: 2,
			Next: &leetcode61_xuanzhuanliebiao.ListNode{
				Val: 3,
				Next: &leetcode61_xuanzhuanliebiao.ListNode{
					Val: 4,
					Next: &leetcode61_xuanzhuanliebiao.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(leetcode61_xuanzhuanliebiao.RotateRight(&head, 2).Val)
}
