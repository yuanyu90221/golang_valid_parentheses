# valid parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 
## Examples

```
Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
```
## 解析

每個字元照順序找對應的配對字元

要能夠完全配對代表每次輸入的字元配對字元在順序上遇早出現遇晚找到

這個題目可以透過 stack 來做處理

stack 特性後進先出

剛好可以讓這個來做比對

每個對應的配對字元可以透過 map 來做配對

## 程式碼

```golang
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

```