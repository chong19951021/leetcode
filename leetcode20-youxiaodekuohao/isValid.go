package leetcode20_youxiaodekuohao

func IsValid(s string) bool {

	if len(s)%2 == 1 {
		return false
	}

	tmpMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := make([]string, 0)

	for _, str := range s {
		if string(str) == "(" || string(str) == "[" || string(str) == "{" {
			stack = append(stack, string(str))
		} else if len(stack) == 0 || stack[len(stack)-1] != tmpMap[string(str)] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
