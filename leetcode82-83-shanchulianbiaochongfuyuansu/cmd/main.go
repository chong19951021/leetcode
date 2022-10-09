package main

import (
	"coding/common"
	leetcode80_83_shanchulianbiaochongfuyuansu "coding/leetcode82-83-shanchulianbiaochongfuyuansu"
	"fmt"
)

func main() {
	head := &common.ListNode{
		Next: &common.ListNode{
			Next: &common.ListNode{
				Next: &common.ListNode{
					Next: nil,
					Val:  5,
				},
				Val: 2,
			},
			Val: 2,
		},
		Val: 1,
	}
	leetcode80_83_shanchulianbiaochongfuyuansu.DeleteDuplicates(head)
	fmt.Println(head)
}
