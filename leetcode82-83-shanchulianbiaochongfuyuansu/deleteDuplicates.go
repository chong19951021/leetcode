package leetcode82_83_shanchulianbiaochongfuyuansu

import (
	"coding/common"
)

// leetcode83 重复的留一个

func DeleteDuplicates(head *common.ListNode) *common.ListNode {
	lastNode := &common.ListNode{
		Val:  -1,
		Next: head,
	}
	res := lastNode
	curMap := make(map[int]struct{})
	for head != nil {
		if _, ok := curMap[head.Val]; ok {
			if head.Next != nil && head.Next.Val != lastNode.Val || head.Next == nil {
				lastNode.Next = head.Next
				lastNode = head.Next
			}
		} else {
			curMap[head.Val] = struct{}{}
			lastNode = head
		}
		head = head.Next
	}
	return res.Next
}

// leetcode82 要求是只要有重复的，就不留
func deleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// leetcode82 要求是只要有重复的，就不留
func deleteDuplicates82(head *common.ListNode) *common.ListNode {
	curMap := make(map[int]int)
	curHead := head
	for head != nil {
		curMap[head.Val]++
		head = head.Next
	}
	lastNode := &common.ListNode{
		Next: nil,
		Val:  -1,
	}
	var resNode *common.ListNode
	for curHead != nil {
		if val, _ := curMap[curHead.Val]; val > 1 {
			if lastNode.Next == nil {
				resNode = lastNode
			}
			lastNode.Next = curHead.Next
		} else {
			if lastNode.Next == nil {
				resNode = lastNode
			}
			lastNode.Next = curHead
			lastNode = curHead
		}
		curHead = curHead.Next
	}
	if resNode == nil {
		return resNode
	}
	return resNode.Next
}
