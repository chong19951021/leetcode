package main

import (
	"coding/common"
	leetcode206_fanzhuanlianbiao "coding/leetcode206-fanzhuanlianbiao"
	"fmt"
)

func main() {

	head := common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 2,
			Next: &common.ListNode{
				Val: 3,
				Next: &common.ListNode{
					Val: 4,
					Next: &common.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	list := leetcode206_fanzhuanlianbiao.ReverseList(&head)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
