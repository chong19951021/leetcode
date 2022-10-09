package leetcode61_xuanzhuanliebiao

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

// 示例 1：

// 输入：head = [1,2,3,4,5], k = 2
// rotate1 : 5->1->2->3->4
// rotate2 : 4->5->1->2->3
// 输出：[4,5,1,2,3]

// 示例 2：

// 输入：head = [0,1,2], k = 4
// rotate1 : 0->1->2
// rotate2 : 2->0->1
// rotate3 : 0->1->2
// rotate4 : 2->0->1

// 输出：[2,0,1]

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	node := head
	length := 1
	// 先找出链表的长度
	for node.Next != nil {
		node = node.Next
		length++
	}
	// 此时node指向的是尾节点,再指向头结点，形成环形链表
	node.Next = head
	// 以尾节点的视角思考问题，当移动k次以后,尾结点处于整个链表中的第几个位置
	tailIndex := k % length
	// 距离表示原尾结点到最终结果的尾结点的距离
	distance := length - tailIndex
	for distance > 0 {
		distance--
		node = node.Next
	}
	res := node.Next
	node.Next = nil
	return res
}
