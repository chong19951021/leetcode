package leetcode23_hebingKgelianbiao

import (
	"coding/common"
	"container/heap"
)

func mergeKLists(lists []*common.ListNode) *common.ListNode {
	var h = new(ListsGroup)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}
	var dummy = new(common.ListNode)
	var head = dummy
	for h.Len() > 0 {
		top := heap.Pop(h).(*common.ListNode)
		if top.Next != nil {
			heap.Push(h, top.Next)
		}
		head.Next, head = top, top
	}
	return dummy.Next

}

type ListsGroup []*common.ListNode

func (lg ListsGroup) Len() int {
	return len(lg)
}
func (lg ListsGroup) Less(i, j int) bool {
	return lg[i].Val < lg[j].Val
}
func (lg ListsGroup) Swap(i, j int) {
	lg[i], lg[j] = lg[j], lg[i]
}
func (lg *ListsGroup) Pop() interface{} {
	x := (*lg)[lg.Len()-1]
	*lg = (*lg)[:(lg.Len() - 1)]
	return x
}
func (lg *ListsGroup) Push(any interface{}) {
	node := any.(*common.ListNode)
	*lg = append(*lg, node)
}
