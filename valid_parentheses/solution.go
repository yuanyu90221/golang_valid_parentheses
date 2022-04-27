package valid_parentheses

func isValid(s string) bool {
	// stack
	stack := []rune{}
	popMap := map[rune]rune{'}': '{', ')': '(', ']': '['}
	for _, val := range s {
		if popValue, exist := popMap[val]; !exist {
			stack = append(stack, val)
		} else {
			lenStack := len(stack)
			if lenStack == 0 {
				return false
			}
			if popValue != stack[lenStack-1] {
				return false
			}
			stack = stack[0 : lenStack-1]
		}
	}
	return len(stack) == 0
}
