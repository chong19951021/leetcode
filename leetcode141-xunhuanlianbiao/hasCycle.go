package leetcode141_xunhuanlianbiao

import "coding/common"

// 解法1
// 利用hash存储
func hasCycle(head *common.ListNode) bool {
	tmpMap := make(map[*common.ListNode]struct{})
	for head != nil {
		if _, ok := tmpMap[head]; ok {
			return true
		}
		tmpMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 解法2
// 快慢指针
func hasCycle1(head *common.ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false

}
