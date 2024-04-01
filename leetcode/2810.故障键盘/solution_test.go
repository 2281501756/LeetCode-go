package leetcode

import (
	"testing"
)

func TestFaultyKeyboard(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func finalString(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			res = Reverse(res)
		} else {
			res += string(s[i])
		}
	}
	return res
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

//leetcode submit region end(Prohibit modification and deletion)
