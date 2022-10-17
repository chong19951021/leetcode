package main

import (
	tanxinsuanfa_huichanganpai "coding/c111-huichanganpai"
	"fmt"
)

func main() {
	meetings := []tanxinsuanfa_huichanganpai.Meeting{
		{5, 12},
		{14, 18},
		{3, 9},
		{1, 2},
	}
	max := tanxinsuanfa_huichanganpai.OrderMeeting(meetings)
	fmt.Println(max)

}
