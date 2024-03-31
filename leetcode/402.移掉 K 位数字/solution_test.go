package leetcode

import (
	"testing"
)

func TestRemoveKDigits(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	s := make([]int32, 0)
	for _, ch := range num {
		for k > 0 && len(s) > 0 && s[len(s)-1] > ch {
			s = s[0 : len(s)-1]
			k--
		}
		if ch != '0' || len(s) > 0 {
			s = append(s, ch)
		}
	}
	for k > 0 && len(s) != 0 {
		s = s[0 : len(s)-1]
		k--
	}
	if len(s) == 0 {
		return "0"
	}
	return string(s)
}

//leetcode submit region end(Prohibit modification and deletion)
