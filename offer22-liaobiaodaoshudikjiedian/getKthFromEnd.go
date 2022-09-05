package offer22_liaobiaodaoshudikjiedian

import "coding/common"

// 解法一  结果逆向推导
// 倒数第 k个,即从正开始数 len - k
func GetKthFromEnd(head *common.ListNode, k int) *common.ListNode {
	if head == nil {
		return nil
	}
	curHead := head
	length := 1
	for head.Next != nil {
		head = head.Next
		length++
	}
	step := length - k

	for step = step - 1; step >= 0; step-- {
		curHead = curHead.Next
	}
	return curHead
}

// 解法2
func GetKthFromEnd2(head *common.ListNode, k int) *common.ListNode {
	if head == nil {
		return nil
	}
	curHead := head
	fastHead := head
	for k > 0 && fastHead.Next != nil {
		fastHead = fastHead.Next
		k--
	}
	for fastHead != nil {
		curHead = curHead.Next
		fastHead = fastHead.Next
	}
	return curHead

}
