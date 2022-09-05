package main

import (
	"coding/common"
	leetcode86_分隔链表 "coding/leetcode86-fengelianbiao"
	"fmt"
)

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
//你应当 保留 两个分区中每个节点的初始相对位置。
//
//示例 1：
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//
//示例 2：
//
//输入：head = [2,1], x = 2
//输出：[1,2]

func main() {

	head := &common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 4,
			Next: &common.ListNode{
				Val: 3,
				Next: &common.ListNode{
					Val: 2,
					Next: &common.ListNode{
						Val: 5,
						Next: &common.ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	//head := &common.ListNode{
	//	Val: 2,
	//	Next: &common.ListNode{
	//		Val:  1,
	//		Next: nil,
	//	},
	//}
	fmt.Println(leetcode86_分隔链表.Partition(head, 3).Val)
}
