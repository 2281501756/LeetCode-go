package leetcode

import (
	"testing"
)

func TestRQku0D(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	valid := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			} else {
				l++
				r--
			}
		}
		return true
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return valid(l+1, r) || valid(l, r-1)
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
