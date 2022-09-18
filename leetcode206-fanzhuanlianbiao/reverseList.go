package leetcode206_fanzhuanlianbiao

import "coding/common"

func ReverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	var lastNode *common.ListNode

	for head != nil {
		nextNode := head.Next
		head.Next = lastNode
		lastNode = head
		head = nextNode
	}
	return lastNode

}
