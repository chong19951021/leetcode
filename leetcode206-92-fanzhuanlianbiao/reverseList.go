package leetcode206_92_fanzhuanlianbiao

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

// leetcode 92
//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
//请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
func reverseBetween(head *common.ListNode, left int, right int) *common.ListNode {

	if left == right {
		return head
	}
	// 因为头结点可能会变化，采用dummy节点
	dummy := &common.ListNode{Val: -1, Next: head}
	pre := dummy
	// rightNode 为第right个节点
	rightNode := head
	for right > 1 {
		rightNode = rightNode.Next
		right--
	}
	// nextNode为区间外右节点
	nextNode := rightNode.Next
	// 截断
	rightNode.Next = nil
	// pre为区间外左节点
	for left > 1 {
		pre = pre.Next
		left--
	}
	// 此为列表反转以后的尾节点
	curr := pre.Next
	nn := ReverseList(pre.Next)
	// 令区间外左节点指向反转以后的头节点
	pre.Next = nn
	// 令尾节点指向之前的区间外的右节点
	curr.Next = nextNode
	return dummy.Next
}
