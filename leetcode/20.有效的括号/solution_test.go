package leetcode

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			} else if s[i] == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if s[i] == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else if s[i] == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
