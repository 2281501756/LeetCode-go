package leetcode

import (
	"testing"
)

func TestValidParenthesisString(t *testing.T) {
	validString := checkValidString("((()")
	t.Log(validString)
}

// leetcode submit region begin(Prohibit modification and deletion)
func checkValidString(s string) bool {
	stack, star := make([]int, 0), make([]int, 0)
	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else if ch == '*' {
			star = append(star, i)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else if len(star) > 0 {
				star = star[:len(star)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	if len(star) < len(stack) {
		return false
	}
	for ; len(stack) > 0; stack, star = stack[:len(stack)-1], star[:len(star)-1] {
		if stack[len(stack)-1] > star[len(star)-1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
