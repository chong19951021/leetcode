package leetcode86_fengelianbiao

import "coding/common"

func Partition(head *common.ListNode, x int) *common.ListNode {
	lowerHead := new(common.ListNode)
	highHead := new(common.ListNode)
	resHead := lowerHead
	resHighHead := highHead
	for head != nil {
		if head.Val < x {
			lowerHead.Next = head
			lowerHead = lowerHead.Next
		} else {
			highHead.Next = head
			highHead = highHead.Next
		}
		head = head.Next
	}
	highHead.Next = nil
	lowerHead.Next = resHighHead.Next
	return resHead.Next
}
