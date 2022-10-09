package leetcode21_1669_hebinglianbiao

import "coding/common"

// leetcode-21
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	resNode := &common.ListNode{
		Val:  -1,
		Next: nil,
	}
	curNode := resNode
	for true {
		if list1 == nil {
			resNode.Next = list2
			break
		}
		if list2 == nil {
			resNode.Next = list1
			break
		}
		if list1.Val >= list2.Val {
			resNode.Next = list2
			list2 = list2.Next
		} else {
			resNode.Next = list1
			list1 = list1.Next

		}
		resNode = resNode.Next

	}
	return curNode.Next
}

// leetcode-1669
//给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
//
//请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。
//
//下图中蓝色边和节点展示了操作后的结果：

func mergeInBetween(list1 *common.ListNode, a int, b int, list2 *common.ListNode) *common.ListNode {
	var preNode *common.ListNode
	resNode := list1
	for i := 0; i <= b; i++ {
		if i == a-1 {
			preNode = list1
		}
		list1 = list1.Next
	}
	preNode.Next = list2
	for list2.Next != nil {

		list2 = list2.Next

	}
	list2.Next = list1
	return resNode
}
